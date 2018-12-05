// Copyright Jetstack Ltd. See LICENSE for details.
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

func SetDefaults_Environment(obj *MachineSetSpec) {
	if obj.Replicas == nil {
		var replicas int32 = 1
		obj.Replicas = &replicas
	}
}
