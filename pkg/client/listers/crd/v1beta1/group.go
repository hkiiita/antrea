// Copyright 2024 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "antrea.io/antrea/pkg/apis/crd/v1beta1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// GroupLister helps list Groups.
// All objects returned here must be treated as read-only.
type GroupLister interface {
	// List lists all Groups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Group, err error)
	// Groups returns an object that can list and get Groups.
	Groups(namespace string) GroupNamespaceLister
	GroupListerExpansion
}

// groupLister implements the GroupLister interface.
type groupLister struct {
	listers.ResourceIndexer[*v1beta1.Group]
}

// NewGroupLister returns a new GroupLister.
func NewGroupLister(indexer cache.Indexer) GroupLister {
	return &groupLister{listers.New[*v1beta1.Group](indexer, v1beta1.Resource("group"))}
}

// Groups returns an object that can list and get Groups.
func (s *groupLister) Groups(namespace string) GroupNamespaceLister {
	return groupNamespaceLister{listers.NewNamespaced[*v1beta1.Group](s.ResourceIndexer, namespace)}
}

// GroupNamespaceLister helps list and get Groups.
// All objects returned here must be treated as read-only.
type GroupNamespaceLister interface {
	// List lists all Groups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Group, err error)
	// Get retrieves the Group from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.Group, error)
	GroupNamespaceListerExpansion
}

// groupNamespaceLister implements the GroupNamespaceLister
// interface.
type groupNamespaceLister struct {
	listers.ResourceIndexer[*v1beta1.Group]
}
