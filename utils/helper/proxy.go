
package helper

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetUrl string, pathPrefix string) *httputil.ReverseProxy {

	target, err := url.Parse(targetUrl)

	if err != nil {
		fmt.Println("Error Parsing Target URL:", err)
		return nil
	}

	return &httputil.ReverseProxy{
		Rewrite: func(req *httputil.ProxyRequest) {
			req.SetURL(target)

			if strings.HasPrefix(req.Out.URL.Path, pathPrefix) {
				req.Out.URL.Path = strings.TrimPrefix(
					req.Out.URL.Path,
					pathPrefix,
				)

				if req.Out.URL.Path == "" {
					req.Out.URL.Path = "/"
				}
			}

			fmt.Println("Target URL:", req.Out.URL.String())
		},

		ErrorHandler: func(
			w http.ResponseWriter,
			r *http.Request,
			err error,
		) {
			fmt.Println("Proxy error:", err)

			http.Error(
				w,
				"Failed to fetch from upstream service",
				http.StatusBadGateway,
			)
		},
	}
}
