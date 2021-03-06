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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	scheme "github.com/trussle/terraform-operator/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AwsWafRulesGetter has a method to return a AwsWafRuleInterface.
// A group's client should implement this interface.
type AwsWafRulesGetter interface {
	AwsWafRules(namespace string) AwsWafRuleInterface
}

// AwsWafRuleInterface has methods to work with AwsWafRule resources.
type AwsWafRuleInterface interface {
	Create(*v1.AwsWafRule) (*v1.AwsWafRule, error)
	Update(*v1.AwsWafRule) (*v1.AwsWafRule, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsWafRule, error)
	List(opts meta_v1.ListOptions) (*v1.AwsWafRuleList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsWafRule, err error)
	AwsWafRuleExpansion
}

// awsWafRules implements AwsWafRuleInterface
type awsWafRules struct {
	client rest.Interface
	ns     string
}

// newAwsWafRules returns a AwsWafRules
func newAwsWafRules(c *TrussleV1Client, namespace string) *awsWafRules {
	return &awsWafRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsWafRule, and returns the corresponding awsWafRule object, and an error if there is any.
func (c *awsWafRules) Get(name string, options meta_v1.GetOptions) (result *v1.AwsWafRule, err error) {
	result = &v1.AwsWafRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafRules that match those selectors.
func (c *awsWafRules) List(opts meta_v1.ListOptions) (result *v1.AwsWafRuleList, err error) {
	result = &v1.AwsWafRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafRules.
func (c *awsWafRules) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awswafrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsWafRule and creates it.  Returns the server's representation of the awsWafRule, and an error, if there is any.
func (c *awsWafRules) Create(awsWafRule *v1.AwsWafRule) (result *v1.AwsWafRule, err error) {
	result = &v1.AwsWafRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awswafrules").
		Body(awsWafRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafRule and updates it. Returns the server's representation of the awsWafRule, and an error, if there is any.
func (c *awsWafRules) Update(awsWafRule *v1.AwsWafRule) (result *v1.AwsWafRule, err error) {
	result = &v1.AwsWafRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awswafrules").
		Name(awsWafRule.Name).
		Body(awsWafRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafRule and deletes it. Returns an error if one occurs.
func (c *awsWafRules) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafRules) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafRule.
func (c *awsWafRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsWafRule, err error) {
	result = &v1.AwsWafRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awswafrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
