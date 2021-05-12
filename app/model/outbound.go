package model

import (
	"encoding/json"
	"log"
)

type OutboundObj map[string]interface{}

type OutboundObjs []OutboundObj

func OutboundReadAll() []byte {
	var OutboundObj OutboundObj
	var OutboundObjs OutboundObjs
	records, err := JsonDB.ReadAll("Outbound")
	if err != nil {
		log.Println("Error", err)
	}
	for _, f := range records {
		if err := json.Unmarshal([]byte(f), &OutboundObj); err != nil {
			log.Println("Error", err)
		}
		OutboundObjs = append(OutboundObjs, OutboundObj)
	}
	if ret, err := json.Marshal(OutboundObjs); err != nil {
		log.Println("Error", err)
	} else {
		return ret
	}
	return nil
}

func OutboundWrite(tag string, in map[string]interface{}) {
	if err := JsonDB.Write("Outbound", tag, in); err != nil {
		log.Println("Error", err)
	}
}

func OutboundRead(tag string) []byte {
	var OutboundObj OutboundObj
	if err := JsonDB.Read("Outbound", tag, &OutboundObj); err != nil {
		log.Println("Error", err)
	}
	if ret, err := json.Marshal(OutboundObj); err != nil {
		log.Println("Error", err)
	} else {
		return ret
	}
	return nil
}

func OutboundDel(tag string) {
	if err := JsonDB.Delete("Outbound", tag); err != nil {
		log.Println("Error", err)
	}
}

func OutboundApply(tag string) {
}
