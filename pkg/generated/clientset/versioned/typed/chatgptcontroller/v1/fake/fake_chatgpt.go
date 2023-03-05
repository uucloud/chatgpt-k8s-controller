/*
Copyright 2023 uucloud.

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
	chatgptcontrollerv1 "chatgpt-k8s-controller/pkg/apis/chatgptcontroller/v1"
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeChatGPTs implements ChatGPTInterface
type FakeChatGPTs struct {
	Fake *FakeChatgptcontrollerV1
	ns   string
}

var chatgptsResource = schema.GroupVersionResource{Group: "chatgptcontroller.uucloud.top", Version: "v1", Resource: "chatgpts"}

var chatgptsKind = schema.GroupVersionKind{Group: "chatgptcontroller.uucloud.top", Version: "v1", Kind: "ChatGPT"}

// Get takes name of the chatGPT, and returns the corresponding chatGPT object, and an error if there is any.
func (c *FakeChatGPTs) Get(ctx context.Context, name string, options v1.GetOptions) (result *chatgptcontrollerv1.ChatGPT, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(chatgptsResource, c.ns, name), &chatgptcontrollerv1.ChatGPT{})

	if obj == nil {
		return nil, err
	}
	return obj.(*chatgptcontrollerv1.ChatGPT), err
}

// List takes label and field selectors, and returns the list of ChatGPTs that match those selectors.
func (c *FakeChatGPTs) List(ctx context.Context, opts v1.ListOptions) (result *chatgptcontrollerv1.ChatGPTList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(chatgptsResource, chatgptsKind, c.ns, opts), &chatgptcontrollerv1.ChatGPTList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &chatgptcontrollerv1.ChatGPTList{ListMeta: obj.(*chatgptcontrollerv1.ChatGPTList).ListMeta}
	for _, item := range obj.(*chatgptcontrollerv1.ChatGPTList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested chatGPTs.
func (c *FakeChatGPTs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(chatgptsResource, c.ns, opts))

}

// Create takes the representation of a chatGPT and creates it.  Returns the server's representation of the chatGPT, and an error, if there is any.
func (c *FakeChatGPTs) Create(ctx context.Context, chatGPT *chatgptcontrollerv1.ChatGPT, opts v1.CreateOptions) (result *chatgptcontrollerv1.ChatGPT, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(chatgptsResource, c.ns, chatGPT), &chatgptcontrollerv1.ChatGPT{})

	if obj == nil {
		return nil, err
	}
	return obj.(*chatgptcontrollerv1.ChatGPT), err
}

// Update takes the representation of a chatGPT and updates it. Returns the server's representation of the chatGPT, and an error, if there is any.
func (c *FakeChatGPTs) Update(ctx context.Context, chatGPT *chatgptcontrollerv1.ChatGPT, opts v1.UpdateOptions) (result *chatgptcontrollerv1.ChatGPT, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(chatgptsResource, c.ns, chatGPT), &chatgptcontrollerv1.ChatGPT{})

	if obj == nil {
		return nil, err
	}
	return obj.(*chatgptcontrollerv1.ChatGPT), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeChatGPTs) UpdateStatus(ctx context.Context, chatGPT *chatgptcontrollerv1.ChatGPT, opts v1.UpdateOptions) (*chatgptcontrollerv1.ChatGPT, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(chatgptsResource, "status", c.ns, chatGPT), &chatgptcontrollerv1.ChatGPT{})

	if obj == nil {
		return nil, err
	}
	return obj.(*chatgptcontrollerv1.ChatGPT), err
}

// Delete takes name of the chatGPT and deletes it. Returns an error if one occurs.
func (c *FakeChatGPTs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(chatgptsResource, c.ns, name, opts), &chatgptcontrollerv1.ChatGPT{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeChatGPTs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(chatgptsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &chatgptcontrollerv1.ChatGPTList{})
	return err
}

// Patch applies the patch and returns the patched chatGPT.
func (c *FakeChatGPTs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *chatgptcontrollerv1.ChatGPT, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(chatgptsResource, c.ns, name, pt, data, subresources...), &chatgptcontrollerv1.ChatGPT{})

	if obj == nil {
		return nil, err
	}
	return obj.(*chatgptcontrollerv1.ChatGPT), err
}
