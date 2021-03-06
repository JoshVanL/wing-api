/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package controller

import (
	"github.com/jetstack/wing-api/pkg/controller/machine"
	"github.com/jetstack/wing-api/pkg/controller/machinedeployment"
	"github.com/jetstack/wing-api/pkg/controller/machineset"
	"github.com/jetstack/wing-api/pkg/controller/sharedinformers"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/controller"
	"k8s.io/client-go/rest"
)

func GetAllControllers(config *rest.Config) ([]controller.Controller, chan struct{}) {
	shutdown := make(chan struct{})
	si := sharedinformers.NewSharedInformers(config, shutdown)
	return []controller.Controller{
		machine.NewMachineController(config, si),
		machinedeployment.NewMachineDeploymentController(config, si),
		machineset.NewMachineSetController(config, si),
	}, shutdown
}
