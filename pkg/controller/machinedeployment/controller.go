
// Copyright Jetstack Ltd. See LICENSE for details.


package machinedeployment

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/jetstack/wing-api/pkg/apis/wing/v1alpha1"
	"github.com/jetstack/wing-api/pkg/controller/sharedinformers"
	listers "github.com/jetstack/wing-api/pkg/client/listers_generated/wing/v1alpha1"
)

// +controller:group=wing,version=v1alpha1,kind=MachineDeployment,resource=machinedeployments
type MachineDeploymentControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about MachineDeployment
	lister listers.MachineDeploymentLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *MachineDeploymentControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing machinedeployments labels
	c.lister = arguments.GetSharedInformers().Factory.Wing().V1alpha1().MachineDeployments().Lister()
}

// Reconcile handles enqueued messages
func (c *MachineDeploymentControllerImpl) Reconcile(u *v1alpha1.MachineDeployment) error {
	// Implement controller logic here
	log.Printf("Running reconcile MachineDeployment for %s\n", u.Name)
	return nil
}

func (c *MachineDeploymentControllerImpl) Get(namespace, name string) (*v1alpha1.MachineDeployment, error) {
	return c.lister.MachineDeployments(namespace).Get(name)
}
