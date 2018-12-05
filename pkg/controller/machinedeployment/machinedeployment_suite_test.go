
// Copyright Jetstack Ltd. See LICENSE for details.


package machinedeployment_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/rest"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/test"

	"github.com/jetstack/wing-api/pkg/apis"
	"github.com/jetstack/wing-api/pkg/client/clientset_generated/clientset"
	"github.com/jetstack/wing-api/pkg/openapi"
	"github.com/jetstack/wing-api/pkg/controller/sharedinformers"
	"github.com/jetstack/wing-api/pkg/controller/machinedeployment"
)

var testenv *test.TestEnvironment
var config *rest.Config
var cs *clientset.Clientset
var shutdown chan struct{}
var controller *machinedeployment.MachineDeploymentController
var si *sharedinformers.SharedInformers

func TestMachineDeployment(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "MachineDeployment Suite", []Reporter{test.NewlineReporter{}})
}

var _ = BeforeSuite(func() {
	testenv = test.NewTestEnvironment()
	config = testenv.Start(apis.GetAllApiBuilders(), openapi.GetOpenAPIDefinitions)
	cs = clientset.NewForConfigOrDie(config)

	shutdown = make(chan struct{})
	si = sharedinformers.NewSharedInformers(config, shutdown)
	controller = machinedeployment.NewMachineDeploymentController(config, si)
	controller.Run(shutdown)
})

var _ = AfterSuite(func() {
	close(shutdown)
	testenv.Stop()
})
