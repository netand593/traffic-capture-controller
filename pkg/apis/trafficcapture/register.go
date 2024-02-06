package trafficcapture

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/apimachinery/pkg/runtime/schema"

    "github.com/yourusername/traffic-capture-controller/pkg/apis/trafficcapture/v1"
)

const (
    GroupName = "trafficcaptures.mydomain.org"
    Version   = "v1"
)

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: Version}

func Kind(kind string) schema.GroupKind {
    return SchemeGroupVersion.WithKind(kind).GroupKind()
}

func Resource(resource string) schema.GroupResource {
    return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
    SchemeBuilder = runtime.NewSchemeBuilder(v1.AddKnownTypes)
    AddToScheme   = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
    scheme.AddKnownTypes(SchemeGroupVersion,
        &v1.TrafficCapture{},
        &v1.TrafficCaptureList{},
    )
    metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
    return nil
}

