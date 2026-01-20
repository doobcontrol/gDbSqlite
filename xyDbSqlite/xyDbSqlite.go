package xyDbSqlite

import (
	"database/sql"
	"github.com/doobcontrol/gDb/xyDb"
	_ "modernc.org/sqlite"
)

 // Define DbSqliteAccess
 const S_dbFile = "dbFile"
 const S_sqliteDriverName = "sqlite"
 type DbSqliteAccess struct{
    xyDb.DbAccess
 }
 func (dba *DbSqliteAccess) InitDb(initPars map[string]string, dbStructure xyDb.DbStructure) (string, error) {
	dba.SetDriverName(S_sqliteDriverName)
	db, err := sql.Open(dba.DbDriverName, initPars[S_dbFile])
	if err != nil {
        return "", err
	}
	defer db.Close() // Defer closing the database connection until the main function finishes

    dScript, err := dba.DbScript(dbStructure)
	if err != nil {
        return "", err
	}
	if _, err = db.Exec(dScript); err != nil {
        return "", err
	}

	db, err = sql.Open(dba.DbDriverName, initPars[S_dbFile])
	if err != nil {
        return "", err
	}

	dba.Db = db
    return initPars[S_dbFile], nil
 }