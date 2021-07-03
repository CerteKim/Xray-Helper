package log

import (
	"log"
	"os"
)

var Stdlog *log.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
var Errlog *log.Logger = log.New(os.Stderr, "", log.Ldate|log.Ltime)
