/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/
package v1beta1

import (
	"os"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/arbitrary"
	"github.com/leanovate/gopter/gen"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1"
)

func Test_User_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 50
	arbitraries := arbitrary.DefaultArbitraries()
	arbitraries.RegisterGen(TypeMetaGenerator())

	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from User to hub returns original",
		arbitraries.ForAll(RunResourceConversionTestForUser))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForUser tests if a specific instance of ConfigurationStore round trips to the hub storage version and back losslessly
func RunResourceConversionTestForUser(subject User) (bool, error) {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1.User
	err := copied.ConvertTo(&hub)
	if err != nil {
		return false, err
	}

	// Convert from our hub version
	var actual User
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return false, err
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return false, errors.New(result)
	}

	return true, nil
}

func TypeMetaGenerator() gopter.Gen {
	generators := make(map[string]gopter.Gen)
	typeMetaGenerator := gen.Struct(reflect.TypeOf(metav1.TypeMeta{}), generators)

	return typeMetaGenerator
}
