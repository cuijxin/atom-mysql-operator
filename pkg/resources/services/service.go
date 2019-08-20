package services

import (
	"atom-mysql-operator/pkg/apis/mysql/v1alpha1"
	"atom-mysql-operator/pkg/constants"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// NewForCluster will return a new headless Kubernetes service for a MySQL cluster
func NewForCluster(cluster *v1alpha1.Cluster) *corev1.Service {
	mysqlPort := corev1.ServicePort{Port: 3306}
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels:    map[string]string{constants.ClusterLabel: cluster.Name},
			Name:      cluster.Name,
			Namespace: cluster.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cluster, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    v1alpha1.ClusterCRDResourceKind,
				}),
			},
			Annotations: map[string]string{
				"service.alpha.kubernetes.io/tolerate-unready-endpoints": "true",
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{mysqlPort},
			Selector: map[string]string{
				constants.ClusterLabel: cluster.Name,
			},
			ClusterIP: corev1.ClusterIPNone,
		},
	}

	return svc
}
