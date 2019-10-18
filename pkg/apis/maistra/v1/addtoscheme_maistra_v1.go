package v1

import (
	"github.com/maistra/istio-operator/pkg/apis/maistra/v1"
	"github.com/openshift-knative/knative-openshift-ingress/pkg/apis"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	// Adds schema for maistra
	apis.AddToSchemes = append(apis.AddToSchemes, v1.SchemeBuilder.AddToScheme)
}
