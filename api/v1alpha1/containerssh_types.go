/*
Copyright 2021.

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
	"github.com/containerssh/configuration"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// embeddet type to add DeepCopyInto
type AppConfig configuration.AppConfig
func (in *AppConfig) DeepCopyInto(out *AppConfig) {
	y, _ := yaml.Marshal(*in)
	err = yaml.Unmarshal(y, out)
	return
}

// ContainerSSHSpec defines the desired state of ContainerSSH
type ContainerSSHSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ContainerSSH. Edit ContainerSSH_types.go to remove/update
	Config AppConfig `json:"config,omitempty"`
}

// ContainerSSHStatus defines the observed state of ContainerSSH
type ContainerSSHStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ContainerSSH is the Schema for the containersshes API
type ContainerSSH struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerSSHSpec   `json:"spec,omitempty"`
	Status ContainerSSHStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerSSHList contains a list of ContainerSSH
type ContainerSSHList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerSSH `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerSSH{}, &ContainerSSHList{})
}
