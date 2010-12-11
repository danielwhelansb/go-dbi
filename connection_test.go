package dbi

import (
    "os"
    "testing"
)

type DummyConnection struct {
    host string
    userinfo string
    db string
    options map[string][]string
}

func (self *DummyConnection) Query(sql string, params ...interface{}) (ResultSet, os.Error) {
    return new(DummyResultSet), nil
}

func (self *DummyConnection) Execute(sql string, params ...interface{}) os.Error {
    return nil
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

func TestExecuteMethodWorks(t *testing.T) {
    conn := new(DummyConnection)
    func(c Connection) {
        err := c.Execute("CREATE TABLE users (...)")
        if err != nil {
            t.Fatal("Expected Execute() to pass")
        }
    }(conn)
}

func TestQueryMethodWorks(t *testing.T) {
    conn := new(DummyConnection)
    func (c Connection) {
        _, err := c.Query("SELECT * FROM users")
        if err != nil {
            t.Fatal("Expected Query() to pass")
        }
    }(conn)
}

