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

// RouterSpec defines the desired state of Router
type RouterSpec struct {
	Address string         `json:"address"`
	Tls     RouterTlsSpec  `json:"tls,omitempty"`
	Auth    RouterAuthSpec `json:"auth"`
}

type RouterTlsSpec struct {
	CertificateAuthority string `json:"certificateAuthority"`
}

type RouterAuthSpec struct {
	Type     string                 `json:"type,omitempty"`
	Password RouterAuthPasswordSpec `json:"password,omitempty"`
}

type RouterAuthPasswordSpec struct {
	Username          string             `json:"username"`
	Password          string             `json:"password,omitempty"`
	PasswordSecretRef v1.SecretReference `json:"passwordSecretRef,omitempty"`
}

// RouterStatus defines the observed state of Router
type RouterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Version string `json:"version"`
	Model   string `json:"model"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Router is the Schema for the routers API
type Router struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouterSpec   `json:"spec,omitempty"`
	Status RouterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RouterList contains a list of Router
type RouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Router `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Router{}, &RouterList{})
}
