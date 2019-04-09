package metax

import (
	"net/http"

	"github.com/google/uuid"
)

func MetaContextHttpHandler(handler http.Handler) http.Handler {
	return &metaContextHttpHandler{
		nextHandler: handler,
	}
}

type metaContextHttpHandler struct {
	nextHandler http.Handler
}

func (h *metaContextHttpHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s := req.Header.Get("X-Request-ID")
	if s == "" {
		s = uuid.New().String()
	}
	if h.nextHandler != nil {
		h.nextHandler.ServeHTTP(rw, req.WithContext(ContextWithMeta(req.Context(), ParseMeta(s))))
	}
}
