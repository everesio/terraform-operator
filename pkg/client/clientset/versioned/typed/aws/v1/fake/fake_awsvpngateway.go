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

package fake

import (
	aws_v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAwsVpnGateways implements AwsVpnGatewayInterface
type FakeAwsVpnGateways struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsvpngatewaysResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsvpngateways"}

var awsvpngatewaysKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsVpnGateway"}

// Get takes name of the awsVpnGateway, and returns the corresponding awsVpnGateway object, and an error if there is any.
func (c *FakeAwsVpnGateways) Get(name string, options v1.GetOptions) (result *aws_v1.AwsVpnGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsvpngatewaysResource, c.ns, name), &aws_v1.AwsVpnGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpnGateway), err
}

// List takes label and field selectors, and returns the list of AwsVpnGateways that match those selectors.
func (c *FakeAwsVpnGateways) List(opts v1.ListOptions) (result *aws_v1.AwsVpnGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsvpngatewaysResource, awsvpngatewaysKind, c.ns, opts), &aws_v1.AwsVpnGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsVpnGatewayList{}
	for _, item := range obj.(*aws_v1.AwsVpnGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsVpnGateways.
func (c *FakeAwsVpnGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsvpngatewaysResource, c.ns, opts))

}

// Create takes the representation of a awsVpnGateway and creates it.  Returns the server's representation of the awsVpnGateway, and an error, if there is any.
func (c *FakeAwsVpnGateways) Create(awsVpnGateway *aws_v1.AwsVpnGateway) (result *aws_v1.AwsVpnGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsvpngatewaysResource, c.ns, awsVpnGateway), &aws_v1.AwsVpnGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpnGateway), err
}

// Update takes the representation of a awsVpnGateway and updates it. Returns the server's representation of the awsVpnGateway, and an error, if there is any.
func (c *FakeAwsVpnGateways) Update(awsVpnGateway *aws_v1.AwsVpnGateway) (result *aws_v1.AwsVpnGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsvpngatewaysResource, c.ns, awsVpnGateway), &aws_v1.AwsVpnGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpnGateway), err
}

// Delete takes name of the awsVpnGateway and deletes it. Returns an error if one occurs.
func (c *FakeAwsVpnGateways) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsvpngatewaysResource, c.ns, name), &aws_v1.AwsVpnGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsVpnGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsvpngatewaysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsVpnGatewayList{})
	return err
}

// Patch applies the patch and returns the patched awsVpnGateway.
func (c *FakeAwsVpnGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsVpnGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsvpngatewaysResource, c.ns, name, data, subresources...), &aws_v1.AwsVpnGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpnGateway), err
}
