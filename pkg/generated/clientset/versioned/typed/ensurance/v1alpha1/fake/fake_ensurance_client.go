// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/gocrane/api/pkg/generated/clientset/versioned/typed/ensurance/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeEnsuranceV1alpha1 struct {
	*testing.Fake
}

func (c *FakeEnsuranceV1alpha1) AvoidanceActions() v1alpha1.AvoidanceActionInterface {
	return &FakeAvoidanceActions{c}
}

func (c *FakeEnsuranceV1alpha1) NodeQOSs() v1alpha1.NodeQOSInterface {
	return &FakeNodeQOSs{c}
}

func (c *FakeEnsuranceV1alpha1) PodQOSEnsurancePolicies(namespace string) v1alpha1.PodQOSEnsurancePolicyInterface {
	return &FakePodQOSEnsurancePolicies{c, namespace}
}

func (c *FakeEnsuranceV1alpha1) ServicePolicies() v1alpha1.ServicePolicyInterface {
	return &FakeServicePolicies{c}
}

func (c *FakeEnsuranceV1alpha1) ServiceQOSs() v1alpha1.ServiceQOSInterface {
	return &FakeServiceQOSs{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeEnsuranceV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
