/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package template

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/rotisserie/eris"
)

// URIFromVersion creates an ASO template URI from a version
func URIFromVersion(version string) string {
	tmpl := "https://github.com/Azure/azure-service-operator/releases/download/%s/azureserviceoperator_%s.yaml"
	return strings.ReplaceAll(tmpl, "%s", version)
}

// Get retrieves a template as a string either from a URL or local path
// uri is a URI pointing to either a remote file (http or https) or a local file.
func Get(ctx context.Context, uri string) (string, error) {
	parsedURI, err := url.Parse(uri)
	if err != nil {
		return "", eris.Wrapf(err, "failed to parse URI %s", uri)
	}

	switch parsedURI.Scheme {
	case "": // Assume that this is a local file?
		return load(uri)
	case "http", "https":
		return download(ctx, uri)
	default:
		return "", eris.Errorf("unknown URI scheme %s", parsedURI.Scheme)
	}
}

// Apply injects the specified pattern into the template
func Apply(template string, crdPattern string) (string, error) {
	// This is a bit hacky, but it works for our simple purposes.
	newPattern := fmt.Sprintf("- --crd-pattern=%s", crdPattern)

	// TODO: Maybe fix the YAML to be an actual envsubst template?
	result := strings.Replace(template, "- --crd-pattern=", newPattern, 1)
	if template == result {
		return "", eris.New("failed to inject crd-pattern, was the template used from ASO v2.1.0 or greater?")
	}

	return result, nil
}

// load loads a file from the given path
func load(path string) (string, error) {
	result, err := os.ReadFile(path)
	if err != nil {
		return "", eris.Wrapf(err, "failed to read file %s", path)
	}

	return string(result), nil
}

// download downloads a file from the specified URL
func download(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", eris.Wrap(err, "failed to create request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", eris.Wrapf(err, "failed to GET file at %s", url)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", eris.Errorf("unexpected status code %d from %s", resp.StatusCode, url)
	}

	bldr := &strings.Builder{}
	written, err := io.Copy(bldr, resp.Body)
	if err != nil {
		return "", eris.Wrap(err, "failed to copy body")
	}
	if written == 0 {
		return "", eris.Errorf("file had no content")
	}

	return bldr.String(), nil
}
