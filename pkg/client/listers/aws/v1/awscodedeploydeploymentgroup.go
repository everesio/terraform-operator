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

// AwsCodedeployDeploymentGroupLister helps list AwsCodedeployDeploymentGroups.
type AwsCodedeployDeploymentGroupLister interface {
	// List lists all AwsCodedeployDeploymentGroups in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsCodedeployDeploymentGroup, err error)
	// AwsCodedeployDeploymentGroups returns an object that can list and get AwsCodedeployDeploymentGroups.
	AwsCodedeployDeploymentGroups(namespace string) AwsCodedeployDeploymentGroupNamespaceLister
	AwsCodedeployDeploymentGroupListerExpansion
}

// awsCodedeployDeploymentGroupLister implements the AwsCodedeployDeploymentGroupLister interface.
type awsCodedeployDeploymentGroupLister struct {
	indexer cache.Indexer
}

// NewAwsCodedeployDeploymentGroupLister returns a new AwsCodedeployDeploymentGroupLister.
func NewAwsCodedeployDeploymentGroupLister(indexer cache.Indexer) AwsCodedeployDeploymentGroupLister {
	return &awsCodedeployDeploymentGroupLister{indexer: indexer}
}

// List lists all AwsCodedeployDeploymentGroups in the indexer.
func (s *awsCodedeployDeploymentGroupLister) List(selector labels.Selector) (ret []*v1.AwsCodedeployDeploymentGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsCodedeployDeploymentGroup))
	})
	return ret, err
}

// AwsCodedeployDeploymentGroups returns an object that can list and get AwsCodedeployDeploymentGroups.
func (s *awsCodedeployDeploymentGroupLister) AwsCodedeployDeploymentGroups(namespace string) AwsCodedeployDeploymentGroupNamespaceLister {
	return awsCodedeployDeploymentGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsCodedeployDeploymentGroupNamespaceLister helps list and get AwsCodedeployDeploymentGroups.
type AwsCodedeployDeploymentGroupNamespaceLister interface {
	// List lists all AwsCodedeployDeploymentGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsCodedeployDeploymentGroup, err error)
	// Get retrieves the AwsCodedeployDeploymentGroup from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsCodedeployDeploymentGroup, error)
	AwsCodedeployDeploymentGroupNamespaceListerExpansion
}

// awsCodedeployDeploymentGroupNamespaceLister implements the AwsCodedeployDeploymentGroupNamespaceLister
// interface.
type awsCodedeployDeploymentGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsCodedeployDeploymentGroups in the indexer for a given namespace.
func (s awsCodedeployDeploymentGroupNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsCodedeployDeploymentGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsCodedeployDeploymentGroup))
	})
	return ret, err
}

// Get retrieves the AwsCodedeployDeploymentGroup from the indexer for a given namespace and name.
func (s awsCodedeployDeploymentGroupNamespaceLister) Get(name string) (*v1.AwsCodedeployDeploymentGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awscodedeploydeploymentgroup"), name)
	}
	return obj.(*v1.AwsCodedeployDeploymentGroup), nil
}
