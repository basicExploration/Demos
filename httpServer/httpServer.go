package httpServer



import (
	"io"
	"net/http"
)

func run(h http.ResponseWriter, v *http.Request) {
	io.WriteString(h, "<a>dfdfd</a>")
}

func init() {
  http.HandleFunc("/", run)
  http.ListenAndServe(":4000",nil)
}
 
