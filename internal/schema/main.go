package schema

import (
	"errors"
	"log"

	"github.com/xeipuuv/gojsonschema"
)

// Loader holds the context to the schema
type Loader struct {
	Ref map[string]*gojsonschema.Schema
}

// Validate validates the schema based on the name and return a list of errors if fail
func (l Loader) Validate(ref string, doc interface{}) []error {
	var errs []error
	schema, ok := l.Ref[ref]
	if !ok {
		errs = append(errs, errors.New("schema does not exist"))
		return errs
	}
	result, err := schema.Validate(gojsonschema.NewGoLoader(doc))
	if err != nil {
		errs = append(errs, err)
		return errs
	}
	if result.Valid() {
		return errs
	}
	for _, desc := range result.Errors() {
		errs = append(errs, errors.New(desc.Description()))
	}
	return errs
}

// New creates a new reference to the json schema
func New(schemas map[string]string) *Loader {
	var loader Loader
	loader.Ref = make(map[string]*gojsonschema.Schema)

	for key, value := range schemas {
		s, err := gojsonschema.NewSchema(gojsonschema.NewReferenceLoader(value))
		if err != nil {
			log.Fatal(err.Error())
		}
		loader.Ref[key] = s
	}

	return &loader
}
