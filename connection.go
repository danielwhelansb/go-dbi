package dbi

import "os"

type Connection interface {
    Close() os.Error
}

