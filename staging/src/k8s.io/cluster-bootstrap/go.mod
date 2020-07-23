// This is a generated file. Do not edit directly.

module k8s.io/cluster-bootstrap

go 1.13

require (
	github.com/stretchr/testify v1.4.0
	golang.org/x/crypto v0.0.0-20200220183623-bac4c82f6975 // indirect
	gopkg.in/square/go-jose.v2 v2.2.2
	k8s.io/api v0.19.0-rc.1
	k8s.io/apimachinery v0.19.0-rc.1
	k8s.io/klog/v2 v2.2.0
)

replace (
	golang.org/x/net => golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/cluster-bootstrap => ../cluster-bootstrap
)
