package dbi

import (
    "os"
    "testing"
)

type DummyDriver struct {
}

func (self *DummyDriver) GetConnection(host, username, password, db string, options map[string]interface{}) (Connection, os.Error) {
    return new(DummyConnection), nil
}

func TestGetConnectionMethodWorks(*testing.T) {
    drv := new(DummyDriver)
    func(d Driver) {
        _, _ = d.GetConnection("localhost", "root", "", "testdb", nil)
    }(drv)
}

