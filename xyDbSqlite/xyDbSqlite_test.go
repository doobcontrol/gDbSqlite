package xyDbSqlite

import (
	"fmt"
	"github.com/doobcontrol/gDb/xyDb"
	"os"
	"testing"
)

func TestDbSqliteAccess_SetConnect(t *testing.T) {
	mysqlAccess := &DbSqliteAccess{}

	// Ensure the connection string is valid and can connect to the test database
    err := mysqlAccess.SetConnect(testDFile)
    if err != nil {
        t.Error(fmt.Printf("DbAccess.SetConnect err expected nil, but got an error: %s", err.Error()))
    }
}
func TestDbSqliteAccess_InitDb(t *testing.T) {
	pars := &map[string]string{
		S_dbFile: testDFile,
	}
	db := xyDb.DbStructure{
		DbName: "testDb",
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

	mysqlAccess := &DbSqliteAccess{}

	//clean db
	cleanDb()

    newConnectString, err := mysqlAccess.InitDb(pars, db)
    if err != nil {
        t.Error(fmt.Printf("DbAccess.InitDb err expected nil, but got an error: %s", err.Error()))
    }
    if newConnectString != testDFile {
        t.Error(fmt.Printf("DbAccess.InitDb newConnectString expected: %s, but got: %s", 
		testDFile, newConnectString))
    }

	//clean db
	cleanDb()
}
func TestDbSqliteAccess_sql(t *testing.T) {
	mysqlAccess := &DbSqliteAccess{}

	//clean db
	cleanDb()

    //init db
	initDb(mysqlAccess)
	mysqlAccess.Close()

	//connect
	mysqlAccess.SetConnect(testDFile)

	err := mysqlAccess.ExSql(fmt.Sprintf("insert into %s(%s) values(%s)","table1","F1,F2","'abc','efg'"))
    if err != nil {
        t.Error(fmt.Printf("DbAccess.ExSql.insert err expected nil, but got an error: %s", err.Error()))
    } else {
		t.Log("DbAccess.ExSql.insert succeed")
	}
	err = mysqlAccess.ExSql(fmt.Sprintf("insert into %s(%s) values(%s)","table1","F1,F2","'abc1','efg2'"))
    if err != nil {
        t.Error(fmt.Printf("DbAccess.ExSql.insert err expected nil, but got an error: %s", err.Error()))
	} else {
		t.Log("DbAccess.ExSql.insert succeed")
	}

	err = mysqlAccess.ExSql(fmt.Sprintf("update %s set %s where %s='%s'","table1","F2='zzz'","F1", "abc"))
    if err != nil {
        t.Error(fmt.Printf("DbAccess.ExSql.update err expected nil, but got an error: %s", err.Error()))
    } else {
		t.Log("DbAccess.ExSql.update succeed")
	}

	record, err := mysqlAccess.Query(fmt.Sprintf("select * from %s","table1"))
    if err != nil {
        t.Error(fmt.Printf("DbAccess.ExSql.select err expected nil, but got an error: %s", err.Error()))
    } else {
		if len(*record) != 2 {
			t.Error(fmt.Printf("DbAccess.ExSql.select expected: %d records, but got: %d records", 2, len(*record)))
		} else {
			t.Log("DbAccess.ExSql.select succeed")
		}
	}

	err = mysqlAccess.ExSql(fmt.Sprintf("delete from %s where %s='%s'","table1","F1", "abc"))
    if err != nil {
        t.Error(fmt.Printf("DbAccess.ExSql.delete err expected nil, but got an error: %s", err.Error()))
    } else {
		t.Log("DbAccess.ExSql.delete succeed")
	}

	//clean db
	cleanDb()
}

var testDFile = "./testDb"
func cleanDb(){
	os.Remove(testDFile)
}
func initDb(mysqlAccess *DbSqliteAccess){
	pars := &map[string]string{
		S_dbFile: testDFile,
	}
	db := xyDb.DbStructure{
		DbName: "testDb",
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
	mysqlAccess.InitDb(pars, db)
}
