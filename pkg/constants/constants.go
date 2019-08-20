package constants

// ClusterLabel is applied to all components of a MySQL ClusterLabel
const ClusterLabel = "v1alpha1.mysql.oracle.com/cluster"

// MySQLOperatorVersionLabel denotes the version of the MySQLOperator and
// MySQLOperatorAgent running in the cluster.
const MySQLOperatorVersionLabel = "v1alpha1.mysql.oracle.com/version"

// LabelClusterRole specifies the role of a Pod within a cluster.
const LabelClusterRole = "v1alpha1.mysql.oracle.com/role"

// ClusterRolePrimary denotes a primary InnoDB cluster member.
const ClusterRolePrimary = "primary"

// ClusterRoleSecondary denotes a secondary InnoDB cluster member.
const ClusterRoleSecondary = "secondary"
