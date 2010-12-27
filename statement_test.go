package dbi

import (
    "os"
    "testing"
)

type DummyStatement struct {
}

func (stmt *DummyStatement) Query(params ...interface{}) (ResultSet, os.Error) {
    return nil, nil
}

func (stmt *DummyStatement) Execute(params ...interface{}) os.Error {
    return nil
}

func (stmt *DummyStatement) Close() os.Error {
    return nil
}

func TestStatementQueryMethodWorks(t *testing.T) {
    func(s Statement) {
        _, _ = s.Query(1, "foo", nil)
    }(new(DummyStatement))
}

func TestStatementExecuteMethodWorks(t *testing.T) {
    func(s Statement) {
        _ = s.Execute(3, "bar", nil)
    }(new(DummyStatement))
}

func TestStatementCloseMethodWorks(t *testing.T) {
    func(s Statement) {
        _ = s.Close()
    }(new(DummyStatement))
}

