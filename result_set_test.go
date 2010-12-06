package dbi

import (
    "os"
    "testing"
)

type DummyResultSet struct {
}

func (self *DummyResultSet) Next() bool {
    return false
}

func (self *DummyResultSet) GetStringByCol(col int) (string, os.Error) {
    return "hello", nil
}

func (self *DummyResultSet) GetString(col string) (string, os.Error) {
    return "hello with string", nil
}

func (self *DummyResultSet) Close() os.Error {
    return nil
}

func TestResultSetNextMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Next()
    }(new(DummyResultSet))
}

func TestResultSetGetStringByColMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetStringByCol(1)
    }(new(DummyResultSet))
}

func TestResultSetGetStringMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetString("foo")
    }(new(DummyResultSet))
}

func TestResultSetCloseMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Close()
    }(new(DummyResultSet))
}

