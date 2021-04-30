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
	"context"

	v2alpha1 "github.com/apache/apisix-ingress-controller/pkg/kube/apisix/apis/config/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApisixClusterConfigs implements ApisixClusterConfigInterface
type FakeApisixClusterConfigs struct {
	Fake *FakeApisixV2alpha1
}

var apisixclusterconfigsResource = schema.GroupVersionResource{Group: "apisix.apache.org", Version: "v2alpha1", Resource: "apisixclusterconfigs"}

var apisixclusterconfigsKind = schema.GroupVersionKind{Group: "apisix.apache.org", Version: "v2alpha1", Kind: "ApisixClusterConfig"}

// Get takes name of the apisixClusterConfig, and returns the corresponding apisixClusterConfig object, and an error if there is any.
func (c *FakeApisixClusterConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.ApisixClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(apisixclusterconfigsResource, name), &v2alpha1.ApisixClusterConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ApisixClusterConfig), err
}

// List takes label and field selectors, and returns the list of ApisixClusterConfigs that match those selectors.
func (c *FakeApisixClusterConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.ApisixClusterConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(apisixclusterconfigsResource, apisixclusterconfigsKind, opts), &v2alpha1.ApisixClusterConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.ApisixClusterConfigList{ListMeta: obj.(*v2alpha1.ApisixClusterConfigList).ListMeta}
	for _, item := range obj.(*v2alpha1.ApisixClusterConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apisixClusterConfigs.
func (c *FakeApisixClusterConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(apisixclusterconfigsResource, opts))
}

// Create takes the representation of a apisixClusterConfig and creates it.  Returns the server's representation of the apisixClusterConfig, and an error, if there is any.
func (c *FakeApisixClusterConfigs) Create(ctx context.Context, apisixClusterConfig *v2alpha1.ApisixClusterConfig, opts v1.CreateOptions) (result *v2alpha1.ApisixClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(apisixclusterconfigsResource, apisixClusterConfig), &v2alpha1.ApisixClusterConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ApisixClusterConfig), err
}

// Update takes the representation of a apisixClusterConfig and updates it. Returns the server's representation of the apisixClusterConfig, and an error, if there is any.
func (c *FakeApisixClusterConfigs) Update(ctx context.Context, apisixClusterConfig *v2alpha1.ApisixClusterConfig, opts v1.UpdateOptions) (result *v2alpha1.ApisixClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(apisixclusterconfigsResource, apisixClusterConfig), &v2alpha1.ApisixClusterConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ApisixClusterConfig), err
}

// Delete takes name of the apisixClusterConfig and deletes it. Returns an error if one occurs.
func (c *FakeApisixClusterConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(apisixclusterconfigsResource, name), &v2alpha1.ApisixClusterConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApisixClusterConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(apisixclusterconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.ApisixClusterConfigList{})
	return err
}

// Patch applies the patch and returns the patched apisixClusterConfig.
func (c *FakeApisixClusterConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.ApisixClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(apisixclusterconfigsResource, name, pt, data, subresources...), &v2alpha1.ApisixClusterConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ApisixClusterConfig), err
}
