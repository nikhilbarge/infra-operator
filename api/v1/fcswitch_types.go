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

// FcSwitchSpec defines the desired state of FcSwitch
type FcSwitchSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of FcSwitch. Edit FcSwitch_types.go to remove/update
	Name string `json:"name,omitempty"`
	IPAddress string `json:"ipaddress,omitempty"`
	TimeStamp string `json:"timestamp,omitempty"`
}

// FcSwitchStatus defines the observed state of FcSwitch
type FcSwitchStatus struct {
	Status string `json:"status,omitempty"` 
	GatewayIP string `json:"gatewayip,omitempty"`
	LastUpdated string `json:"lastupdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// FcSwitch is the Schema for the fcswitches API
type FcSwitch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FcSwitchSpec   `json:"spec,omitempty"`
	Status FcSwitchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FcSwitchList contains a list of FcSwitch
type FcSwitchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FcSwitch `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FcSwitch{}, &FcSwitchList{})
}
