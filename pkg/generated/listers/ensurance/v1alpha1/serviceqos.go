// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gocrane/api/ensurance/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceQOSLister helps list ServiceQOSs.
// All objects returned here must be treated as read-only.
type ServiceQOSLister interface {
	// List lists all ServiceQOSs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceQOS, err error)
	// Get retrieves the ServiceQOS from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServiceQOS, error)
	ServiceQOSListerExpansion
}

// serviceQOSLister implements the ServiceQOSLister interface.
type serviceQOSLister struct {
	indexer cache.Indexer
}

// NewServiceQOSLister returns a new ServiceQOSLister.
func NewServiceQOSLister(indexer cache.Indexer) ServiceQOSLister {
	return &serviceQOSLister{indexer: indexer}
}

// List lists all ServiceQOSs in the indexer.
func (s *serviceQOSLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceQOS, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceQOS))
	})
	return ret, err
}

// Get retrieves the ServiceQOS from the index for a given name.
func (s *serviceQOSLister) Get(name string) (*v1alpha1.ServiceQOS, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("serviceqos"), name)
	}
	return obj.(*v1alpha1.ServiceQOS), nil
}
