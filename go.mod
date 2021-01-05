module github.com/openshift/client-go

go 1.13

require (
	github.com/openshift/api v0.0.0-20201203102015-275406142edb
	github.com/spf13/pflag v1.0.5
	gopkg.in/yaml.v2 v2.3.0 // indirect
	k8s.io/api v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v0.19.2
	k8s.io/code-generator v0.19.2
	k8s.io/klog/v2 v2.3.0 // indirect
)

replace github.com/openshift/api => github.com/jhadvig/api v0.0.0-20210104163517-3e2cb2ea1f59
