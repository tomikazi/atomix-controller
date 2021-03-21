// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v2beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ValueSpec specifies a Value
type ValueSpec struct {
	Storage StorageSpec    `json:"storage,omitempty"`
	Cache   ValueCacheSpec `json:"cache,omitempty"`
}

// ValueCacheSpec specifies a Value cache configuration
type ValueCacheSpec struct {
	Enabled bool `json:"enabled,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Value is the Schema for the Value API
// +k8s:openapi-gen=true
type Value struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ValueSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ValueList contains a list of Value
type ValueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is the Value of items in the list
	Items []Value `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Value{}, &ValueList{})
}
