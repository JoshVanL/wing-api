
// Copyright Jetstack Ltd. See LICENSE for details.


package v1alpha1

import (
	"log"
	"context"

	"k8s.io/apimachinery/pkg/runtime"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/jetstack/tarmak/pkg/apis/wing"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Machine
// +k8s:openapi-gen=true
// +resource:path=machines,strategy=MachineStrategy
type Machine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MachineSpec   `json:"spec,omitempty"`
	Status MachineStatus `json:"status,omitempty"`
}

// MachineSpec defines the desired state of Machine
type MachineSpec struct {
}

// MachineStatus defines the observed state of Machine
type MachineStatus struct {
}

// Validate checks that an instance of Machine is well formed
func (MachineStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*wing.Machine)
	log.Printf("Validating fields for Machine %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default Machine field values
func (MachineSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*Machine)
	// set default field values here
	log.Printf("Defaulting fields for Machine %s\n", obj.Name)
}
