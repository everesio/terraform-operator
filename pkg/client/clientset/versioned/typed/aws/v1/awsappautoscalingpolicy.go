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

// AwsAppautoscalingPoliciesGetter has a method to return a AwsAppautoscalingPolicyInterface.
// A group's client should implement this interface.
type AwsAppautoscalingPoliciesGetter interface {
	AwsAppautoscalingPolicies(namespace string) AwsAppautoscalingPolicyInterface
}

// AwsAppautoscalingPolicyInterface has methods to work with AwsAppautoscalingPolicy resources.
type AwsAppautoscalingPolicyInterface interface {
	Create(*v1.AwsAppautoscalingPolicy) (*v1.AwsAppautoscalingPolicy, error)
	Update(*v1.AwsAppautoscalingPolicy) (*v1.AwsAppautoscalingPolicy, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsAppautoscalingPolicy, error)
	List(opts meta_v1.ListOptions) (*v1.AwsAppautoscalingPolicyList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsAppautoscalingPolicy, err error)
	AwsAppautoscalingPolicyExpansion
}

// awsAppautoscalingPolicies implements AwsAppautoscalingPolicyInterface
type awsAppautoscalingPolicies struct {
	client rest.Interface
	ns     string
}

// newAwsAppautoscalingPolicies returns a AwsAppautoscalingPolicies
func newAwsAppautoscalingPolicies(c *TrussleV1Client, namespace string) *awsAppautoscalingPolicies {
	return &awsAppautoscalingPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsAppautoscalingPolicy, and returns the corresponding awsAppautoscalingPolicy object, and an error if there is any.
func (c *awsAppautoscalingPolicies) Get(name string, options meta_v1.GetOptions) (result *v1.AwsAppautoscalingPolicy, err error) {
	result = &v1.AwsAppautoscalingPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsappautoscalingpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAppautoscalingPolicies that match those selectors.
func (c *awsAppautoscalingPolicies) List(opts meta_v1.ListOptions) (result *v1.AwsAppautoscalingPolicyList, err error) {
	result = &v1.AwsAppautoscalingPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsappautoscalingpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAppautoscalingPolicies.
func (c *awsAppautoscalingPolicies) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsappautoscalingpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsAppautoscalingPolicy and creates it.  Returns the server's representation of the awsAppautoscalingPolicy, and an error, if there is any.
func (c *awsAppautoscalingPolicies) Create(awsAppautoscalingPolicy *v1.AwsAppautoscalingPolicy) (result *v1.AwsAppautoscalingPolicy, err error) {
	result = &v1.AwsAppautoscalingPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsappautoscalingpolicies").
		Body(awsAppautoscalingPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAppautoscalingPolicy and updates it. Returns the server's representation of the awsAppautoscalingPolicy, and an error, if there is any.
func (c *awsAppautoscalingPolicies) Update(awsAppautoscalingPolicy *v1.AwsAppautoscalingPolicy) (result *v1.AwsAppautoscalingPolicy, err error) {
	result = &v1.AwsAppautoscalingPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsappautoscalingpolicies").
		Name(awsAppautoscalingPolicy.Name).
		Body(awsAppautoscalingPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAppautoscalingPolicy and deletes it. Returns an error if one occurs.
func (c *awsAppautoscalingPolicies) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsappautoscalingpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAppautoscalingPolicies) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsappautoscalingpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAppautoscalingPolicy.
func (c *awsAppautoscalingPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsAppautoscalingPolicy, err error) {
	result = &v1.AwsAppautoscalingPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsappautoscalingpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}