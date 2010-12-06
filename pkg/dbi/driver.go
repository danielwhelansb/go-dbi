package dbi

import "os"

type Driver interface {
    GetConnection(host, username, password string, options map[string]interface{}) (Connection, os.Error)
}

