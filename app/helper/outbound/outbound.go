package outbound

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/certekim/xray-helper/app/middle/outbound"
	"github.com/certekim/xray-helper/app/utils"
	"github.com/certekim/xray-helper/app/xray"
	"github.com/julienschmidt/httprouter"
)

func WriteOutboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res, err, params := utils.Convert(r)
	if err != nil {
		log.Println(err)
		return
	}
	var ret map[string]interface{}
	err = json.Unmarshal(res, &ret)
	if err != nil {
		log.Println(err)
		return
	}
	if tag, err := params["tag"].(string); err {
		outbound.Write(tag, ret)
		log.Println(err)
		return
	} else {
		log.Println(err)
	}
}

func ReadOutboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tag := ps.ByName("tag")
	var data map[string]interface{}
	err := json.Unmarshal(outbound.Read(tag), &data)
	if err != nil {
		log.Println(err)
		return
	}
	if res, err := json.Marshal(data); err != nil {
		log.Println(err)
		return
	} else {
		fmt.Fprintf(w, "%s", res)
	}
}

func DeleteOutboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tag := ps.ByName("tag")
	outbound.Del(tag)
}

func ApplyOutboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tag := ps.ByName("tag")
	var data map[string]interface{}
	err := json.Unmarshal(outbound.Read(tag), &data)
	if err != nil {
		log.Println(err)
		return
	}
	res, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	xray.AddOutbound(res)
}
