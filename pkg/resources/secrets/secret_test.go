package secrets

import (
	"atom-mysql-operator/pkg/apis/mysql/v1alpha1"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetRootPasswordSecretName(t *testing.T) {
	cluster := &v1alpha1.Cluster{
		ObjectMeta: metav1.ObjectMeta{Name: "example-cluster"},
		Spec:       v1alpha1.ClusterSpec{},
	}

	actual := GetRootPasswordSecretName(cluster)

	if actual != "example-cluster-root-password" {
		t.Errorf("Expected example-cluster-root-password but got %s", actual)
	}
}
