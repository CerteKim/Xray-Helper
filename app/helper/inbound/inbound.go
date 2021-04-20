package inbound

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/certekim/xray-helper/app/middle/inbound"
	"github.com/certekim/xray-helper/app/utils"
	"github.com/certekim/xray-helper/app/xray"
	"github.com/julienschmidt/httprouter"
)

func WriteInboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
		inbound.Write(tag, ret)
		log.Println(err)
		return
	} else {
		log.Println(err)
	}
}

func ReadInboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tag := ps.ByName("tag")
	var data map[string]interface{}
	err := json.Unmarshal(inbound.Read(tag), &data)
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

func DeleteInboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tag := ps.ByName("tag")
	inbound.Del(tag)
}

func ApplyInboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tag := ps.ByName("tag")
	var data map[string]interface{}
	err := json.Unmarshal(inbound.Read(tag), &data)
	if err != nil {
		log.Println(err)
		return
	}
	res, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	xray.AddInbound(res)
}
