/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package arm

import (
	"strconv"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const (
	PollerResumeTokenAnnotation = "serviceoperator.azure.com/poller-resume-token"
	PollerResumeIDAnnotation    = "serviceoperator.azure.com/poller-resume-id"
	LatestReconciledGeneration  = "serviceoperator.azure.com/latest-reconciled-generation"
)

// GetPollerResumeToken returns a poller ID and the poller token
func GetPollerResumeToken(obj genruntime.MetaObject) (string, string, bool) {
	token, hasResumeToken := obj.GetAnnotations()[PollerResumeTokenAnnotation]
	id, hasResumeID := obj.GetAnnotations()[PollerResumeIDAnnotation]

	return id, token, hasResumeToken && hasResumeID
}

func SetPollerResumeToken(obj genruntime.MetaObject, id string, token string) {
	genruntime.AddAnnotation(obj, PollerResumeTokenAnnotation, token)
	genruntime.AddAnnotation(obj, PollerResumeIDAnnotation, id)
}

// ClearPollerResumeToken clears the poller resume token and ID annotations
func ClearPollerResumeToken(obj genruntime.MetaObject) {
	genruntime.RemoveAnnotation(obj, PollerResumeTokenAnnotation)
	genruntime.RemoveAnnotation(obj, PollerResumeIDAnnotation)
}

func SetLatestReconciledGeneration(obj genruntime.MetaObject) {
	genruntime.AddAnnotation(obj, LatestReconciledGeneration, strconv.FormatInt(obj.GetGeneration(), 10))
}

func GetLatestReconciledGeneration(obj genruntime.MetaObject) (int64, bool) {
	val, hasGeneration := obj.GetAnnotations()[LatestReconciledGeneration]
	gen, err := strconv.Atoi(val)
	if err != nil {
		return 0, false
	}
	return int64(gen), hasGeneration
}
