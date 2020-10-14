/*


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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// InfrastructurePoolSpec defines the desired state of InfrastructurePool
type InfrastructurePoolSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Name string `json:"name,omitempty"`
	IPAddress string `json:"ipaddress,omitempty"`
	TimeStamp string `json:"timestamp,omitempty"`
}

// InfrastructurePoolStatus defines the observed state of InfrastructurePool
type InfrastructurePoolStatus struct {
	Status string `json:"status,omitempty"` 
	IPAddress string `json:"ipaddress,omitempty"`
	LastUpdated string `json:"lastupdated,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// InfrastructurePool is the Schema for the infrastructurepools API
type InfrastructurePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InfrastructurePoolSpec   `json:"spec,omitempty"`
	Status InfrastructurePoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InfrastructurePoolList contains a list of InfrastructurePool
type InfrastructurePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InfrastructurePool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InfrastructurePool{}, &InfrastructurePoolList{})
}
