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

// HandlerHelloWorldWithName handle helloword with name http request
func (hl *Helloword) HandlerHelloWorldWithName(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// swagger:operation GET /helloword/{name} hellowordWithName
	// ---
	// parameters :
	// - name : name
	//   in : path
	//   description : your name here bro
	//   required : true
	//   type : string
	// responses :
	//  '200' : HellowordResponse
	d := &data{}
	d.Text = hl.d.Text + " " + ps.ByName("name")
	httputil.ResponseJSON(d, writer)
}
