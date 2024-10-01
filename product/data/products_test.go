package data

import "testing"



func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name: "Amr",
		Price: 22,
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}


}
