package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// QuayEcosystemSpec defines the desired state of QuayEcosystem
type QuayEcosystemSpec struct {
	ImagePullSecretName string `json:"imagePullSecretName,omitempty"`
	Quay                Quay   `json:"quay"`
	Redis               Redis  `json:"redis"`
}

// QuayEcosystemStatus defines the observed state of QuayEcosystem
type QuayEcosystemStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// QuayEcosystem is the Schema for the quayecosystems API
// +k8s:openapi-gen=true
type QuayEcosystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QuayEcosystemSpec   `json:"spec,omitempty"`
	Status QuayEcosystemStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// QuayEcosystemList contains a list of QuayEcosystem
type QuayEcosystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuayEcosystem `json:"items"`
}

// Quay defines the properies of a deployment of Quay
type Quay struct {
	EnableNodePortService bool   `json:"enableNodePortService,omitempty"`
	Image                 string `json:"image,omitempty"`
	RouteHost             string `json:"routeHost,omitempty"`
	Replicas              *int32 `json:"replicas,omitempty"`
}

// Redis defines the properies of a deployment of Redis
type Redis struct {
	Skip     bool   `json:"skip,omitempty"`
	Image    string `json:"image,omitempty"`
	Replicas *int32 `json:"replicas,omitempty"`
}

func init() {
	SchemeBuilder.Register(&QuayEcosystem{}, &QuayEcosystemList{})
}