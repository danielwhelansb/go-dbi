package dbi

import (
    "os"
    "testing"
)

type DummyResultSet struct {
}

func (self *DummyResultSet) RowCount() (uint64, os.Error) {
    return 0, nil
}

func (self *DummyResultSet) Next() bool {
    return false
}

func (self *DummyResultSet) Scan(refs ...interface{}) os.Error {
    return nil
}

func (self *DummyResultSet) NamedScan(refs ...interface{}) os.Error {
    return nil
}

func (self *DummyResultSet) Close() os.Error {
    return nil
}

// ---------------------

func TestResultSetRowCountMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.RowCount()
    }(new(DummyResultSet))
}

func TestResultSetNextMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Next()
    }(new(DummyResultSet))
}

func TestResultSetScanMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        var foo string
        rs.Scan(&foo)
    }(new(DummyResultSet))
}

func TestResultSetNamedScanMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        var foo string
        rs.NamedScan("foo", &foo)
    }(new(DummyResultSet))
}

func TestResultSetCloseMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Close()
    }(new(DummyResultSet))
}

