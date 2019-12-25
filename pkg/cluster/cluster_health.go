package cluster

import (
	"atom-mysql-operator/pkg/cluster/innodb"
	"sync"

	"github.com/heptiolabs/healthcheck"
	"github.com/pkg/errors"
)

var (
	status      *innodb.ClusterStatus
	statusMutex sync.Mutex
)

// SetStatus sets the status of the local mysql cluster. The cluster manager
// controller is responsible for updating.
func SetStatus(new *innodb.ClusterStatus) {
	statusMutex.Lock()
	defer statusMutex.Unlock()
	status = new.DeepCopy()
}

// GetStatus fetches a copy of the latest cluster status.
func GetStatus() *innodb.ClusterStatus {
	statusMutex.Lock()
	defer statusMutex.Unlock()
	if status == nil {
		return nil
	}
	return status.DeepCopy()
}

// NewHealthCheck constructs a healthcheck for the local instance which checks
// cluster status using mysqlsh.
func NewHealthCheck() (healthcheck.Check, error) {
	instance, err := NewLocalInstance()
	if err != nil {
		return nil, errors.Wrap(err, "getting local mysql instance")
	}

	return func() error {
		s := GetStatus()
		if s == nil || s.GetInstanceStatus(instance.Name()) != innodb.InstanceStatusOnline {
			return errors.New("database still requires management")
		}
		return nil
	}, nil
}
