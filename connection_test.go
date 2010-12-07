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

func (self *DummyConnection) Query(sql string, params ...interface{}) (ResultSet, os.Error) {
    return new(DummyResultSet), nil
}

func (self *DummyConnection) Execute(sql string, params ...interface{}) os.Error {
    return nil
}

func (self *DummyConnection) GetOne(sql string, params ...interface{}) (interface{}, os.Error) {
    return 42, nil
}

func (self *DummyConnection) GetRow(sql string, params ...interface{}) (map[string]interface{}, os.Error) {
    return map[string]interface{}{"bar": "baz"}, nil
}

func (self *DummyConnection) GetAll(sql string, params ...interface{}) ([]map[string]interface{}, os.Error) {
    return [](map[string]interface{}){
        map[string]interface{}{"value": 1},
        map[string]interface{}{"value": 2},
    }, nil
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

func TestGetRowMethodWorks(t *testing.T) {
    conn := new(DummyConnection)
    func(c Connection) {
        res, err := c.GetRow("SELECT * FROM bar")
        if err != nil {
            t.Fatal("Expected GetRow() to pass")
        }
        value, found := res["bar"]
        if !found {
            t.Fatal("Expected row to contain key: 'bar'")
        }
        if value.(string) != "baz" {
            t.Fatal("Expected value == 'baz'")
        }
    }(conn)
}

func TestGetAllMethodWorks(t *testing.T) {
    conn := new(DummyConnection)
    func(c Connection) {
        res, err := c.GetAll("SELECT * FROM bar")
        if err != nil {
            t.Fatal("Expected GetAll() to pass")
        }

        value1, found := res[0]["value"]
        if !found {
            t.Fatal("Expected to find value field")
        }
        if value1.(int) != 1 {
            t.Fatal("Expected first row value = 1")
        }

        value2, found := res[1]["value"]
        if !found {
            t.Fatal("Expected to find value field")
        }
        if value2.(int) != 2 {
            t.Fatal("Expected second row value = 2")
        }
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

