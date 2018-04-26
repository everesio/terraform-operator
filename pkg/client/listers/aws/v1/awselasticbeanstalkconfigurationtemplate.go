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

// AwsElasticBeanstalkConfigurationTemplateLister helps list AwsElasticBeanstalkConfigurationTemplates.
type AwsElasticBeanstalkConfigurationTemplateLister interface {
	// List lists all AwsElasticBeanstalkConfigurationTemplates in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsElasticBeanstalkConfigurationTemplate, err error)
	// AwsElasticBeanstalkConfigurationTemplates returns an object that can list and get AwsElasticBeanstalkConfigurationTemplates.
	AwsElasticBeanstalkConfigurationTemplates(namespace string) AwsElasticBeanstalkConfigurationTemplateNamespaceLister
	AwsElasticBeanstalkConfigurationTemplateListerExpansion
}

// awsElasticBeanstalkConfigurationTemplateLister implements the AwsElasticBeanstalkConfigurationTemplateLister interface.
type awsElasticBeanstalkConfigurationTemplateLister struct {
	indexer cache.Indexer
}

// NewAwsElasticBeanstalkConfigurationTemplateLister returns a new AwsElasticBeanstalkConfigurationTemplateLister.
func NewAwsElasticBeanstalkConfigurationTemplateLister(indexer cache.Indexer) AwsElasticBeanstalkConfigurationTemplateLister {
	return &awsElasticBeanstalkConfigurationTemplateLister{indexer: indexer}
}

// List lists all AwsElasticBeanstalkConfigurationTemplates in the indexer.
func (s *awsElasticBeanstalkConfigurationTemplateLister) List(selector labels.Selector) (ret []*v1.AwsElasticBeanstalkConfigurationTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsElasticBeanstalkConfigurationTemplate))
	})
	return ret, err
}

// AwsElasticBeanstalkConfigurationTemplates returns an object that can list and get AwsElasticBeanstalkConfigurationTemplates.
func (s *awsElasticBeanstalkConfigurationTemplateLister) AwsElasticBeanstalkConfigurationTemplates(namespace string) AwsElasticBeanstalkConfigurationTemplateNamespaceLister {
	return awsElasticBeanstalkConfigurationTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsElasticBeanstalkConfigurationTemplateNamespaceLister helps list and get AwsElasticBeanstalkConfigurationTemplates.
type AwsElasticBeanstalkConfigurationTemplateNamespaceLister interface {
	// List lists all AwsElasticBeanstalkConfigurationTemplates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsElasticBeanstalkConfigurationTemplate, err error)
	// Get retrieves the AwsElasticBeanstalkConfigurationTemplate from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsElasticBeanstalkConfigurationTemplate, error)
	AwsElasticBeanstalkConfigurationTemplateNamespaceListerExpansion
}

// awsElasticBeanstalkConfigurationTemplateNamespaceLister implements the AwsElasticBeanstalkConfigurationTemplateNamespaceLister
// interface.
type awsElasticBeanstalkConfigurationTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsElasticBeanstalkConfigurationTemplates in the indexer for a given namespace.
func (s awsElasticBeanstalkConfigurationTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsElasticBeanstalkConfigurationTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsElasticBeanstalkConfigurationTemplate))
	})
	return ret, err
}

// Get retrieves the AwsElasticBeanstalkConfigurationTemplate from the indexer for a given namespace and name.
func (s awsElasticBeanstalkConfigurationTemplateNamespaceLister) Get(name string) (*v1.AwsElasticBeanstalkConfigurationTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awselasticbeanstalkconfigurationtemplate"), name)
	}
	return obj.(*v1.AwsElasticBeanstalkConfigurationTemplate), nil
}