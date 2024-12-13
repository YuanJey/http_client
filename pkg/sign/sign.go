package sign

import "net/http"

type RequestSign interface {
	Sign(request *http.Request)
}
