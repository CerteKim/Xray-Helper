#!/bin/bash

hwi() {
    curl \
    --header "Content-Type: application/json" \
    --data '
    {
        "tag": "http-test",
        "port": 2080,
        "protocol": "http",
        "settings": {}
    }
    ' \
    http://127.0.0.1:8080/api/helper/inbound/add
}

hri() {
    curl \
    http://127.0.0.1:8080/api/helper/inbound/read/http-test
}

hdi() {
    curl \
    http://127.0.0.1:8080/api/helper/inbound/delete/http-test
}

hai() {
    curl \
    http://127.0.0.1:8080/api/helper/inbound/apply/http-test
}


case "$1" in
    hwi)
        hwi
        ;;
    hri)
        hri
        ;;
    hdi)
        hdi
        ;;
    hai)
        hai
        ;;
    *)
        echo "$0:  usage:  $0 {hwi|hri|hdi|hai}"
        ;;
esac