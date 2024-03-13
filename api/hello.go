package vercel2go

import (
	"fmt"
	"net/http"
)

func hello(rw http.ResponseWriter, req *http.Request) {
	msg := req.URL.Query().Get("msg")
	fmt.Fprint(rw, msg)
}
