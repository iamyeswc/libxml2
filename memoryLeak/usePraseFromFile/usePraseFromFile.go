package main

import (
	"fmt"
	"os"

	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xsd"
)

const xsdSchemaPath = `./svg_1-2_ps.xsd`

func xsdValidator(svgPath string) error {
	schema, err := xsd.ParseFromFile(xsdSchemaPath)
	if err != nil {
		return err
	}
	defer schema.Free()

	svgData, err := os.ReadFile(svgPath)
	if err != nil {
		return err
	}

	doc, err := libxml2.ParseString(string(svgData))
	if err != nil {
		return err
	}
	defer doc.Free()

	err = schema.Validate(doc)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	svgPath := `../bimi-sq.svg`
	if err := xsdValidator(svgPath); err != nil {
		fmt.Println("Error validating SVG: %v", err)
		panic(err)
	}
}
