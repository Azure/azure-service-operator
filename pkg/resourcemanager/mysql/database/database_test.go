// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package database

import (
	"fmt"
	"time"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	server "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/server"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PSQL database", func() {

	var rgName string
	var location string
	var psqlServer string
	var PSQLServerManager server.PostgreSQLServerManager
	var PSQLDatabaseManager PostgreSQLDatabaseManager

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.ResourceGroupName
		location = tc.ResourceGroupLocation
		PSQLServerManager = tc.postgreSQLServerManager
		PSQLDatabaseManager = tc.postgreSQLDatabaseManager

		// Create PSQL server
		psqlServer = "t-dev-psql-srv-" + helpers.RandomString(10)

		pSQLSku := psql.Sku{
			Name:     to.StringPtr("B_Gen5_2"),
			Tier:     psql.SkuTier("Basic"),
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
			_, err := PSQLServerManager.GetServer(ctx, rgName, psqlServer)
			if err == nil {
				return true
			}
			_, err = PSQLServerManager.CreateServerIfValid(
				ctx,
				psqlServer,
				rgName,
				location,
				tags,
				psql.OneZero,
				psql.SslEnforcementEnumEnabled,
				pSQLSku,
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

	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
		// Delete PSQL server
		Eventually(func() bool {
			time.Sleep(3 * time.Second)
			_, err := PSQLServerManager.GetServer(ctx, rgName, psqlServer)
			if err != nil {
				return true
			}
			_, err = PSQLServerManager.DeleteServer(ctx, psqlServer, rgName)
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

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {
		It("should create and delete PSQL database in azure", func() {

			defer GinkgoRecover()

			psqlDatabase := "t-dev-psql-db-" + helpers.RandomString(10)

			// Create PSQL database instance
			Eventually(func() bool {
				time.Sleep(3 * time.Second)
				_, err := PSQLDatabaseManager.GetDatabase(ctx, rgName, psqlServer, psqlDatabase)
				if err == nil {
					return true
				}
				_, err = PSQLDatabaseManager.CreateDatabaseIfValid(ctx, psqlDatabase, psqlServer, rgName)
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

			// Delete PSQL database instance
			Eventually(func() bool {
				time.Sleep(3 * time.Second)
				_, err := PSQLDatabaseManager.GetDatabase(ctx, rgName, psqlServer, psqlDatabase)
				if err != nil {
					return true
				}
				_, err = PSQLDatabaseManager.DeleteDatabase(ctx, psqlDatabase, psqlServer, rgName)
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
