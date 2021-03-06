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

// FakeAwsCodepipelines implements AwsCodepipelineInterface
type FakeAwsCodepipelines struct {
	Fake *FakeTrussleV1
	ns   string
}

var awscodepipelinesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awscodepipelines"}

var awscodepipelinesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsCodepipeline"}

// Get takes name of the awsCodepipeline, and returns the corresponding awsCodepipeline object, and an error if there is any.
func (c *FakeAwsCodepipelines) Get(name string, options v1.GetOptions) (result *aws_v1.AwsCodepipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscodepipelinesResource, c.ns, name), &aws_v1.AwsCodepipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCodepipeline), err
}

// List takes label and field selectors, and returns the list of AwsCodepipelines that match those selectors.
func (c *FakeAwsCodepipelines) List(opts v1.ListOptions) (result *aws_v1.AwsCodepipelineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscodepipelinesResource, awscodepipelinesKind, c.ns, opts), &aws_v1.AwsCodepipelineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsCodepipelineList{}
	for _, item := range obj.(*aws_v1.AwsCodepipelineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCodepipelines.
func (c *FakeAwsCodepipelines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscodepipelinesResource, c.ns, opts))

}

// Create takes the representation of a awsCodepipeline and creates it.  Returns the server's representation of the awsCodepipeline, and an error, if there is any.
func (c *FakeAwsCodepipelines) Create(awsCodepipeline *aws_v1.AwsCodepipeline) (result *aws_v1.AwsCodepipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscodepipelinesResource, c.ns, awsCodepipeline), &aws_v1.AwsCodepipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCodepipeline), err
}

// Update takes the representation of a awsCodepipeline and updates it. Returns the server's representation of the awsCodepipeline, and an error, if there is any.
func (c *FakeAwsCodepipelines) Update(awsCodepipeline *aws_v1.AwsCodepipeline) (result *aws_v1.AwsCodepipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscodepipelinesResource, c.ns, awsCodepipeline), &aws_v1.AwsCodepipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCodepipeline), err
}

// Delete takes name of the awsCodepipeline and deletes it. Returns an error if one occurs.
func (c *FakeAwsCodepipelines) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscodepipelinesResource, c.ns, name), &aws_v1.AwsCodepipeline{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCodepipelines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscodepipelinesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsCodepipelineList{})
	return err
}

// Patch applies the patch and returns the patched awsCodepipeline.
func (c *FakeAwsCodepipelines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsCodepipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscodepipelinesResource, c.ns, name, data, subresources...), &aws_v1.AwsCodepipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCodepipeline), err
}
