package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	gjson "github.com/gorilla/rpc/json"
)

type Args struct {
	ID string
}

type Employee struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

type JSONServer struct{}

func (s *JSONServer) Get(r *http.Request, args *Args, reply *Employee) error {
	var employee []Employee

	b, err := ioutil.ReadFile("./employee.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &employee); err != nil {
		return err
	}

	for _, v := range employee {
		if v.ID == args.ID {
			*reply = v
			break
		}
	}

	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(gjson.NewCodec(), "application/json")
	s.RegisterService(new(JSONServer), "")

	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":8080", r)
}
