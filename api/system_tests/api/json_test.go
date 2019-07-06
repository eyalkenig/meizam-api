package api_test

import "encoding/json"

func toJSON(jsonMap map[string]interface{}) []byte {
	json, err := json.Marshal(jsonMap)
	if err != nil {
		panic(err)
	}
	return json
}

func fromJSON(bytes []byte) map[string]interface{} {
	var jsonMap map[string]interface{}
	err := json.Unmarshal(bytes, &jsonMap)
	if err != nil {
		panic(err)
	}
	return jsonMap
}
