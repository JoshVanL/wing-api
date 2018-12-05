
// Copyright Jetstack Ltd. See LICENSE for details.


package machine

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/jetstack/wing-api/pkg/apis/wing/v1alpha1"
	"github.com/jetstack/wing-api/pkg/controller/sharedinformers"
	listers "github.com/jetstack/wing-api/pkg/client/listers_generated/wing/v1alpha1"
)

// +controller:group=wing,version=v1alpha1,kind=Machine,resource=machines
type MachineControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Machine
	lister listers.MachineLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *MachineControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing machines labels
	c.lister = arguments.GetSharedInformers().Factory.Wing().V1alpha1().Machines().Lister()
}

// Reconcile handles enqueued messages
func (c *MachineControllerImpl) Reconcile(u *v1alpha1.Machine) error {
	// Implement controller logic here
	log.Printf("Running reconcile Machine for %s\n", u.Name)
	return nil
}

func (c *MachineControllerImpl) Get(namespace, name string) (*v1alpha1.Machine, error) {
	return c.lister.Machines(namespace).Get(name)
}
