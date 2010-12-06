package dbi

import (
    "os"
    "testing"
)

type DummyConnection struct {
    host string
    username string
    password string
    db string
    options map[string]interface{}
}

func (self *DummyConnection) Close() os.Error {
    return nil
}

func TestCloseMethodWorks(*testing.T) {
    conn := new(DummyConnection)
    func(c Connection) {
        c.Close()
    }(conn)
}

