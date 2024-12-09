package main

import (
	"encoding/json"
	"os"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
)

func main() {
	input, err := os.ReadFile("./openapi/book/v1/book_controller.swagger.json")
	if err != nil {
		panic(err)
	}

	var doc2 openapi2.T

	if err = json.Unmarshal(input, &doc2); err != nil {
		panic(err)
	}

	doc3, err := openapi2conv.ToV3(&doc2)
	if err != nil {
		panic(err)
	}

	output, err := json.MarshalIndent(doc3, "", "  ")
	if err != nil {
		panic(err)
	}
	if err = os.WriteFile("./openapi/book/v1/book_controller.openapi.json", output, 0644); err != nil {
		panic(err)
	}
}
