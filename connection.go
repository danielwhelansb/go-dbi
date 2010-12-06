package dbi

import "os"

type Connection interface {
    Query(sql string, params ...interface{}) (ResultSet, os.Error)
    Execute(sql string, params ...interface{}) os.Error
    GetOne(sql string, params ...interface{}) (interface{}, os.Error)
    Close() os.Error
}

