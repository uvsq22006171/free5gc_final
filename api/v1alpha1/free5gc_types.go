/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or im*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Free5GCSpec définit la configuration de Free5GC
type Free5GCSpec struct {
	Version string `json:"version,omitempty"`
	PFCP    PFCP   `json:"pfcp,omitempty"`
	GTPU    GTPU   `json:"gtpu,omitempty"`
	DNNList []DNN  `json:"dnnList,omitempty"`
}

// PFCP définit les paramètres du protocole PFCP
type PFCP struct {
	Addr           string `json:"addr"`
	NodeID         string `json:"nodeID"`
	RetransTimeout string `json:"retransTimeout"`
	MaxRetrans     int    `json:"maxRetrans"`
}

// GTPU définit la configuration des interfaces GTP-U
type GTPU struct {
	Forwarder string   `json:"forwarder"`
	IfList    []IfItem `json:"ifList"`
}

// IfItem représente une interface réseau
type IfItem struct {
	Addr string `json:"addr"`
	Type string `json:"type"`
}

// DNN définit un Data Network Name
type DNN struct {
	CIDR      string `json:"cidr"`
	DNN       string `json:"dnn"`
	NatifName string `json:"natifname"`
}

// Free5GCStatus contient l'état de Free5GC
type Free5GCStatus struct {
	Ready bool `json:"ready,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Free5GC est le schéma principal
type Free5GC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Free5GCSpec   `json:"spec,omitempty"`
	Status Free5GCStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Free5GCList contient une liste d'objets Free5GC
type Free5GCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Free5GC `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Free5GC{}, &Free5GCList{})
}
