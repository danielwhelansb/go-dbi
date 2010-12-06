package dbi

import "os"

type ResultSet interface {
    Len() uint64
    Next() bool
    GetStringByCol(col int) (string, os.Error)
    GetString(col string) (string, os.Error)
    GetValueByCol(col int) (interface{}, os.Error)
    GetValue(col string) (interface{}, os.Error)
    Close() os.Error
}

