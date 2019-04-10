package metax

import (
	"errors"
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
	req.Header.Set("X-Request-ID", MetaFromContext(req.Context()).String())
	if h.next == nil {
		panic(errors.New("need use before other RoundTripper"))
	}
	return h.next.RoundTrip(req)
}
