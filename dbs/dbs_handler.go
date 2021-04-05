package dbs

import (
	"fmt"
	"github.com/TraiOi/util"
	"go.mongodb.org/mongo-driver/bson"
)

func getInfo(collection string) []map[string]interface{} {
	var info []map[string]interface{}

	cursor, err := DB.Collection(collection).Find(ctx, bson.M{})
	if err != nil {
		util.ErrorLogger.Print(err)
	}

	if err = cursor.All(ctx, &info); err != nil {
		util.ErrorLogger.Print(err)
	} else {
		util.InfoLogger.Print("Getting all '" + collection + "' ...")
	}

	return info
}

func FormatIP(ip string) string {
	return fmt.Sprintf("%s/32", ip)
}

func GetSWInfo() interface{} {
	return getInfo("SW_INFO")
}
