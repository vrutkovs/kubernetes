// This is a generated file. Do not edit directly.

module k8s.io/cloud-provider

go 1.13

require (
	github.com/google/go-cmp v0.4.0
	github.com/stretchr/testify v1.4.0
	k8s.io/api v0.19.0-rc.1
	k8s.io/apimachinery v0.19.0-rc.1
	k8s.io/client-go v0.19.0-rc.1
	k8s.io/component-base v0.0.0
	k8s.io/klog/v2 v2.2.0
	k8s.io/utils v0.0.0-20200619165400-6e3d28b6ed19
)

replace (
	golang.org/x/net => golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/client-go => ../client-go
	k8s.io/cloud-provider => ../cloud-provider
	k8s.io/component-base => ../component-base
)
