package secrets

import (
	"atom-mysql-operator/pkg/apis/mysql/v1alpha1"
	"atom-mysql-operator/pkg/constants"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// NewMysqlRootPassword returns a Kubernetes secret containing a
// generated MySQL root password.
func NewMysqlRootPassword(cluster *v1alpha1.Cluster) *corev1.Secret {
	CreateSecret := RandomAlphanumericString(16)
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{constants.ClusterLabel: cluster.Name},
			Name:   GetRootPasswordSecretName(cluster),
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cluster, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    v1alpha1.ClusterCRDResourceKind,
				}),
			},
			Namespace: cluster.Namespace,
		},
		Data: map[string][]byte{"password": []byte(CreateSecret)},
	}
	return secret
}

// GetRootPasswordSecretName returns the root password secret name for the
// given mysql cluster.
func GetRootPasswordSecretName(cluster *v1alpha1.Cluster) string {
	return fmt.Sprintf("%s-root-password", cluster.Name)
}
