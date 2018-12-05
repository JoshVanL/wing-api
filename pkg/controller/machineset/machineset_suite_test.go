
// Copyright Jetstack Ltd. See LICENSE for details.


package machineset_test

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
	"github.com/jetstack/wing-api/pkg/controller/machineset"
)

var testenv *test.TestEnvironment
var config *rest.Config
var cs *clientset.Clientset
var shutdown chan struct{}
var controller *machineset.MachineSetController
var si *sharedinformers.SharedInformers

func TestMachineSet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "MachineSet Suite", []Reporter{test.NewlineReporter{}})
}

var _ = BeforeSuite(func() {
	testenv = test.NewTestEnvironment()
	config = testenv.Start(apis.GetAllApiBuilders(), openapi.GetOpenAPIDefinitions)
	cs = clientset.NewForConfigOrDie(config)

	shutdown = make(chan struct{})
	si = sharedinformers.NewSharedInformers(config, shutdown)
	controller = machineset.NewMachineSetController(config, si)
	controller.Run(shutdown)
})

var _ = AfterSuite(func() {
	close(shutdown)
	testenv.Stop()
})
