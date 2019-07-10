package resourcegroups

import (
	"context"
	"testing"

	helpers "Telstra.Dx.AzureOperator/helpers"
	resourcemanagerconfig "Telstra.Dx.AzureOperator/resourcemanager/config"
)

func TestCreatingResouceGroup(t *testing.T) {
	resourcemanagerconfig.LoadSettings()
	resourcegroupName := "t-resourcegroup-" + helpers.RandomString(10)
	resourcegroupLocation := "westus"
	var err error

	_, err = CreateGroup(context.Background(), resourcegroupName, resourcegroupLocation)
	if err != nil {
		t.Errorf("ERROR")
	}

}
