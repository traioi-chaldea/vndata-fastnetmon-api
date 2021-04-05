package dbs

import (
	"context"
)

// General
var ctx = context.TODO()

var db_name = "fastnetmon"
var db_url  = "mongodb://172.21.1.13:27017"

// Connect to db
var DB = Use(db_name)

// List special Collections
var teleCollection   = DB.Collection("TELE_INFO")
var apiCollection    = DB.Collection("API_INFO")
var switchCollection = DB.Collection("SW_INFO")
var vlanCollection   = DB.Collection("VLAN_INFO")
var routeCollection  = DB.Collection("ROUTE_INFO")
var aCollection      = DB.Collection("ACL")
var bCollection      = DB.Collection("ROUTE")

