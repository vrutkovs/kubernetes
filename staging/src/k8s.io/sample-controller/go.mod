// This is a generated file. Do not edit directly.

module k8s.io/sample-controller

go 1.13

require (
	google.golang.org/appengine v1.6.5 // indirect
	k8s.io/api v0.19.0-rc.1
	k8s.io/apimachinery v0.19.0-rc.1
	k8s.io/client-go v0.19.0-rc.1
	k8s.io/code-generator v0.19.0-rc.1
	k8s.io/klog/v2 v2.2.0
)

replace (
	golang.org/x/net => golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/client-go => ../client-go
	k8s.io/code-generator => ../code-generator
	k8s.io/sample-controller => ../sample-controller
)
