package utils

import (
	"encoding/json"
)

func JsonEncode(v interface{}) []byte {
	res, err := json.Marshal(v)
	CheckError(err)
	return res
}

func JsonDecode(j []byte) (v interface{}) {
	var res interface{}
	err := json.Unmarshal(j, &res)
	CheckError(err)
	return res
}
