package outbound

import (
	"encoding/json"
	"log"

	"github.com/certekim/xray-helper/app/middle"
	"github.com/certekim/xray-helper/app/xray"
)

type OutboundObj map[string]interface{}

type OutboundObjs []OutboundObj

func ReadAll() []byte {
	var OutboundObj OutboundObj
	var OutboundObjs OutboundObjs
	records, err := middle.JsonDB.ReadAll("Outbound")
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

func Write(tag string, in map[string]interface{}) {
	if err := middle.JsonDB.Write("Outbound", tag, in); err != nil {
		log.Println("Error", err)
	}
}

func Read(tag string) []byte {
	var OutboundObj OutboundObj
	if err := middle.JsonDB.Read("Outbound", tag, &OutboundObj); err != nil {
		log.Println("Error", err)
	}
	if ret, err := json.Marshal(OutboundObj); err != nil {
		log.Println("Error", err)
	} else {
		return ret
	}
	return nil
}

func Del(tag string) {
	if err := middle.JsonDB.Delete("Outbound", tag); err != nil {
		log.Println("Error", err)
	}
}

func Apply(tag string) {
	xray.AddOutbound(Read(tag))
}
