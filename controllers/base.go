package controllers

import (
	"example/go-crud/initializers"
	"example/go-crud/utils"
	"fmt"
	"reflect"
)

// getModelName returns the name of the type of the provided value.
func getModelName(model interface{}) string {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}

func CheckByID(id string, model interface{}) error {
	int_id, err := utils.ConvertStringToUint(id)
	if err != nil {
		return err
	}

	err = initializers.DB.First(model, int_id).Error
	if err != nil {
		modelName := getModelName(model)
		return fmt.Errorf("%s with ID: %d not found", modelName, int_id)
	}

	return nil
}
