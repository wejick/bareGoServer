package helloWorld

type (
	//Helloword struct to containt Helloword object
	Helloword struct {
		d data
	}
	data struct {
		Text string `json:"text"`
	}
)
