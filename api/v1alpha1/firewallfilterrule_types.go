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

// FirewallFilterRuleSpec defines the desired state of FirewallFilterRule
type FirewallFilterRuleSpec struct {
	Chain             string   `json:"chain"`
	Action            string   `json:"action,omitempty"`
	Comment           string   `json:"comment,omitempty"`
	ConnectionState   []string `json:"connectionState,omitempty"`
	DestAddress       string   `json:"destAddress,omitempty"`
	DestAddressList   string   `json:"destAddressList,omitempty"`
	DestAddressType   string   `json:"destAddressType,omitempty"`
	DestPort          string   `json:"destPort,omitempty"`
	InBridgePort      string   `json:"inBridgePort,omitempty"`
	InBridgePortList  string   `json:"inBridgePortList,omitempty"`
	InInterface       string   `json:"inInterface,omitempty"`
	InInterfaceList   string   `json:"inInterfaceList,omitempty"`
	LogPrefix         string   `json:"logPrefix,omitempty"`
	OutBridgePort     string   `json:"outBridgePort,omitempty"`
	OutBridgePortList string   `json:"outBridgePortList,omitempty"`
	OutInterface      string   `json:"outInterface,omitempty"`
	OutInterfaceList  string   `json:"outInterfaceList,omitempty"`
	Port              int      `json:"port,omitempty"`
	Priority          int      `json:"priority,omitempty"`
	Protocol          string   `json:"protocol,omitempty"`
	SrcAddress        string   `json:"srcAddress,omitempty"`
	SrcAddressList    string   `json:"srcAddressList,omitempty"`
	SrcAddressType    string   `json:"srcAddressType,omitempty"`
	SrcPort           int      `json:"srcPort,omitempty"`
}

// FirewallFilterRuleStatus defines the observed state of FirewallFilterRule
type FirewallFilterRuleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// FirewallFilterRule is the Schema for the firewallfilterrules API
type FirewallFilterRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirewallFilterRuleSpec   `json:"spec,omitempty"`
	Status FirewallFilterRuleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FirewallFilterRuleList contains a list of FirewallFilterRule
type FirewallFilterRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallFilterRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FirewallFilterRule{}, &FirewallFilterRuleList{})
}
