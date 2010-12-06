package dbi

import (
    "os"
    "testing"
)

type DummyConnection struct {
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

