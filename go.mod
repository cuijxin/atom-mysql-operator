module atom-mysql-operator

go 1.12

require (
	github.com/blang/semver v3.5.1+incompatible
	// github.com/blang/semver v3.5.1+incompatible
	github.com/coreos/go-semver v0.3.0
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/heptiolabs/healthcheck v0.0.0-20180807145615-6ff867650f40
	github.com/oracle/mysql-operator v0.0.0-20190515081336-9aebcc37a080
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.3.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.4.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	golang.org/x/sys v0.0.0-20191224085550-c709ea063b76 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.2.7
	k8s.io/api v0.0.0-20190814101207-0772a1bdf941
	k8s.io/apimachinery v0.0.0-20190814100815-533d101be9a6
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20190809000727-6c36bc71fc4a // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.0+incompatible
)
