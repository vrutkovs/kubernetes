// This is a generated file. Do not edit directly.

module k8s.io/metrics

go 1.13

require (
	github.com/gogo/protobuf v1.3.1
	github.com/stretchr/testify v1.4.0
	google.golang.org/appengine v1.6.5 // indirect
	k8s.io/api v0.19.0-rc.1
	k8s.io/apimachinery v0.19.0-rc.1
	k8s.io/client-go v0.19.0-rc.1
	k8s.io/code-generator v0.19.0-rc.1
)

replace (
	golang.org/x/net => golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/client-go => ../client-go
	k8s.io/code-generator => ../code-generator
	k8s.io/metrics => ../metrics
)
