package sqlclient

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	"github.com/Azure/azure-service-operator/pkg/util"
	"github.com/marstr/randname"
)

var tc TestContext
var rgm resourcegroups.AzureResourceGroupManager
var groupName string
var ctx context.Context
var location string
var err error

func addLocalEnvAndParse() error {
	// parse env at top-level (also controls dotenv load)
	err := config.ParseEnvironment()
	if err != nil {
		return fmt.Errorf("failed to add top-level env: %v", err.Error())
	}
	return nil
}

func setup() error {
	err = addLocalEnvAndParse()
	if err != nil {
		return err
	}

	rgm = resourcegroups.AzureResourceGroupManager{}
	location = config.DefaultLocation()

	groupName = config.GenerateGroupName("SQLCreateDeleteTest")
	config.SetGroupName(groupName)

	ctx = context.Background()

	// create the resource group
	_, err = rgm.CreateGroup(ctx, groupName, location)
	if err != nil {
		util.PrintAndLog(err.Error())
	}
	util.PrintAndLog("Setup:Created resource group " + groupName)

	return nil
}

func teardown() error {
	if !config.KeepResources() {
		// does not wait
		// delete the resource group
		util.PrintAndLog("deleting resource group...")
		time.Sleep(time.Second)
		_, err = rgm.DeleteGroup(ctx, groupName)
		if !errhelp.IsAsynchronousOperationNotComplete(err) {
			util.PrintAndLog(fmt.Sprintf("cannot delete resource group: %v", err))
			return err
		}
		util.PrintAndLog("waiting for resource group delete future to come back")

		then := time.Now()
		for {
			if time.Since(then) > tc.timeout {
				util.PrintAndLog("test timed out")
				return fmt.Errorf("Timed out trying to clean up RG")
			}
			time.Sleep(time.Second * 10)
			_, err := resourcegroups.GetGroup(ctx, groupName)
			if err == nil {
				util.PrintAndLog("waiting for resource group to be deleted")
			} else {
				if errhelp.IsGroupNotFound(err) {
					util.PrintAndLog("resource group deleted")
					break
				} else {
					util.PrintAndLog(fmt.Sprintf("cannot delete resource group: %v", err))
					return err
				}
			}
		}
	}
	return nil
}

// test helpers
func generateName(prefix string) string {
	return strings.ToLower(randname.GenerateWithPrefix(prefix, 5))
}

// TestMain is the main entry point for tests
func TestMain(m *testing.M) {
	var err error
	var code int

	err = setup()
	if err != nil {
		util.PrintAndLog(fmt.Sprintf("could not set up environment: %v\n", err))
	}

	code = m.Run()

	err = teardown()
	if err != nil {
		util.PrintAndLog(fmt.Sprintf("could not tear down environment: %v\n; original exit code: %v\n", err, code))
	}

	os.Exit(code)
}
