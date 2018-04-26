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

// FakeAwsVpcEndpointSubnetAssociations implements AwsVpcEndpointSubnetAssociationInterface
type FakeAwsVpcEndpointSubnetAssociations struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsvpcendpointsubnetassociationsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsvpcendpointsubnetassociations"}

var awsvpcendpointsubnetassociationsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsVpcEndpointSubnetAssociation"}

// Get takes name of the awsVpcEndpointSubnetAssociation, and returns the corresponding awsVpcEndpointSubnetAssociation object, and an error if there is any.
func (c *FakeAwsVpcEndpointSubnetAssociations) Get(name string, options v1.GetOptions) (result *aws_v1.AwsVpcEndpointSubnetAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsvpcendpointsubnetassociationsResource, c.ns, name), &aws_v1.AwsVpcEndpointSubnetAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpcEndpointSubnetAssociation), err
}

// List takes label and field selectors, and returns the list of AwsVpcEndpointSubnetAssociations that match those selectors.
func (c *FakeAwsVpcEndpointSubnetAssociations) List(opts v1.ListOptions) (result *aws_v1.AwsVpcEndpointSubnetAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsvpcendpointsubnetassociationsResource, awsvpcendpointsubnetassociationsKind, c.ns, opts), &aws_v1.AwsVpcEndpointSubnetAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsVpcEndpointSubnetAssociationList{}
	for _, item := range obj.(*aws_v1.AwsVpcEndpointSubnetAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsVpcEndpointSubnetAssociations.
func (c *FakeAwsVpcEndpointSubnetAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsvpcendpointsubnetassociationsResource, c.ns, opts))

}

// Create takes the representation of a awsVpcEndpointSubnetAssociation and creates it.  Returns the server's representation of the awsVpcEndpointSubnetAssociation, and an error, if there is any.
func (c *FakeAwsVpcEndpointSubnetAssociations) Create(awsVpcEndpointSubnetAssociation *aws_v1.AwsVpcEndpointSubnetAssociation) (result *aws_v1.AwsVpcEndpointSubnetAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsvpcendpointsubnetassociationsResource, c.ns, awsVpcEndpointSubnetAssociation), &aws_v1.AwsVpcEndpointSubnetAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpcEndpointSubnetAssociation), err
}

// Update takes the representation of a awsVpcEndpointSubnetAssociation and updates it. Returns the server's representation of the awsVpcEndpointSubnetAssociation, and an error, if there is any.
func (c *FakeAwsVpcEndpointSubnetAssociations) Update(awsVpcEndpointSubnetAssociation *aws_v1.AwsVpcEndpointSubnetAssociation) (result *aws_v1.AwsVpcEndpointSubnetAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsvpcendpointsubnetassociationsResource, c.ns, awsVpcEndpointSubnetAssociation), &aws_v1.AwsVpcEndpointSubnetAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpcEndpointSubnetAssociation), err
}

// Delete takes name of the awsVpcEndpointSubnetAssociation and deletes it. Returns an error if one occurs.
func (c *FakeAwsVpcEndpointSubnetAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsvpcendpointsubnetassociationsResource, c.ns, name), &aws_v1.AwsVpcEndpointSubnetAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsVpcEndpointSubnetAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsvpcendpointsubnetassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsVpcEndpointSubnetAssociationList{})
	return err
}

// Patch applies the patch and returns the patched awsVpcEndpointSubnetAssociation.
func (c *FakeAwsVpcEndpointSubnetAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsVpcEndpointSubnetAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsvpcendpointsubnetassociationsResource, c.ns, name, data, subresources...), &aws_v1.AwsVpcEndpointSubnetAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpcEndpointSubnetAssociation), err
}