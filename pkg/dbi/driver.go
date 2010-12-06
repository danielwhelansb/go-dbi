package dbi

import "os"

type Driver interface {
    GetConnection(host, username, password, db string, options map[string]interface{}) (Connection, os.Error)
}

