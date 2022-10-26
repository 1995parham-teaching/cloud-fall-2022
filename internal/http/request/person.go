package request

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Person struct {
	Name   string `json:"name"`
	Family string `json:"family"`
	Age    int    `json:"age"`
}

func (p Person) Validate() error {
	if err := validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(1, 0), is.UTFLetterNumeric),
		validation.Field(&p.Family, validation.Required, validation.Length(1, 0), is.UTFLetterNumeric),
		validation.Field(&p.Age, validation.Required, validation.Min(0)),
	); err != nil {
		return fmt.Errorf("person request data validation failed %w", err)
	}

	return nil
}
