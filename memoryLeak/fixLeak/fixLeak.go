package main

import (
	"fmt"
	"os"

	"github.com/iamyeswc/libxml2"
	"github.com/iamyeswc/libxml2/xsd"
)

const xsdSchemaPath = `./svg_1-2_ps.xsd`

func xsdValidator(svgPath string) error {
	xsdData, err := os.ReadFile(xsdSchemaPath)
	if err != nil {
		return err
	}
	schema, err := xsd.Parse(xsdData)
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
