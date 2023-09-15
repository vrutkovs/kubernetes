package apiserverconfig

import (
	"net/http"

	"k8s.io/apimachinery/pkg/util/sets"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	genericfilters "k8s.io/apiserver/pkg/server/filters"
)

var (
	kubeLongRunningFunc = genericfilters.BasicLongRunningRequestCheck(
		sets.NewString("watch", "proxy"),
		sets.NewString("attach", "exec", "proxy", "log", "portforward"),
	)
)

func IsLongRunningRequest(r *http.Request, requestInfo *apirequest.RequestInfo) bool {
	if requestInfo.APIPrefix == "apis" && requestInfo.APIGroup == "build.openshift.io" && requestInfo.APIVersion == "v1" && requestInfo.Resource == "buildconfigs" && requestInfo.Subresource == "instantiatebinary" {
		return true
	}
	if requestInfo.APIPrefix == "apis" && requestInfo.APIGroup == "image.openshift.io" && requestInfo.APIVersion == "v1" && requestInfo.Resource == "imagestreamimports" {
		return true
	}
	if kubeLongRunningFunc(r, requestInfo) {
		return true
	}
	return false
}
