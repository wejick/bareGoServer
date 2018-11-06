package helloWorld

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wejick/bareGoServer/pkg/httputil"
)

//NewHelloWorld create new instance of helloword
func NewHelloWorld(hellowordString string) *Helloword {
	hl := &Helloword{}

	hl.d.Text = hellowordString

	return hl
}

//HandlerHelloWorld handle helloword http request
func (hl *Helloword) HandlerHelloWorld(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httputil.ResponseJSON(hl.d, writer)
}

//HandlerHelloWorldWithName handle helloword with name http request
func (hl *Helloword) HandlerHelloWorldWithName(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	d := &data{}
	d.Text = hl.d.Text + " " + ps.ByName("name")
	httputil.ResponseJSON(d, writer)
}
