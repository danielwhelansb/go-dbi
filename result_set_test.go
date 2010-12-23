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

func (self *DummyResultSet) Scan(refs ...interface{}) os.Error {
    return nil
}

func (self *DummyResultSet) String(col interface{}) (string, os.Error) {
    return "hello", nil
}

func (self *DummyResultSet) Int8(col interface{}) (int8, os.Error) {
    return 127, nil
}

func (self *DummyResultSet) Int16(col interface{}) (int16, os.Error) {
    return 12345, nil
}

func (self *DummyResultSet) Int32(col interface{}) (int32, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) Int64(col interface{}) (int64, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) Uint8(col interface{}) (uint8, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) Uint16(col interface{}) (uint16, os.Error) {
    return 65535, nil
}

func (self *DummyResultSet) Uint32(col interface{}) (uint32, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) Uint64(col interface{}) (uint64, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) Float32(col interface{}) (float32, os.Error) {
    return 0, nil
}

func (self *DummyResultSet) Float64(col interface{}) (float64, os.Error) {
    return 0, nil
}

func (self *DummyResultSet) Value(col interface{}) (interface{}, os.Error) {
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

func TestResultSetScanMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        var foo string
        rs.Scan(&foo)
        rs.Scan("foo", &foo)
    }(new(DummyResultSet))
}

func TestResultSetGetStringMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.String(1)
        rs.String("foo")
    }(new(DummyResultSet))
}

func TestResultSetGetInt8MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Int8(0)
        rs.Int8("blort")
    }(new(DummyResultSet))
}

func TestResultSetGetInt16MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Int16(0)
        rs.Int16("snert")
    }(new(DummyResultSet))
}

func TestResultSetGetInt32MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Int32(1)
        rs.Int32("bar")
    }(new(DummyResultSet));
}

func TestResultSetGetInt64MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Int64(1)
        rs.Int64("baz")
    }(new(DummyResultSet))
}

func TestResultSetGetUint8MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Uint8(0)
        rs.Uint8("bleh")
    }(new(DummyResultSet))
}

func TestResultSetGetUint16MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Uint16(1)
        rs.Uint16("bort")
    }(new(DummyResultSet))
}

func TestResultSetGetUint32MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Uint32(1)
        rs.Uint32("boz")
    }(new(DummyResultSet))
}

func TestResultSetGetUint64MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Uint64(1)
        rs.Uint64("bleh")
    }(new(DummyResultSet))
}

func TestResultSetGetFloat32MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Float32(8)
        rs.Float32("moo")
    }(new(DummyResultSet))
}

func TestResultSetGetFloat64MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Float64(2)
        rs.Float64("bah")
    }(new(DummyResultSet))
}

func TestResultSetGetValueMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Value("test")
        rs.Value(0)
    }(new(DummyResultSet))
}

func TestResultSetCloseMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Close()
    }(new(DummyResultSet))
}

