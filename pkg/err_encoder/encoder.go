package err_encoder

import (
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdhttp "net/http"
)

func NewHTTPError(code int, field string, detail string) *HTTPError {
	return &HTTPError{
		Code: code,
		Errors: map[string][]string{
			field: {detail},
		},
	}
}

type HTTPError struct {
	Errors map[string][]string `json:"error"`
	Code   int                 `json:"-"`
}

func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if se := new(HTTPError); errors.As(err, &se) {
		return se
	}
	fmt.Println(err)
	//default
	return &HTTPError{
		Code: 500,
		Errors: map[string][]string{
			"error": {"system error"},
		},
	}
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError %d", e.Code)
}

func ErrorEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
	se := FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.WriteHeader(se.Code)
	_, _ = w.Write(body)
}
