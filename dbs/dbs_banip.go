package dbs

import (
	"github.com/TraiOi/util"
)

func banIPtoACL(content ACL) bool {
	var result bool
	col := DB.Collection("ACL")
	if _, err := col.InsertOne(ctx, content); err != nil {
		util.ErrorLogger.Print(err)
		result = false
	} else {
		util.InfoLogger.Print("ACL: Banned success")
		result = true
	}
	return result
}

func banIPtoROUTE(content ACL) bool {
	var result bool
	col := DB.Collection("ROUTE")
	if _, err := col.InsertOne(ctx, content); err != nil {
		util.ErrorLogger.Print(err)
		result = false
	} else {
		util.InfoLogger.Print("ROUTE: Banned success")
		result = true
	}
	return result
}

func BanIP(acl ACL, route ACL) bool {
	return banIPtoACL(acl) && banIPtoROUTE(route)
}
