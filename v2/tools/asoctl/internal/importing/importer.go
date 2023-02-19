/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"bufio"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-service-operator/v2/api"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/version"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/tools/generator/pkg/naming"
	"github.com/Azure/azure-service-operator/v2/tools/generator/pkg/versions"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	"io"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	"net/url"
	"os"
	"sigs.k8s.io/yaml"
	"strings"
)

type Importer struct {
	scheme            *runtime.Scheme
	importedResources []genruntime.MetaObject
	cloudConfig       cloud.Configuration
	factory           ResourceImporterFactory
	armFactory        ARMResourceImporterFactory
}

var userAgent = "asoctl/" + version.BuildVersion

func NewImporter(
	config cloud.Configuration,
) (*Importer, error) {

	client, err := createARMClient(config)
	if err != nil {
		return nil, err
	}

	serviceConfig := config.Services[cloud.ResourceManager]

	factory := newResourceImporterFactory(api.CreateScheme())
	armFactory := factory.CreateARMFactory(client, serviceConfig)

	return &Importer{
		scheme:      api.CreateScheme(),
		cloudConfig: config,
		factory:     factory,
		armFactory:  armFactory,
	}, nil
}

// Import downloads the specified resource and adds it to our list for export
func (i *Importer) ImportFromARMID(ctx context.Context, armID string) error {

	importer, err := i.armFactory.CreateForARMID(armID)
	if err != nil {
		return errors.Wrap(err, "unable to import resource")
	}

	result, err := importer.Import(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to import resource")
	}

	i.importedResources = append(i.importedResources, result.Object)

	return nil

	//armConfig := cloud.AzurePublic.Services[cloud.ResourceManager]
	//
	//f := newResourceImporterFactory(api.CreateScheme())
	//f.configureARM(client, armConfig)
	//
	//importer, err := f.CreateForARMID(resource)
	//if err != nil {
	//	return errors.Wrap(err, "unable to import resource")
	//}
	//
	//result, err := importer.Import(context.Background())
	//
	//klog.Infof("Importing %s", resource)
	//
	//gk, err := i.parseGroupKind(resource)
	//
	//gvk, err := i.findGVK(gk)
	//if err != nil {
	//	return errors.Wrap(err, "unable to import resource")
	//}
	//
	//klog.Infof("Importing %s/%s/%s", gvk.Group, gvk.Version, gvk.Kind)
	//
	//obj, err := i.scheme.New(gvk)
	//if err != nil {
	//	return errors.Wrap(err, "unable to import resource")
	//}
	//
	//obj.GetObjectKind().SetGroupVersionKind(gvk)
	//
	//kr, ok := obj.(genruntime.ARMMetaObject)
	//if !ok {
	//	return errors.Errorf("unable to import resource %s/%s/%s", gvk.Group, gvk.Version, gvk.Kind)
	//}
	//
	//client, err := i.createARMClient()
	//if err != nil {
	//	return errors.Wrap(err, "unable to import resource")
	//}
	//
	//req, err := i.createGetByIDRequest(context.TODO(), resource, kr.GetAPIVersion())
	//if err != nil {
	//	return errors.Wrap(err, "unable to import resource")
	//}
	//
	//resp, err := client.Do(req)
	//if err != nil {
	//	return errors.Wrap(err, "unable to import resource")
	//}
	//
	//armStatus, err := genruntime.NewEmptyARMStatus(kr, i.scheme)
	//if err != nil {
	//	return err
	//}
	//
	//if !azruntime.HasStatusCode(resp, http.StatusOK) {
	//	return azruntime.NewResponseError(resp)
	//}
	//
	//// Create an owner reference
	//
	//var knownOwner genruntime.ArbitraryOwnerReference
	////if owner := kr.Owner(); owner != nil {
	////	knownOwner = genruntime.ArbitraryOwnerReference{
	////		Name:  owner.Name,
	////		Group: owner.Group,
	////		Kind:  owner.Kind,
	////	}
	////}
	//
	//if err := azruntime.UnmarshalAsJSON(resp, armStatus); err != nil {
	//	return err
	//}
	//
	//// Convert the ARM shape to the Kube shape
	//status, err := genruntime.NewEmptyVersionedStatus(kr, i.scheme)
	//if err != nil {
	//	return errors.Wrapf(err, "constructing Kube status object for resource: %q", resource)
	//}
	//
	//if s, ok := status.(genruntime.FromARMConverter); ok {
	//	err = s.PopulateFromARM(knownOwner, reflecthelpers.ValueOfPtr(armStatus)) // TODO: PopulateFromArm expects a value... ick
	//	if err != nil {
	//		return errors.Wrapf(err, "converting ARM status to Kubernetes status")
	//	}
	//}
	//
	//err = kr.SetStatus(status)
	//if err != nil {
	//	return errors.Wrapf(err, "setting status on Kubernetes resource")
	//}
	//
	//i.importedResources = append(i.importedResources, kr)
	//
	//return nil
}

