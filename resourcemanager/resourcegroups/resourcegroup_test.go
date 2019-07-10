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

	// flag := true
	// for {
	// 	result, err := CheckExistence(context.Background(), resourcegroupName)
	// 	if err != nil {
	// 		flag = false
	// 		t.Errorf("ERROR")

	// 	}
	// 	if result.Response != nil {

	// 	}
	// 	time.Sleep(1 * time.Second)
	// 	if !flag {
	// 		break
	// 	}
	// }

}
