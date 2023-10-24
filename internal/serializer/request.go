package serializer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		SendResponseMessage(w, 400, "Invalid request body")
		return err
	}
	return nil
}

func jsonTag(v interface{}, fieldName string) (string, bool) {

	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		value = reflect.Indirect(value)
	}

	t := value.Type()
	sf, ok := t.FieldByName(fieldName)
	if !ok {
		return "", false
	}
	return sf.Tag.Lookup("json")
}

func ValidateRequestJson(w http.ResponseWriter, r *http.Request, validate *validator.Validate, v interface{}) error {
	if err := DecodeJSONBody(w, r, v); err != nil {
		return err
	}
	if err := validate.Struct(v); err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			SendResponseMessage(w, 400, "Error validating request body")
			return err
		}

		err := err.(validator.ValidationErrors)[0]

		if key, ok := jsonTag(v, err.Field()); ok {
			SendResponseMessage(w, 400, fmt.Sprintf(`%v field "%v"`, err.Tag(), key))
			return err
		}
		SendResponseMessage(w, 400, fmt.Sprintf(`%v field "%v"`, err.Tag(), err.Field()))
		return err
	}
	return nil
}
