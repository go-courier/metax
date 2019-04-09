package metax

import (
	"fmt"
	"net/http"
	"testing"
)

type roundTrigger struct {
}

func (h *roundTrigger) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, nil
}

func TestLogIDRoundTripper(t *testing.T) {
	for i := 0; i < 1000; i++ {
		meta := Meta{
			"id": {fmt.Sprintf("%d", i)},
		}

		go func() {
			req, _ := http.NewRequest(http.MethodGet, "/", nil)

			req = req.WithContext(ContextWithMeta(req.Context(), meta))

			MetaContextRoundTripper(&roundTrigger{}).RoundTrip(req)

			if req.Header.Get("X-Request-ID") != meta.String() {
				t.Fatal()
			}
		}()
	}
}
