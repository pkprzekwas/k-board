package namespace

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Service struct {
	Client v1.NamespaceInterface
}

func (s *Service) List() (*corev1.NamespaceList, error) {
	listOpts := metav1.ListOptions{}
	nsList, err := s.Client.List(listOpts)
	if err != nil {
		return nil, err
	}
	return nsList, nil
}

