package huiter

import "net/http"

//huiter the core struct in this package.
type Huiter func(w http.ResponseWriter,r *http.Request)
