package dbs

import (
	"context"
)

// General
var ctx = context.TODO()

var db_name = "fastnetmon"
var db_url  = "mongodb://127.0.0.1:27017"

// Connect to db
var DB = Use(db_name)

