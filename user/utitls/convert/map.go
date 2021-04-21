package convert

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func StructToMap(data interface{}) map[string]interface{} {
	bytes, err := json.Marshal(data)
	if err!=nil {
		log.Error(err)
		return nil
	}
	var returnMap = make(map[string]interface{})

	err = json.Unmarshal(bytes, &returnMap)
	if err!=nil {
		log.Error(err)
		return nil
	}

	return returnMap
}
