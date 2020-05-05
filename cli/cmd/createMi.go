/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"encoding/json"
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

		// interpolate params into command
		c := fmt.Sprintf(
			"az identity create -g %s -n %s --subscription %s -o json",
			config.ResourceGroup,
			config.ManagedIdentity,
			config.Subscription,
		)

		out, err := RunCommand(c)
		if err != nil {
			return err
		}

		err = json.Unmarshal(out, &config)
		if err != nil {
			return err
		}

		c2 := fmt.Sprintf(
			`az role assignment create --role "Managed Identity Operator" --assignee %s --scope "%s"`,
			config.ServicePrincipal,
			config.ManagedIdentityID,
		)

		out, err = RunCommand(c2)
		if err != nil {
			return err
		}

		for {
			c3 := fmt.Sprintf(
				`az role assignment create --role "Contributor" --assignee %s --scope /subscriptions/%s`,
				config.ManagedIdentityClientID,
				config.Subscription,
			)

			out, err = RunCommand(c3)
			if err != nil {
				if strings.Contains(string(out), "No matches in graph database for") || strings.Contains(string(out), "does not exist in the directory") {
					time.Sleep(10 * time.Second)
					continue
				}
				return err
			}

			break
		}

		fmt.Println()
		fmt.Println("Install AAD Pod Identity:")
		fmt.Println("make install-aad-pod-identity")
		fmt.Println()

		tpl := `
apiVersion: "aadpodidentity.k8s.io/v1"
kind: AzureIdentity
metadata:
	name: {{ .ManagedIdentity }}
	namespace: azureoperator-system
spec:
	type: 0
	ResourceID: /subscriptions/{{ .Subscription }}/resourcegroups/{{ .ResourceGroup }}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{{ .ManagedIdentity }}
	ClientID: {{ .ManagedIdentityClientID }}
---
apiVersion: "aadpodidentity.k8s.io/v1"
kind: AzureIdentityBinding
metadata:
	name: aso-identity-binding
	namespace: azureoperator-system
spec:
	AzureIdentity: {{ .ManagedIdentity }}
	Selector: aso_manager_binding`
		fmt.Println()

		t := template.Must(template.New("manifests").Parse(tpl))
		err = t.Execute(os.Stdout, config)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createMiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	createMiCmd.PersistentFlags().StringVarP(&config.ResourceGroup, "resource-group", "g", "", "resource group to associate managed identity with")
	createMiCmd.PersistentFlags().StringVarP(&config.ManagedIdentity, "managed-identity", "i", "", "managed identity name to create and set up")
	createMiCmd.PersistentFlags().StringVarP(&config.ServicePrincipal, "service-principal", "p", "", "service principal associated with kube cluster cluster")
	createMiCmd.PersistentFlags().StringVarP(&config.Subscription, "subscription", "s", "", "azure subscription to act against")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createMiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func RunCommand(c string) ([]byte, error) {
	fmt.Println(c)
	parts, err := shlex.Split(c)
	if err != nil {
		return []byte(""), err
	}

	stmt := exec.Command(parts[0], parts[1:]...)
	out, err := stmt.CombinedOutput()
	fmt.Println(string(out))
	return out, err
}
