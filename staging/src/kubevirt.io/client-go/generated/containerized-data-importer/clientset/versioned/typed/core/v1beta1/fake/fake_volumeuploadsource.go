/*
Copyright The KubeVirt Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// FakeVolumeUploadSources implements VolumeUploadSourceInterface
type FakeVolumeUploadSources struct {
	Fake *FakeCdiV1beta1
	ns   string
}

var volumeuploadsourcesResource = schema.GroupVersionResource{Group: "cdi.kubevirt.io", Version: "v1beta1", Resource: "volumeuploadsources"}

var volumeuploadsourcesKind = schema.GroupVersionKind{Group: "cdi.kubevirt.io", Version: "v1beta1", Kind: "VolumeUploadSource"}

// Get takes name of the volumeUploadSource, and returns the corresponding volumeUploadSource object, and an error if there is any.
func (c *FakeVolumeUploadSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VolumeUploadSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(volumeuploadsourcesResource, c.ns, name), &v1beta1.VolumeUploadSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeUploadSource), err
}

// List takes label and field selectors, and returns the list of VolumeUploadSources that match those selectors.
func (c *FakeVolumeUploadSources) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VolumeUploadSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(volumeuploadsourcesResource, volumeuploadsourcesKind, c.ns, opts), &v1beta1.VolumeUploadSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VolumeUploadSourceList{ListMeta: obj.(*v1beta1.VolumeUploadSourceList).ListMeta}
	for _, item := range obj.(*v1beta1.VolumeUploadSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeUploadSources.
func (c *FakeVolumeUploadSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(volumeuploadsourcesResource, c.ns, opts))

}

// Create takes the representation of a volumeUploadSource and creates it.  Returns the server's representation of the volumeUploadSource, and an error, if there is any.
func (c *FakeVolumeUploadSources) Create(ctx context.Context, volumeUploadSource *v1beta1.VolumeUploadSource, opts v1.CreateOptions) (result *v1beta1.VolumeUploadSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(volumeuploadsourcesResource, c.ns, volumeUploadSource), &v1beta1.VolumeUploadSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeUploadSource), err
}

// Update takes the representation of a volumeUploadSource and updates it. Returns the server's representation of the volumeUploadSource, and an error, if there is any.
func (c *FakeVolumeUploadSources) Update(ctx context.Context, volumeUploadSource *v1beta1.VolumeUploadSource, opts v1.UpdateOptions) (result *v1beta1.VolumeUploadSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(volumeuploadsourcesResource, c.ns, volumeUploadSource), &v1beta1.VolumeUploadSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeUploadSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVolumeUploadSources) UpdateStatus(ctx context.Context, volumeUploadSource *v1beta1.VolumeUploadSource, opts v1.UpdateOptions) (*v1beta1.VolumeUploadSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(volumeuploadsourcesResource, "status", c.ns, volumeUploadSource), &v1beta1.VolumeUploadSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeUploadSource), err
}

// Delete takes name of the volumeUploadSource and deletes it. Returns an error if one occurs.
func (c *FakeVolumeUploadSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(volumeuploadsourcesResource, c.ns, name), &v1beta1.VolumeUploadSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeUploadSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(volumeuploadsourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VolumeUploadSourceList{})
	return err
}

// Patch applies the patch and returns the patched volumeUploadSource.
func (c *FakeVolumeUploadSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeUploadSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(volumeuploadsourcesResource, c.ns, name, pt, data, subresources...), &v1beta1.VolumeUploadSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeUploadSource), err
}
