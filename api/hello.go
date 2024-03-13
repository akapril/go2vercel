package vercel2go

import (
	"fmt"
	"net/http"
)

func Hello(rw http.ResponseWriter, req *http.Request) {
	msg := req.URL.Query().Get("msg")
	fmt.Fprint(rw, msg)
}
