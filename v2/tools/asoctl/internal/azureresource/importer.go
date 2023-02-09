/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package azureresource

import (
	"bufio"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-service-operator/v2/api"
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
	importedResources []genruntime.KubernetesResource
}

func NewImporter() *Importer {
	return &Importer{
		scheme: api.CreateScheme(),
	}
}

// Import downloads the specified resource and adds it to our list for export
func (i *Importer) Import(resource string) error {
	klog.Infof("Importing %s", resource)

	gk, err := i.parseGroupKind(resource)
	if err != nil {
		return errors.Wrap(err, "unable to import resource")
	}

	gvk, err := i.findGVK(gk)
	if err != nil {
		return errors.Wrap(err, "unable to import resource")
	}

	klog.Infof("Importing %s/%s/%s", gvk.Group, gvk.Version, gvk.Kind)

	obj, err := i.scheme.New(gvk)
	if err != nil {
		return errors.Wrap(err, "unable to import resource")
	}

	obj.GetObjectKind().SetGroupVersionKind(gvk)

	kr, ok := obj.(genruntime.KubernetesResource)
	if !ok {
		return errors.Errorf("unable to import resource %s/%s/%s", gvk.Group, gvk.Version, gvk.Kind)
	}

	i.importedResources = append(i.importedResources, kr)

	return nil
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
