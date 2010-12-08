package dbi

import (
    "os"
    "testing"
)

type DummyResultSet struct {
}

func (self *DummyResultSet) Len() uint64 {
    return 0
}

func (self *DummyResultSet) Next() bool {
    return false
}

func (self *DummyResultSet) GetString(col interface{}) (string, os.Error) {
    return "hello", nil
}

func (self *DummyResultSet) GetInt32(col interface{}) (int32, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) GetInt64(col interface{}) (int64, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) GetValue(col interface{}) (interface{}, os.Error) {
    return 123, nil
}

func (self *DummyResultSet) Close() os.Error {
    return nil
}

// ---------------------

func TestResultSetLenMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Len()
    }(new(DummyResultSet))
}

func TestResultSetNextMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Next()
    }(new(DummyResultSet))
}

func TestResultSetGetStringMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetString(1)
        rs.GetString("foo")
    }(new(DummyResultSet))
}

func TestResultSetGetInt32MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetInt32(1)
        rs.GetInt32("bar")
    }(new(DummyResultSet));
}

func TestResultSetGetInt64MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetInt64(1)
        rs.GetInt64("baz")
    }(new(DummyResultSet))
}

func TestResultSetGetValueMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetValue("test")
        rs.GetValue(0)
    }(new(DummyResultSet))
}

func TestResultSetCloseMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Close()
    }(new(DummyResultSet))
}

