package exception

import (
	"encoding/json"
	"net/http"
)

func ErrorHandler(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			switch e := err.(type) {
			case *BadRequestError:
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(e)
			default:
				http.Error(
					w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError,
				)
			}
		}
	}
}
