package metax

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

type handler struct {
}

func (h *handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	meta := MetaFromContext(req.Context())
	fmt.Println(meta)
}

func TestMetaContextHttpHandler(t *testing.T) {
	h := &handler{}

	for i := 0; i < 20; i++ {
		go func() {
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			MetaContextHttpHandler(h).ServeHTTP(http.ResponseWriter(nil), req)
		}()
	}

	time.Sleep(10 * time.Millisecond)
}
