package utils

import (
	"encoding/json"
	"io/ioutil"
	"lo_project/internal/domain/model"
)

func GetDataModel() model.Order {
	file, _ := ioutil.ReadFile("model.json")

	data := model.Order{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}
