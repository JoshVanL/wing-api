// Copyright Jetstack Ltd. See LICENSE for details.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/jetstack/wing-api/pkg/apis/wing/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MachineSetLister helps list MachineSets.
type MachineSetLister interface {
	// List lists all MachineSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MachineSet, err error)
	// MachineSets returns an object that can list and get MachineSets.
	MachineSets(namespace string) MachineSetNamespaceLister
	MachineSetListerExpansion
}

// machineSetLister implements the MachineSetLister interface.
type machineSetLister struct {
	indexer cache.Indexer
}

// NewMachineSetLister returns a new MachineSetLister.
func NewMachineSetLister(indexer cache.Indexer) MachineSetLister {
	return &machineSetLister{indexer: indexer}
}

// List lists all MachineSets in the indexer.
func (s *machineSetLister) List(selector labels.Selector) (ret []*v1alpha1.MachineSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MachineSet))
	})
	return ret, err
}

// MachineSets returns an object that can list and get MachineSets.
func (s *machineSetLister) MachineSets(namespace string) MachineSetNamespaceLister {
	return machineSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachineSetNamespaceLister helps list and get MachineSets.
type MachineSetNamespaceLister interface {
	// List lists all MachineSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MachineSet, err error)
	// Get retrieves the MachineSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MachineSet, error)
	MachineSetNamespaceListerExpansion
}

// machineSetNamespaceLister implements the MachineSetNamespaceLister
// interface.
type machineSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MachineSets in the indexer for a given namespace.
func (s machineSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MachineSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MachineSet))
	})
	return ret, err
}

// Get retrieves the MachineSet from the indexer for a given namespace and name.
func (s machineSetNamespaceLister) Get(name string) (*v1alpha1.MachineSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("machineset"), name)
	}
	return obj.(*v1alpha1.MachineSet), nil
}