func (i *Importer) SaveToWriter(destination io.Writer) error {
	buf := bufio.NewWriter(destination)
	defer func(buf *bufio.Writer) {
		_ = buf.Flush()
	}(buf)

	buf.WriteString("---\n")
	for _, resource := range i.importedResources {
		data, err := yaml.Marshal(resource)
		if err != nil {
			return errors.Wrap(err, "unable to save to writer")
		}

		buf.Write(data)
		buf.WriteString("---\n")
	}

	return nil
}

func (i *Importer) SaveToFile(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return errors.Wrapf(err, "unable to create file %s", filepath)
	}

	defer func() {
		file.Close()

		// if we are panicking, the file will be in a broken
		// state, so remove it
		if r := recover(); r != nil {
			os.Remove(filepath)
			panic(r)
		}
	}()

	err = i.SaveToWriter(file)
	if err != nil {
		// cleanup in case of errors
		file.Close()
		os.Remove(filepath)
	}

	return errors.Wrapf(err, "unable to save to file %s", filepath)
}

// parseGroupKind parses a GroupKind from the resource URL, allowing us to look up the actual resource
func (i *Importer) parseGroupKind(resource string) (schema.GroupKind, error) {
	id, err := i.parseIdFromResourceUrl(resource)
	if err != nil {
		return schema.GroupKind{},
			errors.Wrapf(err, "unable to parse GroupKind from resource URL: %s", resource)
	}

	return schema.GroupKind{
		Group: i.extractGroupFromId(id),
		Kind:  i.extractKindFromId(id),
	}, nil
}

// parseIdFromResourceUrl parses an ARM ID from the resource URL
func (i *Importer) parseIdFromResourceUrl(resourceUrl string) (*arm.ResourceID, error) {
	u, err := url.Parse(resourceUrl)
	if err != nil {
		klog.Errorf("failed to parse resource URL: %s", err)
		return nil, errors.Wrapf(err, "failed to parse resource URL: %s", resourceUrl)
	}

	id, err := arm.ParseResourceID(u.Path)
	if err != nil {
		klog.Errorf("failed to parse ARM ID: %s", err)
		return nil, errors.Wrapf(err, "failed to parse ARM ID: %s", u.Path)
	}

	return id, nil
}

// extractGroupFromId extracts an ASO group name from the ARM ID
func (i *Importer) extractGroupFromId(id *arm.ResourceID) string {
	parts := strings.Split(id.ResourceType.Namespace, ".")
	last := len(parts) - 1
	return strings.ToLower(parts[last]) + ".azure.com"
}

// extractKindFromId extracts an ASO kind from the ARM ID
func (i *Importer) extractKindFromId(id *arm.ResourceID) string {
	if len(id.ResourceType.Types) != 1 {
		panic("Don't currently know how to handle nested resources")
	}

	kind := naming.Singularize(id.ResourceType.Types[0])
	return kind
}

func (i *Importer) findGVK(gk schema.GroupKind) (schema.GroupVersionKind, error) {
	knownVersions := i.scheme.VersionsForGroupKind(gk)
	if len(knownVersions) == 0 {
		return schema.GroupVersionKind{},
			errors.Errorf("no known versions for Group %s, Kind %s", gk.Group, gk.Kind)
	}

	return i.selectVersion(gk, knownVersions), nil
}

func (i *Importer) selectVersion(gk schema.GroupKind, knownVersions []schema.GroupVersion) schema.GroupVersionKind {
	// Sort the versions the same way we do in the generator, for consistency
	slices.SortFunc(
		knownVersions,
		func(left schema.GroupVersion, right schema.GroupVersion) bool {
			return versions.Compare(left.Version, right.Version)
		})

	// Ideally we want to find the latest stable version, but if there isn't one we'll take the latest preview.
	// Preview versions might introduce odd behaviour, so we err on the side of caution.
	// Storage versions need to be skipped though, as they don't have a fixed OriginalVersion()
	var previewVersion schema.GroupVersion
	var stableVersion schema.GroupVersion
	for _, gv := range knownVersions {
		if i.isStorageVersion(gv.Version) {
			// Skip storage versions
			continue
		}

		if versions.IsPreview(gv.Version) {
			previewVersion = gv
		} else {
			stableVersion = gv
		}
	}

	if !stableVersion.Empty() {
		return stableVersion.WithKind(gk.Kind)
	}

	return previewVersion.WithKind(gk.Kind)
}

func (i *Importer) isStorageVersion(version string) bool {
	return strings.HasSuffix(version, "storage")
}

func createARMClient(cloudConfig cloud.Configuration) (*azruntime.Pipeline, error) {
	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get default azure credential")
	}

	opts := &arm.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cloudConfig,
			PerCallPolicies: []policy.Policy{
				genericarmclient.NewUserAgentPolicy(userAgent),
			},
		},
	}

	pipeline, err := armruntime.NewPipeline("generic", version.BuildVersion, creds, azruntime.PipelineOptions{}, opts)
	return &pipeline, err
}
