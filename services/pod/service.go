package pod

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Service struct {
	Client v1.CoreV1Interface
}

func (s *Service) List(namespace string) (*corev1.PodList, error) {
	listOpts := metav1.ListOptions{}
	podList, err := s.Client.Pods(namespace).List(listOpts)
	if err != nil {
		return nil, err
	}
	return podList, nil
}
