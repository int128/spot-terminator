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

// SpotInterruptionSpec represents an EC2 Spot Instance Interruption Warning event.
// See https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-instance-termination-notices.html for details.
type SpotInterruptionSpec struct {
	// Queue is the reference to the origin of the event.
	Queue QueueReference `json:"queue,omitempty"`

	// EventTimestamp is the time at which the event was generated by EventBridge.
	EventTimestamp metav1.Time `json:"eventTimestamp,omitempty"`

	// InstanceID represents the instance affected by the event.
	InstanceID string `json:"instanceID,omitempty"`

	// AvailabilityZone represents the availability zone in which the instance is running.
	AvailabilityZone string `json:"availabilityZone,omitempty"`
}

// SpotInterruptionStatus defines the observed state of SpotInterruption
type SpotInterruptionStatus struct {
	// Timestamp at which the SpotInterruption was reconciled successfully.
	// +optional
	ReconciledAt metav1.Time `json:"reconciledAt,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// SpotInterruption is the Schema for the spotinterruptions API
type SpotInterruption struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpotInterruptionSpec   `json:"spec,omitempty"`
	Status SpotInterruptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpotInterruptionList contains a list of SpotInterruption
type SpotInterruptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotInterruption `json:"items"`
}

// SpotInterruptionReference represents a reference to a SpotInterruption object.
type SpotInterruptionReference struct {
	// Name is the name of the SpotInterruption object.
	Name string `json:"name,omitempty"`
}

// SpotInterruptionReferenceTo returns a reference to the given SpotInterruption object.
func SpotInterruptionReferenceTo(spotInterruption SpotInterruption) SpotInterruptionReference {
	return SpotInterruptionReference{
		Name: spotInterruption.Name,
	}
}

func init() {
	SchemeBuilder.Register(&SpotInterruption{}, &SpotInterruptionList{})
}
