package dbs

import (
	"github.com/TraiOi/util"
)

func unbanIPfromACL(content ACL) bool {
	var result bool
	col := DB.Collection("ACL")
	if _, err := col.DeleteOne(ctx, content); err != nil {
		util.ErrorLogger.Print(err)
		result = false
	} else {
		util.InfoLogger.Print("ACL: Unbanned success")
		result = true
	}
	return result
}

func unbanIPfromROUTE(content ACL) bool {
	var result bool
	col := DB.Collection("ROUTE")
	if _, err := col.DeleteOne(ctx, content); err != nil {
		util.ErrorLogger.Print(err)
		result = false
	} else {
		util.InfoLogger.Print("ROUTE: Unbanned success")
		result = true
	}
	return result
}

func UnbanIP(acl ACL, route ACL) bool {
	return unbanIPfromACL(acl) && unbanIPfromROUTE(route)
}
