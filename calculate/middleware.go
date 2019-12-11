package calculate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

type RequestMiddleware struct {
	Handler http.Handler
}

type Request struct {
	A int
	B int
}

type Error struct {
	Error string
}

func NewRequestMiddleware(handler http.Handler) http.Handler {
	return &RequestMiddleware{handler}
}

func (rm *RequestMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req Request
	w.Header().Set("Content-Type", "application/json")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(HandleError())
		return
	}

	// Checks if a || b is present because unmarshalling will fail
	// if it can't find both as key
	if err := json.Unmarshal(data, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(HandleError())
		return
	}

	switch {
	// Checks if a is not a negative integer
	case math.Signbit(float64(req.A)):
		w.WriteHeader(http.StatusBadRequest)
		w.Write(HandleError())
		return

	// Checks if b is not a negative integer
	case math.Signbit(float64(req.B)):
		w.WriteHeader(http.StatusBadRequest)
		w.Write(HandleError())
		return
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	rm.Handler.ServeHTTP(w, r)
}

func HandleError() []byte {
	errMessage := Error{Error: "Incorrect input"}
	c, err := json.Marshal(errMessage)
	if err != nil {
		fmt.Errorf("Unable to marshal request %+v", err)
	}
	return c
}
