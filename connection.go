package dbi

import "os"

type Connection interface {
    Execute(sql string, params ...interface{}) os.Error
    GetOne(sql string, params ...interface{}) (interface{}, os.Error)
    Close() os.Error
}

