package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// This package will auto register types with the Kubernetes API

var (
	// SchemeBuilder collects the scheme builder functions for the MySQL
	// Operator API.
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)

	// AddToScheme applies the SchemeBuilder functions to a specified scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

// GroupName is the group name for the MySQL Operator API.
const GroupName = "mysql.oracle.com"

// SchemeGroupVersion is the GroupVersion for the MySQL Operator API.
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

const (
	// ClusterCRDResourceKind is the Kind of a Cluster.
	ClusterCRDResourceKind = "Cluster"
	// BackupCRDResourceKind is the Kind of a Backup.
	BackupCRDResourceKind = "Backup"
	// RestoreCRDResourceKind is the Kind of a Restore.
	RestoreCRDResourceKind = "Restore"
	// BackupScheduleCRDResourceKind is the Kind of a BackupSchedule.
	BackupScheduleCRDResourceKind = "BackupSchedule"
)

// Resource gets a MySQL Operator GroupResource for a specified resource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// addKnownTypes adds the set of types defined in this package to the supplied
// scheme.
func addKnownTypes(s *runtime.Scheme) error {
	s.AddKnownTypes(SchemeGroupVersion,
		&Cluster{},
		&ClusterList{},
		&Backup{},
		&BackupList{},
		&Restore{},
		&RestoreList{},
		&BackupSchedule{},
		&BackupScheduleList{})
	metav1.AddToGroupVersion(s, SchemeGroupVersion)
	return nil
}
