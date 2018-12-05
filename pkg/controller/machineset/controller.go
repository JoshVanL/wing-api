
// Copyright Jetstack Ltd. See LICENSE for details.


package machineset

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/jetstack/wing-api/pkg/apis/wing/v1alpha1"
	"github.com/jetstack/wing-api/pkg/controller/sharedinformers"
	listers "github.com/jetstack/wing-api/pkg/client/listers_generated/wing/v1alpha1"
)

// +controller:group=wing,version=v1alpha1,kind=MachineSet,resource=machinesets
type MachineSetControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about MachineSet
	lister listers.MachineSetLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *MachineSetControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing machinesets labels
	c.lister = arguments.GetSharedInformers().Factory.Wing().V1alpha1().MachineSets().Lister()
}

// Reconcile handles enqueued messages
func (c *MachineSetControllerImpl) Reconcile(u *v1alpha1.MachineSet) error {
	// Implement controller logic here
	log.Printf("Running reconcile MachineSet for %s\n", u.Name)
	return nil
}

func (c *MachineSetControllerImpl) Get(namespace, name string) (*v1alpha1.MachineSet, error) {
	return c.lister.MachineSets(namespace).Get(name)
}
