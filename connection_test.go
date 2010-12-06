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
    options map[string][]string
}

func (self *DummyConnection) GetOne(sql string, params ...interface{}) (interface{}, os.Error) {
    return 42, nil
}

func (self *DummyConnection) Close() os.Error {
    return nil
}

//
// Below we test everything via the Connection interface to ensure that
// the method really needs to be provided by DummyConnection.
//

func TestCloseMethodWorks(*testing.T) {
    conn := new(DummyConnection)
    func(c Connection) {
        c.Close()
    }(conn)
}

func TestGetOneMethodWorks(t *testing.T) {
    conn := new(DummyConnection)
    func(c Connection) {
        res, err := c.GetOne("SELECT * FROM foo")
        if err != nil {
            t.Fatal("Expected GetOne() to pass")
        }
        if res.(int) != 42 {
            t.Fatal("Expected result == 42")
        }
    }(conn)
}

