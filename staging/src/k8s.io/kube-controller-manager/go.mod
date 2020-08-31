// This is a generated file. Do not edit directly.

module k8s.io/kube-controller-manager

go 1.15

require (
	k8s.io/apimachinery v0.19.0
	k8s.io/component-base v0.0.0
)

replace (
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/client-go => ../client-go
	k8s.io/component-base => ../component-base
	k8s.io/klog/v2 => k8s.io/klog/v2 v2.2.0
	k8s.io/kube-controller-manager => ../kube-controller-manager
)
