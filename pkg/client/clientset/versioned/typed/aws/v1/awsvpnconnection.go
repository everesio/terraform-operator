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

// AwsVpnConnectionsGetter has a method to return a AwsVpnConnectionInterface.
// A group's client should implement this interface.
type AwsVpnConnectionsGetter interface {
	AwsVpnConnections(namespace string) AwsVpnConnectionInterface
}

// AwsVpnConnectionInterface has methods to work with AwsVpnConnection resources.
type AwsVpnConnectionInterface interface {
	Create(*v1.AwsVpnConnection) (*v1.AwsVpnConnection, error)
	Update(*v1.AwsVpnConnection) (*v1.AwsVpnConnection, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsVpnConnection, error)
	List(opts meta_v1.ListOptions) (*v1.AwsVpnConnectionList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsVpnConnection, err error)
	AwsVpnConnectionExpansion
}

// awsVpnConnections implements AwsVpnConnectionInterface
type awsVpnConnections struct {
	client rest.Interface
	ns     string
}

// newAwsVpnConnections returns a AwsVpnConnections
func newAwsVpnConnections(c *TrussleV1Client, namespace string) *awsVpnConnections {
	return &awsVpnConnections{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsVpnConnection, and returns the corresponding awsVpnConnection object, and an error if there is any.
func (c *awsVpnConnections) Get(name string, options meta_v1.GetOptions) (result *v1.AwsVpnConnection, err error) {
	result = &v1.AwsVpnConnection{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsvpnconnections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsVpnConnections that match those selectors.
func (c *awsVpnConnections) List(opts meta_v1.ListOptions) (result *v1.AwsVpnConnectionList, err error) {
	result = &v1.AwsVpnConnectionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsvpnconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsVpnConnections.
func (c *awsVpnConnections) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsvpnconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsVpnConnection and creates it.  Returns the server's representation of the awsVpnConnection, and an error, if there is any.
func (c *awsVpnConnections) Create(awsVpnConnection *v1.AwsVpnConnection) (result *v1.AwsVpnConnection, err error) {
	result = &v1.AwsVpnConnection{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsvpnconnections").
		Body(awsVpnConnection).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsVpnConnection and updates it. Returns the server's representation of the awsVpnConnection, and an error, if there is any.
func (c *awsVpnConnections) Update(awsVpnConnection *v1.AwsVpnConnection) (result *v1.AwsVpnConnection, err error) {
	result = &v1.AwsVpnConnection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsvpnconnections").
		Name(awsVpnConnection.Name).
		Body(awsVpnConnection).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsVpnConnection and deletes it. Returns an error if one occurs.
func (c *awsVpnConnections) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsvpnconnections").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsVpnConnections) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsvpnconnections").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsVpnConnection.
func (c *awsVpnConnections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsVpnConnection, err error) {
	result = &v1.AwsVpnConnection{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsvpnconnections").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}