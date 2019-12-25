package operator

import (
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestEnsureDefaults(t *testing.T) {
	server := MySQLOperatorOpts{}
	server.EnsureDefaults()
	assertRequiredDefaults(t, server)
}

func assertRequiredDefaults(t *testing.T, s MySQLOperatorOpts) {
	if &s == nil {
		t.Error("MySQLOperatorServer: was nil, expected a valid configuration.")
	}
	if len(s.Hostname) <= 0 {
		t.Errorf("MySQLOperatorServer: expected a non-zero length hostname, got '%s'.", s.Hostname)
	}
	if &s.Images == nil {
		t.Error("MySQLOperatorServer.Images: was nil, expected a valid configuration.")
	}
	if s.Images.MySQLAgentImage != mysqlAgent {
		t.Errorf("MySQLOperatorServer.Images.MySQLAgentImage: was '%s', expected '%s'.", s.Images.MySQLAgentImage, mysqlAgent)
	}
	expectedDuration := v1.Duration{Duration: 43200000000000}
	if &s.MinResyncPeriod == nil {
		t.Errorf("MySQLOperatorServer.MinResyncPeriod: was nil, expected '%s'.", expectedDuration)
	}
	if s.MinResyncPeriod != expectedDuration {
		t.Errorf("MySQLOperatorServer.MinResyncPeriod: was '%s', expected '%s'.", s.MinResyncPeriod, expectedDuration)
	}
}

func TestEnsureDefaultsOverrideSafety(t *testing.T) {
	expected := mockMySQLOperatorOpts()
	ensured := mockMySQLOperatorOpts()
	ensured.EnsureDefaults()
	if expected != ensured {
		t.Errorf("MySQLOperatorOpts.EnsureDefaults() should not modify pre-configured values.")
	}
}

func mockMySQLOperatorOpts() MySQLOperatorOpts {
	return MySQLOperatorOpts{
		KubeConfig: "some-kube-config",
		Master:     "some-master",
		Hostname:   "some-hostname",
		Images: Images{
			MySQLAgentImage:         "some-agent-img",
			DefaultMySQLServerImage: "mysql/mysql-server",
		},
		MinResyncPeriod: v1.Duration{Duration: 42},
	}
}
