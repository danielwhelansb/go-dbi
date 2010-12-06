package dbi

import "os"

type ResultSet interface {
    Next() bool
    FetchArray() []interface{}
    FetchMap() map[string]interface{}
    Close() os.Error
}

