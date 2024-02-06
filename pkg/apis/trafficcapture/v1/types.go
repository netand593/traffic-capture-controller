package v1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TrafficCapture specifies a pod and interface to capture traffic from, along with the destination for the traffic.
type TrafficCapture struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec TrafficCaptureSpec `json:"spec,omitempty"`
}

type TrafficCaptureSpec struct {
    PodName          string `json:"podName"`
    PodInterface     string `json:"podInterface"`
    ExternalEndpoint string `json:"externalEndpoint"`
    ExternalPort     int    `json:"externalPort"`
}

// TrafficCaptureList is a list of TrafficCapture resources.
type TrafficCaptureList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []TrafficCapture `json:"items"`
}

