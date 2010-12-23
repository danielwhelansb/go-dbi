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

func (self *DummyResultSet) String(col string) (string, os.Error) {
    return "hello", nil
}

func (self *DummyResultSet) Int8(col string) (int8, os.Error) {
    return 127, nil
}

func (self *DummyResultSet) Int16(col string) (int16, os.Error) {
    return 12345, nil
}

func (self *DummyResultSet) Int32(col string) (int32, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) Int64(col string) (int64, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) Uint8(col string) (uint8, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) Uint16(col string) (uint16, os.Error) {
    return 65535, nil
}

func (self *DummyResultSet) Uint32(col string) (uint32, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) Uint64(col string) (uint64, os.Error) {
    return 1, nil
}

func (self *DummyResultSet) Float32(col string) (float32, os.Error) {
    return 0, nil
}

func (self *DummyResultSet) Float64(col string) (float64, os.Error) {
    return 0, nil
}

func (self *DummyResultSet) Value(col string) (interface{}, os.Error) {
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
        rs.String("foo")
    }(new(DummyResultSet))
}

func TestResultSetGetInt8MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Int8("blort")
    }(new(DummyResultSet))
}

func TestResultSetGetInt16MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Int16("snert")
    }(new(DummyResultSet))
}

func TestResultSetGetInt32MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Int32("bar")
    }(new(DummyResultSet));
}

func TestResultSetGetInt64MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Int64("baz")
    }(new(DummyResultSet))
}

func TestResultSetGetUint8MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Uint8("bleh")
    }(new(DummyResultSet))
}

func TestResultSetGetUint16MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Uint16("bort")
    }(new(DummyResultSet))
}

func TestResultSetGetUint32MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Uint32("boz")
    }(new(DummyResultSet))
}

func TestResultSetGetUint64MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Uint64("bleh")
    }(new(DummyResultSet))
}

func TestResultSetGetFloat32MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Float32("moo")
    }(new(DummyResultSet))
}

func TestResultSetGetFloat64MethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Float64("bah")
    }(new(DummyResultSet))
}

func TestResultSetGetValueMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Value("test")
    }(new(DummyResultSet))
}

func TestResultSetCloseMethodWorks(t *testing.T) {
    func(rs ResultSet) {
        rs.Close()
    }(new(DummyResultSet))
}

