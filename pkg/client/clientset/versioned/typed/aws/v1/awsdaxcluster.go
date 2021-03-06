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

// AwsDaxClustersGetter has a method to return a AwsDaxClusterInterface.
// A group's client should implement this interface.
type AwsDaxClustersGetter interface {
	AwsDaxClusters(namespace string) AwsDaxClusterInterface
}

// AwsDaxClusterInterface has methods to work with AwsDaxCluster resources.
type AwsDaxClusterInterface interface {
	Create(*v1.AwsDaxCluster) (*v1.AwsDaxCluster, error)
	Update(*v1.AwsDaxCluster) (*v1.AwsDaxCluster, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsDaxCluster, error)
	List(opts meta_v1.ListOptions) (*v1.AwsDaxClusterList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsDaxCluster, err error)
	AwsDaxClusterExpansion
}

// awsDaxClusters implements AwsDaxClusterInterface
type awsDaxClusters struct {
	client rest.Interface
	ns     string
}

// newAwsDaxClusters returns a AwsDaxClusters
func newAwsDaxClusters(c *TrussleV1Client, namespace string) *awsDaxClusters {
	return &awsDaxClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsDaxCluster, and returns the corresponding awsDaxCluster object, and an error if there is any.
func (c *awsDaxClusters) Get(name string, options meta_v1.GetOptions) (result *v1.AwsDaxCluster, err error) {
	result = &v1.AwsDaxCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdaxclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDaxClusters that match those selectors.
func (c *awsDaxClusters) List(opts meta_v1.ListOptions) (result *v1.AwsDaxClusterList, err error) {
	result = &v1.AwsDaxClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdaxclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDaxClusters.
func (c *awsDaxClusters) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsdaxclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsDaxCluster and creates it.  Returns the server's representation of the awsDaxCluster, and an error, if there is any.
func (c *awsDaxClusters) Create(awsDaxCluster *v1.AwsDaxCluster) (result *v1.AwsDaxCluster, err error) {
	result = &v1.AwsDaxCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsdaxclusters").
		Body(awsDaxCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDaxCluster and updates it. Returns the server's representation of the awsDaxCluster, and an error, if there is any.
func (c *awsDaxClusters) Update(awsDaxCluster *v1.AwsDaxCluster) (result *v1.AwsDaxCluster, err error) {
	result = &v1.AwsDaxCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsdaxclusters").
		Name(awsDaxCluster.Name).
		Body(awsDaxCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDaxCluster and deletes it. Returns an error if one occurs.
func (c *awsDaxClusters) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdaxclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDaxClusters) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdaxclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDaxCluster.
func (c *awsDaxClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsDaxCluster, err error) {
	result = &v1.AwsDaxCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsdaxclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
