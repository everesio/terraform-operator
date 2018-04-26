/*
Copyright The Kubernetes Authors.

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

package v1

import (
	v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AwsRoute53ZoneLister helps list AwsRoute53Zones.
type AwsRoute53ZoneLister interface {
	// List lists all AwsRoute53Zones in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsRoute53Zone, err error)
	// AwsRoute53Zones returns an object that can list and get AwsRoute53Zones.
	AwsRoute53Zones(namespace string) AwsRoute53ZoneNamespaceLister
	AwsRoute53ZoneListerExpansion
}

// awsRoute53ZoneLister implements the AwsRoute53ZoneLister interface.
type awsRoute53ZoneLister struct {
	indexer cache.Indexer
}

// NewAwsRoute53ZoneLister returns a new AwsRoute53ZoneLister.
func NewAwsRoute53ZoneLister(indexer cache.Indexer) AwsRoute53ZoneLister {
	return &awsRoute53ZoneLister{indexer: indexer}
}

// List lists all AwsRoute53Zones in the indexer.
func (s *awsRoute53ZoneLister) List(selector labels.Selector) (ret []*v1.AwsRoute53Zone, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsRoute53Zone))
	})
	return ret, err
}

// AwsRoute53Zones returns an object that can list and get AwsRoute53Zones.
func (s *awsRoute53ZoneLister) AwsRoute53Zones(namespace string) AwsRoute53ZoneNamespaceLister {
	return awsRoute53ZoneNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsRoute53ZoneNamespaceLister helps list and get AwsRoute53Zones.
type AwsRoute53ZoneNamespaceLister interface {
	// List lists all AwsRoute53Zones in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsRoute53Zone, err error)
	// Get retrieves the AwsRoute53Zone from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsRoute53Zone, error)
	AwsRoute53ZoneNamespaceListerExpansion
}

// awsRoute53ZoneNamespaceLister implements the AwsRoute53ZoneNamespaceLister
// interface.
type awsRoute53ZoneNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsRoute53Zones in the indexer for a given namespace.
func (s awsRoute53ZoneNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsRoute53Zone, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsRoute53Zone))
	})
	return ret, err
}

// Get retrieves the AwsRoute53Zone from the indexer for a given namespace and name.
func (s awsRoute53ZoneNamespaceLister) Get(name string) (*v1.AwsRoute53Zone, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsroute53zone"), name)
	}
	return obj.(*v1.AwsRoute53Zone), nil
}