// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// This class is intended to host all CRUD calls for testing secrets primitive resources

// CreateSAFromName generate a vanilla service-account
// with the provided name
func (t *TestBuild) CreateSAFromName(saName string) error {
	client := t.Clientset.CoreV1().ServiceAccounts(t.Namespace)
	_, err := client.Create(&corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: saName,
		}})
	if err != nil {
		return err
	}
	return nil
}

// GetSA retrieves an existing service-account by name
func (t *TestBuild) GetSA(saName string) (*v1.ServiceAccount, error) {
	client := t.Clientset.CoreV1().ServiceAccounts(t.Namespace)
	sa, err := client.Get(saName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return sa, nil
}
