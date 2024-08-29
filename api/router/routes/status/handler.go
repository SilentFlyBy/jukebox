package status

import "net/http"

func GetHealthStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte{})
}
