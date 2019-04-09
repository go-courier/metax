package metax

import (
	"fmt"
	"net/http"
	"time"
)

type handler struct {
}

func (h *handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	meta := MetaFromContext(req.Context())
	fmt.Println(meta)
}

func ExampleMetaContextHttpHandler() {
	h := &handler{}

	for i := 0; i < 20; i++ {
		go func() {
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			MetaContextHttpHandler(h).ServeHTTP(http.ResponseWriter(nil), req)
		}()
	}

	time.Sleep(10 * time.Millisecond)
}
