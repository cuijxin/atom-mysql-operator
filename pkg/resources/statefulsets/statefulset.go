package statefulsets

import (
	"atom-mysql-operator/pkg/apis/mysql/v1alpha1"
	"atom-mysql-operator/pkg/resources/secrets"
	"fmt"
	"strconv"
	"strings"

	"github.com/blang/semver"
	v1 "k8s.io/api/core/v1"
)

const (
	// MySQLServerName is the static name of all 'mysql(-server)' containers.
	MySQLServerName = "mysql"
	// MySQLAgentName is the static name of all 'mysql-agent' containers.
	MySQLAgentName = "mysql-agent"
	// MySQLAgentBasePath defines the volume mount path for the MySQL agent
	MySQLAgentBasePath = "/var/lib/mysql-agent"

	mySQLbackupVolumeName = "mysqlbackupvolume"
	mySQLVolumeName       = "mysqlvolume"
	mySQLSSLVolumeName    = "mysqlsslvolume"

	replicationGroupPort = 13306

	minMysqlVersionWithGroupExitStateArgs = "8.0.12"
)

func volumeMounts(cluster *v1alpha1.Cluster) []v1.VolumeMount {
	var mounts []v1.VolumeMount

	name := mySQLVolumeName
	if cluster.Spec.VolumeClaimTemplate != nil {
		name = cluster.Spec.VolumeClaimTemplate.Name
	}

	mounts = append(mounts, v1.VolumeMount{
		Name:      name,
		MountPath: "/var/lib/mysql",
		SubPath:   "mysql",
	})

	backupName := mySQLbackupVolumeName
	if cluster.Spec.BackupVolumeClaimTemplate != nil {
		backupName = cluster.Spec.BackupVolumeClaimTemplate.Name
	}
	mounts = append(mounts, v1.VolumeMount{
		Name:      backupName,
		MountPath: MySQLAgentBasePath,
		SubPath:   "mysql",
	})

	// A user may explicitly define a my.cnf configuration file for
	// their MySQL cluster.
	if cluster.RequiresConfigMount() {
		mounts = append(mounts, v1.VolumeMount{
			Name:      cluster.Name,
			MountPath: "/etc/my.cnf",
			SubPath:   "my.cnf",
		})
	}

	if cluster.RequiresCustomSSLSetup() {
		mounts = append(mounts, v1.VolumeMount{
			Name:      mySQLSSLVolumeName,
			MountPath: "/etc/ssl/mysql",
		})
	}

	return mounts
}

func clusterNameEnvVar(cluster *v1alpha1.Cluster) v1.EnvVar {
	return v1.EnvVar{Name: "MYSQL_CLUSTER_NAME", Value: cluster.Name}
}

func namespaceEnvVar() v1.EnvVar {
	return v1.EnvVar{
		Name: "POD_NAMESPACE",
		ValueFrom: &v1.EnvVarSource{
			FieldRef: &v1.ObjectFieldSelector{
				FieldPath: "metadata.namespace",
			},
		},
	}
}

func replicationGroupSeedsEnvVar(replicationGroupSeeds string) v1.EnvVar {
	return v1.EnvVar{
		Name:  "REPLICATION_GROUP_SEEDS",
		Value: replicationGroupSeeds,
	}
}

func multiMasterEnvVar(enabled bool) v1.EnvVar {
	return v1.EnvVar{
		Name:  "MYSQL_CLUSTER_MULTI_MASTER",
		Value: strconv.FormatBool(enabled),
	}
}

// Returns the MySQL_ROOT_PASSWORD environment variable
// If a user specifies a secret in the spec we use that
// else we create a secret with a random password
func mysqlRootPassword(cluster *v1alpha1.Cluster) v1.EnvVar {
	var secretName string
	if cluster.RequiresSecret() {
		secretName = secrets.GetRootPasswordSecretName(cluster)
	} else {
		secretName = cluster.Spec.RootPasswordSecret.Name
	}

	return v1.EnvVar{
		Name: "MYSQL_ROOT_PASSWORD",
		ValueFrom: &v1.EnvVarSource{
			SecretKeyRef: &v1.SecretKeySelector{
				LocalObjectReference: v1.LocalObjectReference{
					Name: secretName,
				},
				Key: "password",
			},
		},
	}
}

func getReplicationGroupSeeds(name string, members int) string {
	seeds := []string{}
	for i := 0; i < members; i++ {
		seeds = append(seeds, fmt.Sprintf("%[1]s-%[2]d.%[1]s:%[3]d", name, i, replicationGroupPort))
	}
	return strings.Join(seeds, ",")
}

func checkSupportGroupExitStateArgs(deployingVersion string) (supportedVer bool) {
	defer func() {
		if r := recover(); r != nil {

		}
	}()

	supportedVer = false

	ver := semver.New(deployingVersion)
	minVer := semver.New(minMysqlVersionWithGroupExitStateArgs)

	if ver.LessThan(*minVer) {
		return
	}

	supportedVer = true
	return
}

// Builds the MySQl operator container for a cluster.
// The 'mysqlImage' paramter is the image name of the mysql server to use with
// no version information.. e.g. 'mysql/mysql-server'
func mysqlServerContainer(cluster *v1alpha1.Cluster, mysqlServerImage string, rootPassword v1.EnvVar, members int, baseServerID uint32) v1.Container {
	args := []string{
		"--server_id=$(expr $base + $index)",
		"--datadir=/var/lib/mysql",
		"--user=mysql",
		"--gtid_mode=ON",
		"--log-bin",
		"--binlog_checksum=NONE",
	}
}
