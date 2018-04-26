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

// FakeAwsRoute53ZoneAssociations implements AwsRoute53ZoneAssociationInterface
type FakeAwsRoute53ZoneAssociations struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsroute53zoneassociationsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsroute53zoneassociations"}

var awsroute53zoneassociationsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsRoute53ZoneAssociation"}

// Get takes name of the awsRoute53ZoneAssociation, and returns the corresponding awsRoute53ZoneAssociation object, and an error if there is any.
func (c *FakeAwsRoute53ZoneAssociations) Get(name string, options v1.GetOptions) (result *aws_v1.AwsRoute53ZoneAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsroute53zoneassociationsResource, c.ns, name), &aws_v1.AwsRoute53ZoneAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsRoute53ZoneAssociation), err
}

// List takes label and field selectors, and returns the list of AwsRoute53ZoneAssociations that match those selectors.
func (c *FakeAwsRoute53ZoneAssociations) List(opts v1.ListOptions) (result *aws_v1.AwsRoute53ZoneAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsroute53zoneassociationsResource, awsroute53zoneassociationsKind, c.ns, opts), &aws_v1.AwsRoute53ZoneAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsRoute53ZoneAssociationList{}
	for _, item := range obj.(*aws_v1.AwsRoute53ZoneAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsRoute53ZoneAssociations.
func (c *FakeAwsRoute53ZoneAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsroute53zoneassociationsResource, c.ns, opts))

}

// Create takes the representation of a awsRoute53ZoneAssociation and creates it.  Returns the server's representation of the awsRoute53ZoneAssociation, and an error, if there is any.
func (c *FakeAwsRoute53ZoneAssociations) Create(awsRoute53ZoneAssociation *aws_v1.AwsRoute53ZoneAssociation) (result *aws_v1.AwsRoute53ZoneAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsroute53zoneassociationsResource, c.ns, awsRoute53ZoneAssociation), &aws_v1.AwsRoute53ZoneAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsRoute53ZoneAssociation), err
}

// Update takes the representation of a awsRoute53ZoneAssociation and updates it. Returns the server's representation of the awsRoute53ZoneAssociation, and an error, if there is any.
func (c *FakeAwsRoute53ZoneAssociations) Update(awsRoute53ZoneAssociation *aws_v1.AwsRoute53ZoneAssociation) (result *aws_v1.AwsRoute53ZoneAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsroute53zoneassociationsResource, c.ns, awsRoute53ZoneAssociation), &aws_v1.AwsRoute53ZoneAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsRoute53ZoneAssociation), err
}

// Delete takes name of the awsRoute53ZoneAssociation and deletes it. Returns an error if one occurs.
func (c *FakeAwsRoute53ZoneAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsroute53zoneassociationsResource, c.ns, name), &aws_v1.AwsRoute53ZoneAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsRoute53ZoneAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsroute53zoneassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsRoute53ZoneAssociationList{})
	return err
}

// Patch applies the patch and returns the patched awsRoute53ZoneAssociation.
func (c *FakeAwsRoute53ZoneAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsRoute53ZoneAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsroute53zoneassociationsResource, c.ns, name, data, subresources...), &aws_v1.AwsRoute53ZoneAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsRoute53ZoneAssociation), err
}