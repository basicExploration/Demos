//package huiter support a  richer router than go's official library net/http
package huiter

import (
	"net/http"
	"strings"
)

func (h *Huiter) ServeHTTP(writer http.ResponseWriter, reader *http.Request) {

}

var d Huiter = func(w http.ResponseWriter, r *http.Request) {
	pathValue := strings.Split(r.URL.Path, "/")
	for _, v := range pathValue {
		if strings.HasPrefix(v,":"){

		}
	}
}



