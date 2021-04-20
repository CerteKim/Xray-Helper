package inbound

import (
	"encoding/json"
	"log"

	"github.com/certekim/xray-helper/app/middle"
	"github.com/certekim/xray-helper/app/xray"
)

type InboundObj map[string]interface{}
type InboundObjs []InboundObj

func ReadAll() []byte {
	var InboundObj InboundObj
	var InboundObjs InboundObjs
	records, err := middle.JsonDB.ReadAll("Inbound")
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

func Write(tag string, in map[string]interface{}) {
	if err := middle.JsonDB.Write("Inbound", tag, in); err != nil {
		log.Println("Error", err)
	}
}

func Read(tag string) []byte {
	var InboundObj InboundObj
	if err := middle.JsonDB.Read("Inbound", tag, &InboundObj); err != nil {
		log.Println("Error", err)
	}
	if ret, err := json.Marshal(InboundObj); err != nil {
		log.Println("Error", err)
	} else {
		return ret
	}
	return nil
}

func Del(tag string) {
	if err := middle.JsonDB.Delete("Inbound", tag); err != nil {
		log.Println("Error", err)
	}
}

func Apply(tag string) {
	xray.AddInbound(Read(tag))
}
