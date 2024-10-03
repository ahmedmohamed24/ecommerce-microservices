package utils

import (
	"encoding/json"

	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/datatypes"
)

func ConvertDatatypesJSONToStruct(jsonData datatypes.JSON) (*structpb.Struct, error) {
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &jsonMap); err != nil {
		return nil, err
	}

	jsonStruct, err := structpb.NewStruct(jsonMap)
	if err != nil {
		return nil, err
	}

	return jsonStruct, nil
}
