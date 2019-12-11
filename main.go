package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/olugbokikisemiu/meant4Task/calculate"
)

func main() {
	router := httprouter.New()
	router.POST("/calculate", Index)
	m := calculate.NewRequestMiddleware(router)
	fmt.Println("Server started and listening on port :8989")
	http.ListenAndServe(":8989", m)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var req calculate.Request
	requestData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(calculate.HandleError())
		return
	}

	json.Unmarshal(requestData, &req)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, calculate.Calculate(req.A, req.B))
}
