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

func (self *DummyResultSet) FetchArray() []interface{} {
    return nil
}

func (self *DummyResultSet) FetchMap() map[string]interface{} {
    return nil
}

func (self *DummyResultSet) Close() os.Error {
    return nil
}

func TestResultSetNextMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Next()
    }(new(DummyResultSet))
}

func TestResultSetFetchArrayMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.FetchArray()
    }(new(DummyResultSet))
}

func TestResultSetFetchMapMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.FetchMap()
    }(new(DummyResultSet))
}

func TestResultSetCloseMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Close()
    }(new(DummyResultSet))
}

