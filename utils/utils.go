package utils

import (
	"encoding/json"
	"log"
)

//InterfaceToBytes this function return an interface to []byte
func InterfaceToBytes(data interface{}) []byte {
	purchaseBytes, errMarshal := json.Marshal(data)
	if errMarshal != nil {
		log.Panicf("Can not convert interface %#v to bytes", data)
	}
	return purchaseBytes
}
