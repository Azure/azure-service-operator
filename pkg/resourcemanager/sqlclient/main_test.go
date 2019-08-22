package sqlclient

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/resources"
	"github.com/Azure/azure-service-operator/pkg/util"
	"github.com/marstr/randname"
)

func addLocalEnvAndParse() error {
	// parse env at top-level (also controls dotenv load)
	err := config.ParseEnvironment()
	if err != nil {
		return fmt.Errorf("failed to add top-level env: %v", err.Error())
	}
	return nil
}

func setup() error {
	var err error
	err = addLocalEnvAndParse()
	if err != nil {
		return err
	}

	return nil
}

func teardown() error {
	if !config.KeepResources() {
		// does not wait
		_, err := resources.DeleteGroup(context.Background(), config.GroupName())
		if err != nil {

			// this indicates that the resource group wasn't created
			// no worries!
			util.PrintAndLog(fmt.Sprintf("could not teardown resource group, this may not be a problem: %v\n", err))
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
