/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodPolicySpec defines the desired state of PodPolicy
type PodPolicySpec struct {
	// TerminateOnSpotInterruption is whether to terminate a Pod on SpotInterruption.
	// +optional
	TerminateOnSpotInterruption bool `json:"terminateOnSpotInterruption,omitempty"`
}

// PodPolicyStatus defines the observed state of PodPolicy
type PodPolicyStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// PodPolicy is the Schema for the podpolicies API
type PodPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodPolicySpec   `json:"spec,omitempty"`
	Status PodPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PodPolicyList contains a list of PodPolicy
type PodPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodPolicy{}, &PodPolicyList{})
}
