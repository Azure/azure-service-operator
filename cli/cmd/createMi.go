// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"os/exec"

	"github.com/google/shlex"
	"github.com/spf13/cobra"
)

type Config struct {
	ResourceGroup           string
	ManagedIdentity         string
	ServicePrincipal        string
	Subscription            string
	ManagedIdentityClientID string `json:"clientId"`
	ManagedIdentityID       string `json:"id"`
}

var config Config

// createMiCmd represents the createMi command
var createMiCmd = &cobra.Command{
	Use:     "createMi",
	Aliases: []string{"createmi", "createMI"},
	Short:   "Set up a new Managed Identity for use with the Azure Service Operator",
	Long: `Set up a new Managed Identity for use with the Azure Service Operator:

Creates a new MI, assigns necessary roles, displays helpful follow-up commands.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		commands := []struct {
			Name   string
			C      string
			Target interface{}
		}{
			{
				Name:   "Create Managed Identity",
				C:      "az identity create -g {{ .ResourceGroup }} -n {{ .ManagedIdentity }} --subscription {{ .Subscription }} -o json",
				Target: &config,
			},
			{
				Name:   "Assign Managed Identity Operator Role",
				C:      `az role assignment create --role "Managed Identity Operator" --assignee {{ .ServicePrincipal }} --scope "{{ .ManagedIdentityID }}"`,
				Target: nil,
			},
			{
				Name:   "Assign Managed Identity Contributor Role",
				C:      `az role assignment create --role "Contributor" --assignee {{ .ManagedIdentityClientID }} --scope /subscriptions/{{ .Subscription }}`,
				Target: nil,
			},
		}

		for _, c := range commands {
			fmt.Println("Starting action:", c.Name)
			fmt.Println("-------------------------")
			t := template.Must(template.New(c.Name).Parse(c.C))

			var rendered bytes.Buffer
			err := t.Execute(&rendered, config)
			if err != nil {
				return err
			}

			fmt.Println(rendered.String())
			fmt.Println()
			fmt.Println()

			for {
				out, err := RunCommand(rendered.String())
				fmt.Println(string(out))
				if err != nil {
					if strings.Contains(string(out), "No matches in graph database") || strings.Contains(string(out), "does not exist in the directory") || strings.Contains(string(out), "Cannot find user or service principal in graph database") {
						time.Sleep(10 * time.Second)
						continue
					}
					return err
				}

				if c.Target != nil {
					err = json.Unmarshal(out, c.Target)
					if err != nil {
						return err
					}
				}

				break
			}

		}

		fmt.Println()
		fmt.Println("Install AAD Pod Identity:")
		fmt.Println("make install-aad-pod-identity")
		fmt.Println()

		tpl := `apiVersion: "aadpodidentity.k8s.io/v1"
kind: AzureIdentity
metadata:
  name: {{ .ManagedIdentity }}
  namespace: azureoperator-system
spec:
  type: 0
  resourceID: /subscriptions/{{ .Subscription }}/resourcegroups/{{ .ResourceGroup }}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{{ .ManagedIdentity }}
  clientID: {{ .ManagedIdentityClientID }}
---
apiVersion: "aadpodidentity.k8s.io/v1"
kind: AzureIdentityBinding
metadata:
  name: aso-identity-binding
  namespace: azureoperator-system
spec:
  azureIdentity: {{ .ManagedIdentity }}
  selector: aso_manager_binding
`

		fmt.Println("cat <<EOF | kubectl apply -f -")
		t := template.Must(template.New("manifests").Parse(tpl))
		err := t.Execute(os.Stdout, config)
		if err != nil {
			return err
		}
		fmt.Println("EOF")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createMiCmd)
	createMiCmd.PersistentFlags().StringVarP(&config.ResourceGroup, "resource-group", "g", "", "resource group to associate managed identity with")
	createMiCmd.PersistentFlags().StringVarP(&config.ManagedIdentity, "managed-identity", "i", "", "managed identity name to create and set up")
	createMiCmd.PersistentFlags().StringVarP(&config.ServicePrincipal, "service-principal", "p", "", "service principal associated with kube cluster cluster")
	createMiCmd.PersistentFlags().StringVarP(&config.Subscription, "subscription", "s", "", "azure subscription to act against")
}

func RunCommand(c string) ([]byte, error) {
	parts, err := shlex.Split(c)
	if err != nil {
		return []byte(""), err
	}

	stmt := exec.Command(parts[0], parts[1:]...)
	out, err := stmt.CombinedOutput()
	return out, err
}
