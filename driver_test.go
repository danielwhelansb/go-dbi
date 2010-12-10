package dbi

import (
    "os"
    "testing"
    "http"
)

type DummyDriver struct {
}

func (self *DummyDriver) GetConnection(url *http.URL) (Connection, os.Error) {
    conn := new(DummyConnection)
    conn.host = url.Host
    conn.userinfo = url.RawUserinfo
    conn.db = url.Path[1:]
    options, err := http.ParseQuery(url.RawQuery)
    if err != nil {
        return nil, err
    }
    conn.options = options
    return conn, nil
}

func TestGetConnectionMethodWorks(*testing.T) {
    drv := new(DummyDriver)
    func(d Driver) {
        url, _ := http.ParseURL("foo://root@localhost/testdb")
        _, _ = d.GetConnection(url)
    }(drv)
}

func TestAddDriver(t *testing.T) {
    drv := new(DummyDriver)
    AddDriver("dummy", drv)
    _, err := Connect("dummy://localhost/foo")
    if err != nil {
        t.Fatal("Expected Connect() to succeed: " + err.String())
    }
}

func TestConnectWithoutAuth(t *testing.T) {
    drv := new(DummyDriver)
    AddDriver("dummy", drv)

    conn, err := Connect("dummy://localhost/foo")
    if err != nil {
        t.Fatal("Expected Connect() to succeed: " + err.String())
    }
    
    dummyConn := conn.(*DummyConnection)
    if dummyConn.host != "localhost" {
        t.Fatal("dummyConn.host != 'localhost': " + err.String())
    }
    if dummyConn.db != "foo" {
        t.Fatal("dummyConn.db != 'foo': " + err.String())
    }
}

func TestConnectWithAuth(t *testing.T) {
    drv := new(DummyDriver)
    AddDriver("dummy", drv)

    conn, err := Connect("dummy://root:muffin@shinetech.com/bar")
    if err != nil {
        t.Fatal("Expected Connect() to succeed: " + err.String())
    }

    dummyConn := conn.(*DummyConnection)
    if dummyConn.host != "shinetech.com" {
        t.Fatal("dummyConn.host != 'shinetech.com': " + dummyConn.host)
    }
    if dummyConn.userinfo != "root:muffin" {
        t.Fatal("dummyConn.userinfo != 'root:muffin': " + dummyConn.userinfo)
    }
    if dummyConn.db != "bar" {
        t.Fatal("dummyConn.db != 'bar': " + dummyConn.db)
    }
}

func TestConnectWithQueryString(t *testing.T) {
    drv := new(DummyDriver)
    AddDriver("dummy", drv)

    conn, err := Connect("dummy://foo/bar?baz=true")
    if err != nil {
        t.Fatal("Expected Connect() to succeed: " + err.String())
    }

    dummyConn := conn.(*DummyConnection)
    value, found := dummyConn.options["baz"]
    if !found {
        t.Fatal("Could not find 'baz' parameter in dummyConn options")
    }
    if len(value) != 1 {
        t.Fatal("Expected exactly one value for the baz option.")
    }
    if value[0] != "true" {
        t.Fatal("Expected baz option to be the string 'true'")
    }
}

