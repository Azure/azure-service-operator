// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"fmt"
	"time"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PSQL server", func() {

	var rgName string
	var location string
	var mysqlServer string
	var MySQLManager MySQLServerManager

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.ResourceGroupName
		location = tc.ResourceGroupLocation
		MySQLManager = tc.mySQLServerManager
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {
		It("should create and delete PSQL server in azure", func() {

			defer GinkgoRecover()

			mysqlServer = "t-dev-psql-srv-" + helpers.RandomString(10)

			mySQLSku := mysql.Sku{
				Name:     to.StringPtr("B_Gen5_2"),
				Tier:     mysql.SkuTier("Basic"),
				Capacity: to.Int32Ptr(2),
				Size:     to.StringPtr("51200"),
				Family:   to.StringPtr("Gen5"),
			}
			tags := map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			}

			// Create PSQL server instance
			Eventually(func() bool {
				time.Sleep(3 * time.Second)
				_, err := MySQLManager.GetServer(ctx, rgName, mysqlServer)
				if err == nil {
					return true
				}
				_, err = MySQLManager.CreateServerIfValid(
					ctx,
					mysqlServer,
					rgName,
					location,
					tags,
					mysql.EightFullStopZero,
					mysql.SslEnforcementEnumEnabled,
					mySQLSku,
					"adm1nus3r",
					"m@#terU$3r",
				)
				if err != nil {
					fmt.Println(err.Error())
					if !errhelp.IsAsynchronousOperationNotComplete(err) {
						fmt.Println("error occured")
						return false
					}
				}
				return true
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())

			time.Sleep(5 * time.Minute)

			// Delete PSQL server instance
			Eventually(func() bool {
				time.Sleep(3 * time.Second)
				_, err := MySQLManager.GetServer(ctx, rgName, mysqlServer)
				if err != nil {
					return true
				}
				_, err = MySQLManager.DeleteServer(ctx, mysqlServer, rgName)
				if err != nil {
					fmt.Println(err.Error())
					if !errhelp.IsAsynchronousOperationNotComplete(err) {
						fmt.Println("error occured")
						return false
					}
				}
				return err == nil
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())
		})

	})
})
