/*
Copyright 2020 reoring.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SleepSpec defines the desired state of Sleep
type SleepSpec struct {
	// Pride name is a name for pride
	PrideName string `json:"prideName"`

	// Bedtime hour. e.g. 21
	Bedtime int `json:"bedtime"`

	// Wakeup hour. e.g. 8
	Wakeup int `json:"wakeup"`
}

// SleepStatus defines the observed state of Sleep
type SleepStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Sleeping bool `json:"sleeping"`
}

// +kubebuilder:object:root=true

// Sleep is the Schema for the sleeps API
type Sleep struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SleepSpec   `json:"spec,omitempty"`
	Status SleepStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SleepList contains a list of Sleep
type SleepList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Sleep `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Sleep{}, &SleepList{})
}
