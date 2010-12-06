package dbi

import "os"

type Connection interface {
    GetOne(sql string, params ...interface{}) (interface{}, os.Error)
    Close() os.Error
}

