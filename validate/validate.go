package validate

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func Validate(json interface{}) error {
	validate := validator.New()
	err := validate.Struct(json)
	if err != nil {
		log.Println(err.Error())
		return err
	} else {
		return nil
	}
}
