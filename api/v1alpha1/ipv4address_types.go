/*
Copyright 2023.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Ipv4AddressSpec defines the desired state of Ipv4Address
type Ipv4AddressSpec struct {
	// The IPV4 Address to assign
	Address   string `json:"address"`
	Network   string `json:"network,omitempty"`
	Interface string `json:"interface,omitempty"`

	//RouterRef v1.ObjectReference `json:"routerRef"`
}

// Ipv4AddressStatus defines the observed state of Ipv4Address
type Ipv4AddressStatus struct {
	Identifier string             `json:"identifier,omitempty"`
	Conditions []metav1.Condition `json:"conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Ipv4Address is the Schema for the ipv4addresses API
type Ipv4Address struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ipv4AddressSpec   `json:"spec,omitempty"`
	Status Ipv4AddressStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Ipv4AddressList contains a list of Ipv4Address
type Ipv4AddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ipv4Address `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Ipv4Address{}, &Ipv4AddressList{})
}
