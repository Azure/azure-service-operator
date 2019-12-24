package sqlclient

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
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

	os.Exit(code)
}
