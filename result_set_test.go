package dbi

import (
    "os"
    "testing"
)

type DummyResultSet struct {
}

func (self *DummyResultSet) RowCount() uint64 {
    return 0
}

func (self *DummyResultSet) Next() bool {
    return false
}

func (self *DummyResultSet) GetString(col interface{}) (string, os.Error) {
    return "hello", nil
}

func (self *DummyResultSet) GetInt8(col interface{}) (int8, os.Error) {
    return 127, nil
}

func (self *DummyResultSet) GetInt16(col interface{}) (int16, os.Error) {
    return 12345, nil
}

func (self *DummyResultSet) GetInt32(col interface{}) (int32, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) GetInt64(col interface{}) (int64, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) GetUint8(col interface{}) (uint8, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) GetUint16(col interface{}) (uint16, os.Error) {
    return 65535, nil
}

func (self *DummyResultSet) GetUint32(col interface{}) (uint32, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) GetUint64(col interface{}) (uint64, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) GetFloat32(col interface{}) (float32, os.Error) {
    return 0, nil
}

func (self *DummyResultSet) GetFloat64(col interface{}) (float64, os.Error) {
    return 0, nil
}

func (self *DummyResultSet) GetValue(col interface{}) (interface{}, os.Error) {
    return 123, nil
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

func TestResultSetGetStringMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetString(1)
        rs.GetString("foo")
    }(new(DummyResultSet))
}

func TestResultSetGetInt8MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetInt8(0)
        rs.GetInt8("blort")
    }(new(DummyResultSet))
}

func TestResultSetGetInt16MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetInt16(0)
        rs.GetInt16("snert")
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

func TestResultSetGetUint8MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetUint8(0)
        rs.GetUint8("bleh")
    }(new(DummyResultSet))
}

func TestResultSetGetUint16MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetUint16(1)
        rs.GetUint16("bort")
    }(new(DummyResultSet))
}

func TestResultSetGetUint32MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetUint32(1)
        rs.GetUint32("boz")
    }(new(DummyResultSet))
}

func TestResultSetGetUint64MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetUint64(1)
        rs.GetUint64("bleh")
    }(new(DummyResultSet))
}

func TestResultSetGetFloat32MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetFloat32(8)
        rs.GetFloat32("moo")
    }(new(DummyResultSet))
}

func TestResultSetGetFloat64MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.GetFloat64(2)
        rs.GetFloat64("bah")
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

