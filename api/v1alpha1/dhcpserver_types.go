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

// DhcpServerSpec defines the desired state of DhcpServer
type DhcpServerSpec struct {
	Name          string `json:"name"`
	AddArp        bool   `json:"addArp,omitempty"`
	AddressPool   string `json:"addressPool,omitempty"`
	Authoritative string `json:"authoritative,omitempty"`
	Disabled      bool   `json:"disabled.omitempty"`
	Interface     string `json:"interface,omitempty"`
	LeaseScript   string `json:"leaseScript,omitempty"`
}

// DhcpServerStatus defines the observed state of DhcpServer
type DhcpServerStatus struct {
	Identifier string `json:"identifier,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DhcpServer is the Schema for the dhcpservers API
type DhcpServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DhcpServerSpec   `json:"spec,omitempty"`
	Status DhcpServerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DhcpServerList contains a list of DhcpServer
type DhcpServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DhcpServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DhcpServer{}, &DhcpServerList{})
}
