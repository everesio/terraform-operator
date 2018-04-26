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

// FakeAwsIamUserPolicyAttachments implements AwsIamUserPolicyAttachmentInterface
type FakeAwsIamUserPolicyAttachments struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsiamuserpolicyattachmentsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsiamuserpolicyattachments"}

var awsiamuserpolicyattachmentsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsIamUserPolicyAttachment"}

// Get takes name of the awsIamUserPolicyAttachment, and returns the corresponding awsIamUserPolicyAttachment object, and an error if there is any.
func (c *FakeAwsIamUserPolicyAttachments) Get(name string, options v1.GetOptions) (result *aws_v1.AwsIamUserPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsiamuserpolicyattachmentsResource, c.ns, name), &aws_v1.AwsIamUserPolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIamUserPolicyAttachment), err
}

// List takes label and field selectors, and returns the list of AwsIamUserPolicyAttachments that match those selectors.
func (c *FakeAwsIamUserPolicyAttachments) List(opts v1.ListOptions) (result *aws_v1.AwsIamUserPolicyAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsiamuserpolicyattachmentsResource, awsiamuserpolicyattachmentsKind, c.ns, opts), &aws_v1.AwsIamUserPolicyAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsIamUserPolicyAttachmentList{}
	for _, item := range obj.(*aws_v1.AwsIamUserPolicyAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIamUserPolicyAttachments.
func (c *FakeAwsIamUserPolicyAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsiamuserpolicyattachmentsResource, c.ns, opts))

}

// Create takes the representation of a awsIamUserPolicyAttachment and creates it.  Returns the server's representation of the awsIamUserPolicyAttachment, and an error, if there is any.
func (c *FakeAwsIamUserPolicyAttachments) Create(awsIamUserPolicyAttachment *aws_v1.AwsIamUserPolicyAttachment) (result *aws_v1.AwsIamUserPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsiamuserpolicyattachmentsResource, c.ns, awsIamUserPolicyAttachment), &aws_v1.AwsIamUserPolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIamUserPolicyAttachment), err
}

// Update takes the representation of a awsIamUserPolicyAttachment and updates it. Returns the server's representation of the awsIamUserPolicyAttachment, and an error, if there is any.
func (c *FakeAwsIamUserPolicyAttachments) Update(awsIamUserPolicyAttachment *aws_v1.AwsIamUserPolicyAttachment) (result *aws_v1.AwsIamUserPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsiamuserpolicyattachmentsResource, c.ns, awsIamUserPolicyAttachment), &aws_v1.AwsIamUserPolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIamUserPolicyAttachment), err
}

// Delete takes name of the awsIamUserPolicyAttachment and deletes it. Returns an error if one occurs.
func (c *FakeAwsIamUserPolicyAttachments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsiamuserpolicyattachmentsResource, c.ns, name), &aws_v1.AwsIamUserPolicyAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIamUserPolicyAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsiamuserpolicyattachmentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsIamUserPolicyAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched awsIamUserPolicyAttachment.
func (c *FakeAwsIamUserPolicyAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsIamUserPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsiamuserpolicyattachmentsResource, c.ns, name, data, subresources...), &aws_v1.AwsIamUserPolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIamUserPolicyAttachment), err
}