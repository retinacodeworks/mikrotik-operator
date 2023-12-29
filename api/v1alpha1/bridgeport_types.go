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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BridgePortSpec defines the desired state of BridgePort
type BridgePortSpec struct {
	BridgeRef  v1.ObjectReference `json:"bridgeRef"`
	Interface  string             `json:"interface"`
	PortVlanId int                `json:"portVlanId,omitempty"`
	Comment    string             `json:"comment"`
}

// BridgePortStatus defines the observed state of BridgePort
type BridgePortStatus struct {
	Identifier string `json:"identifier,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// BridgePort is the Schema for the bridgeports API
type BridgePort struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BridgePortSpec   `json:"spec,omitempty"`
	Status BridgePortStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BridgePortList contains a list of BridgePort
type BridgePortList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BridgePort `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BridgePort{}, &BridgePortList{})
}
