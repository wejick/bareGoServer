// go:generate swagger generate spec
package helloWorld

// Helloword struct to containt Helloword object
// swagger:response HellowordResponse
type Helloword struct {
	d data
}
type data struct {
	Text string `json:"text"`
}
