package storage

import (
	"encoding/json"
	"os"
)

func SaveToJSON(data interface{}, fileName string) error {
	jsonDATA,err:=json.MarshalIndent(data,""," ")
	if  err!= nil {
		return  err
	}
	return os.WriteFile(fileName,jsonDATA,0644)
}