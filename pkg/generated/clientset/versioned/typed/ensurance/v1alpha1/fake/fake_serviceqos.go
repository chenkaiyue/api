// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gocrane/api/ensurance/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceQOSs implements ServiceQOSInterface
type FakeServiceQOSs struct {
	Fake *FakeEnsuranceV1alpha1
}

var serviceqossResource = schema.GroupVersionResource{Group: "ensurance.crane.io", Version: "v1alpha1", Resource: "serviceqoss"}

var serviceqossKind = schema.GroupVersionKind{Group: "ensurance.crane.io", Version: "v1alpha1", Kind: "PodQOS"}

// Get takes name of the serviceQOS, and returns the corresponding serviceQOS object, and an error if there is any.
func (c *FakeServiceQOSs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PodQOS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(serviceqossResource, name), &v1alpha1.PodQOS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodQOS), err
}

// List takes label and field selectors, and returns the list of ServiceQOSs that match those selectors.
func (c *FakeServiceQOSs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceQOSList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(serviceqossResource, serviceqossKind, opts), &v1alpha1.ServiceQOSList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceQOSList{ListMeta: obj.(*v1alpha1.ServiceQOSList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceQOSList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceQOSs.
func (c *FakeServiceQOSs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(serviceqossResource, opts))
}

// Create takes the representation of a serviceQOS and creates it.  Returns the server's representation of the serviceQOS, and an error, if there is any.
func (c *FakeServiceQOSs) Create(ctx context.Context, serviceQOS *v1alpha1.PodQOS, opts v1.CreateOptions) (result *v1alpha1.PodQOS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(serviceqossResource, serviceQOS), &v1alpha1.PodQOS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodQOS), err
}

// Update takes the representation of a serviceQOS and updates it. Returns the server's representation of the serviceQOS, and an error, if there is any.
func (c *FakeServiceQOSs) Update(ctx context.Context, serviceQOS *v1alpha1.PodQOS, opts v1.UpdateOptions) (result *v1alpha1.PodQOS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(serviceqossResource, serviceQOS), &v1alpha1.PodQOS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodQOS), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceQOSs) UpdateStatus(ctx context.Context, serviceQOS *v1alpha1.PodQOS, opts v1.UpdateOptions) (*v1alpha1.PodQOS, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(serviceqossResource, "status", serviceQOS), &v1alpha1.PodQOS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodQOS), err
}

// Delete takes name of the serviceQOS and deletes it. Returns an error if one occurs.
func (c *FakeServiceQOSs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(serviceqossResource, name), &v1alpha1.PodQOS{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceQOSs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(serviceqossResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceQOSList{})
	return err
}

// Patch applies the patch and returns the patched serviceQOS.
func (c *FakeServiceQOSs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PodQOS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(serviceqossResource, name, pt, data, subresources...), &v1alpha1.PodQOS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodQOS), err
}
