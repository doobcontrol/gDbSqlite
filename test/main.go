package main

import (
	"fmt"
	"github.com/doobcontrol/gDb/xyDb"
	"github.com/doobcontrol/gDbSqlite/xyDbSqlite"
)

func main() {
	xyDb.DService.DbAccess = &xyDbSqlite.DbSqliteAccess{}
	dbStru := xyDb.DbStructure{
		DbName: "test",
		Tables: []xyDb.DbTable{
			{
				TableName: "table1",
				Fields: []xyDb.DbField{
					{
						FieldName: "F1",
						DataType: "string",
						Length: 10,
						IsKey: true,
					},
					{
						FieldName: "F2",
						DataType: "string",
						Length: 10,
						IsKey: false,
					},
				},
			},
			{
				TableName: "table2",
				Fields: []xyDb.DbField{
					{
						FieldName: "F1",
						DataType: "string",
						Length: 10,
						IsKey: true,
					},
					{
						FieldName: "F2",
						DataType: "string",
						Length: 10,
						IsKey: false,
					},
				},
			},
		},
	}
	initPars := map[string]string{xyDbSqlite.S_dbFile:"./test.db"}
	connectString, err := xyDb.DService.Init(&initPars, dbStru)
	if err != nil {
		fmt.Println("xyDb Library test error: ", err)
	} else{
		fmt.Println("xyDb Library test done,connectString: ", connectString)
	}
}