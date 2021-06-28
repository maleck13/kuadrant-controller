/*
Copyright 2021 Red Hat, Inc.

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
	v12 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// APIProductSpec defines the desired state of APIProduct
type APIProductSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Information    ProductInformation `json:"information"`
	Routing        Routing            `json:"routing"`
	SecurityScheme []*SecurityScheme  `json:"securityScheme"`
	APIs           []*APISelector     `json:"APIs"`
}

// APIProductStatus defines the observed state of APIProduct
type APIProductStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// TODO(jmprusi): use proper conditions to give feedback to the user
	ObservedGen int64 `json:"observedgen"`
	Ready       bool  `json:"ready"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// APIProduct is the Schema for the apiproducts API
type APIProduct struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   APIProductSpec   `json:"spec,omitempty"`
	Status APIProductStatus `json:"status,omitempty"`
}

type ProductInformation struct {
	Description string `json:"description"`
	Owner       string `json:"owner"`
}

type APIKeyAuthCredentials struct {
	LabelSelectors map[string]string `json:"labelSelectors"`
}

type Routing struct {
	Hosts  []string `json:"hosts"`
	Expose bool     `json:"expose"`
}

type OpenIDConnectAuthCredentials struct {
	Endpoint string `json:"endpoint"`
}

type Destination struct {
	Schema                string `json:"schema,omitempty"`
	*v12.ServiceReference `json:"serviceReference"`
}

type TLSConfig struct {
	PlainHTTP     string `json:"plainHTTP"`
	TLSSecretName string `json:"tlsSecretName"`
}

type APISelector struct {
	Name      string  `json:"name"`
	Namespace string  `json:"namespace"`
	Tag       string  `json:"tag"`
	Mapping   Mapping `json:"mapping,omitempty"`
}

type Mapping struct {
	Prefix string `json:"prefix"`
}

//+kubebuilder:object:root=true

// APIProductList contains a list of APIProduct
type APIProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIProduct `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIProduct{}, &APIProductList{})
}
