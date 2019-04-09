package metax

import (
	"net/http"
)

func MetaContextRoundTripper(rt http.RoundTripper) http.RoundTripper {
	return &metaContextRoundTripper{
		next: rt,
	}
}

type metaContextRoundTripper struct {
	next http.RoundTripper
}

func (h *metaContextRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	meta := MetaFromContext(req.Context())
	req.Header.Set("X-Request-ID", meta.String())
	if h.next != nil {
		return h.next.RoundTrip(req)
	}
	return nil, nil
}
