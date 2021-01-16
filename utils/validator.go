package utils

import (
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

// Inputvalidate Validates the data
func Inputvalidate(dataset interface{}) (bool, map[string]string) {
	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(dataset)
	if err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}
		errors := make(map[string]string)
		refected := reflect.ValueOf(dataset)
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := refected.Type().FieldByName(err.StructField())
			var name string
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}
			switch err.Tag() {
			case "required":
				errors["Error"] = "The " + name + " cant be empty"
				break
			case "email":
				errors["Error"] = "The " + name + " Not Valid"
				break
			case "eqfield":
				errors["Error"] = "The " + name + " Doesnt match"
			default:
				errors["Error"] = "Validation Failed: " + name
			}

		}
		return false, errors

	}
	return true, nil

}
