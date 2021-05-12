package model

import (
	"encoding/json"
	"log"
)

type InboundObj map[string]interface{}
type InboundObjs []InboundObj

func InboundReadAll() []byte {
	var InboundObj InboundObj
	var InboundObjs InboundObjs
	records, err := JsonDB.ReadAll("Inbound")
	if err != nil {
		log.Println("Error", err)
	}
	for _, f := range records {
		if err := json.Unmarshal([]byte(f), &InboundObj); err != nil {
			log.Println("Error", err)
		}
		InboundObjs = append(InboundObjs, InboundObj)
	}
	if ret, err := json.Marshal(InboundObjs); err != nil {
		log.Println("Error", err)
	} else {
		return ret
	}
	return nil
}

func InboundWrite(tag string, in map[string]interface{}) {
	if err := JsonDB.Write("Inbound", tag, in); err != nil {
		log.Println("Error", err)
	}
}

func InboundRead(tag string) []byte {
	var InboundObj InboundObj
	if err := JsonDB.Read("Inbound", tag, &InboundObj); err != nil {
		log.Println("Error", err)
	}
	if ret, err := json.Marshal(InboundObj); err != nil {
		log.Println("Error", err)
	} else {
		return ret
	}
	return nil
}

func InboundDel(tag string) {
	if err := JsonDB.Delete("Inbound", tag); err != nil {
		log.Println("Error", err)
	}
}

func InboundApply(tag string) {

}
