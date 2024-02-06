package controller

import (
    "context"
    "fmt"

    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/types"
    "k8s.io/client-go/kubernetes"
    "k8s.io/klog/v2"

    trafficcapturev1 "github.com/yourusername/traffic-capture-controller/pkg/apis/trafficcapture/v1"
)

type TrafficCaptureController struct {
    kubeClientset *kubernetes.Clientset
}

func NewTrafficCaptureController(kubeClientset *kubernetes.Clientset) *TrafficCaptureController {
    return &TrafficCaptureController{
        kubeClientset: kubeClientset,
    }
}

func (c *TrafficCaptureController) AddSidecarToPod(ctx context.Context, tc *trafficcapturev1.TrafficCapture) error {
    // Fetch the Pod specified in the TrafficCapture resource
    pod, err := c.kubeClientset.CoreV1().Pods(tc.Namespace).Get(ctx, tc.Spec.PodName, metav1.GetOptions{})
    if err != nil {
        klog.Errorf("Failed to get pod %s in namespace %s: %v", tc.Spec.PodName, tc.Namespace, err)
        return err
    }

    // Check if the sidecar is already added
    for _, container := range pod.Spec.Containers {
        if container.Name == "traffic-forwarder-sidecar" {
            klog.Infof("Sidecar already exists in %s", tc.Spec.PodName)
            return nil
        }
    }

    // Define the sidecar container
    sidecarContainer := corev1.Container{
        Name:    "traffic-forwarder-sidecar",
        Image:   "your-sidecar-image:tag", // Specify your sidecar image
        Command: []string{"your-command"},  // Adjust this command to start your traffic forwarding
        Env: []corev1.EnvVar{
            {
                Name:  "POD_INTERFACE",
                Value: tc.Spec.PodInterface,
            },
            {
                Name:  "EXTERNAL_ENDPOINT",
                Value: fmt.Sprintf("%s:%d", tc.Spec.ExternalEndpoint, tc.Spec.ExternalPort),
            },
        },
        SecurityContext: &corev1.SecurityContext{
            Capabilities: &corev1.Capabilities{
                Add: []corev1.Capability{"NET_RAW"},
            },
        },
    }

    // Add the sidecar container to the pod
    pod.Spec.Containers = append(pod.Spec.Containers, sidecarContainer)

    // Update the Pod in the cluster
    _, err = c.kubeClientset.CoreV1().Pods(tc.Namespace).Update(ctx, pod, metav1.UpdateOptions{})
    if err != nil {
        klog.Errorf("Failed to update pod %s in namespace %s with sidecar: %v", tc.Spec.PodName, tc.Namespace, err)
        return err
    }

    klog.Infof("Added sidecar to pod %s in namespace %s", tc.Spec.PodName, tc.Namespace)
    return nil
}

// Implement other controller methods as necessary...

