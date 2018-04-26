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

// FakeAwsSsmAssociations implements AwsSsmAssociationInterface
type FakeAwsSsmAssociations struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsssmassociationsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsssmassociations"}

var awsssmassociationsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsSsmAssociation"}

// Get takes name of the awsSsmAssociation, and returns the corresponding awsSsmAssociation object, and an error if there is any.
func (c *FakeAwsSsmAssociations) Get(name string, options v1.GetOptions) (result *aws_v1.AwsSsmAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsssmassociationsResource, c.ns, name), &aws_v1.AwsSsmAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmAssociation), err
}

// List takes label and field selectors, and returns the list of AwsSsmAssociations that match those selectors.
func (c *FakeAwsSsmAssociations) List(opts v1.ListOptions) (result *aws_v1.AwsSsmAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsssmassociationsResource, awsssmassociationsKind, c.ns, opts), &aws_v1.AwsSsmAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsSsmAssociationList{}
	for _, item := range obj.(*aws_v1.AwsSsmAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSsmAssociations.
func (c *FakeAwsSsmAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsssmassociationsResource, c.ns, opts))

}

// Create takes the representation of a awsSsmAssociation and creates it.  Returns the server's representation of the awsSsmAssociation, and an error, if there is any.
func (c *FakeAwsSsmAssociations) Create(awsSsmAssociation *aws_v1.AwsSsmAssociation) (result *aws_v1.AwsSsmAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsssmassociationsResource, c.ns, awsSsmAssociation), &aws_v1.AwsSsmAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmAssociation), err
}

// Update takes the representation of a awsSsmAssociation and updates it. Returns the server's representation of the awsSsmAssociation, and an error, if there is any.
func (c *FakeAwsSsmAssociations) Update(awsSsmAssociation *aws_v1.AwsSsmAssociation) (result *aws_v1.AwsSsmAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsssmassociationsResource, c.ns, awsSsmAssociation), &aws_v1.AwsSsmAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmAssociation), err
}

// Delete takes name of the awsSsmAssociation and deletes it. Returns an error if one occurs.
func (c *FakeAwsSsmAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsssmassociationsResource, c.ns, name), &aws_v1.AwsSsmAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSsmAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsssmassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsSsmAssociationList{})
	return err
}

// Patch applies the patch and returns the patched awsSsmAssociation.
func (c *FakeAwsSsmAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsSsmAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsssmassociationsResource, c.ns, name, data, subresources...), &aws_v1.AwsSsmAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmAssociation), err
}