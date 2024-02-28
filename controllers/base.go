package controllers

import (
	"example/go-crud/initializers"
	"example/go-crud/utils"
)

func CheckByID(id string, model interface{}) error {
	int_id, err := utils.ConvertStringToUint(id)
	if err != nil {
		return err
	}

	err = initializers.DB.First(model, int_id).Error
	if err != nil {
		return err
	}

	return nil
}
