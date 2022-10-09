/*
Copyright The Kubectl Bind contributors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
)

// ServiceBindingLister helps list ServiceBindings.
// All objects returned here must be treated as read-only.
type ServiceBindingLister interface {
	// List lists all ServiceBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceBinding, err error)
	// Get retrieves the ServiceBinding from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServiceBinding, error)
	ServiceBindingListerExpansion
}

// serviceBindingLister implements the ServiceBindingLister interface.
type serviceBindingLister struct {
	indexer cache.Indexer
}

// NewServiceBindingLister returns a new ServiceBindingLister.
func NewServiceBindingLister(indexer cache.Indexer) ServiceBindingLister {
	return &serviceBindingLister{indexer: indexer}
}

// List lists all ServiceBindings in the indexer.
func (s *serviceBindingLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceBinding))
	})
	return ret, err
}

// Get retrieves the ServiceBinding from the index for a given name.
func (s *serviceBindingLister) Get(name string) (*v1alpha1.ServiceBinding, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicebinding"), name)
	}
	return obj.(*v1alpha1.ServiceBinding), nil
}
