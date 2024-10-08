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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"

	v1beta1 "antrea.io/antrea/pkg/apis/crd/v1beta1"
	scheme "antrea.io/antrea/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// EgressesGetter has a method to return a EgressInterface.
// A group's client should implement this interface.
type EgressesGetter interface {
	Egresses() EgressInterface
}

// EgressInterface has methods to work with Egress resources.
type EgressInterface interface {
	Create(ctx context.Context, egress *v1beta1.Egress, opts v1.CreateOptions) (*v1beta1.Egress, error)
	Update(ctx context.Context, egress *v1beta1.Egress, opts v1.UpdateOptions) (*v1beta1.Egress, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, egress *v1beta1.Egress, opts v1.UpdateOptions) (*v1beta1.Egress, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Egress, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.EgressList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Egress, err error)
	EgressExpansion
}

// egresses implements EgressInterface
type egresses struct {
	*gentype.ClientWithList[*v1beta1.Egress, *v1beta1.EgressList]
}

// newEgresses returns a Egresses
func newEgresses(c *CrdV1beta1Client) *egresses {
	return &egresses{
		gentype.NewClientWithList[*v1beta1.Egress, *v1beta1.EgressList](
			"egresses",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1beta1.Egress { return &v1beta1.Egress{} },
			func() *v1beta1.EgressList { return &v1beta1.EgressList{} }),
	}
}
