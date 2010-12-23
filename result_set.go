package dbi

import "os"

// API to be implemented for query resultsets.
type ResultSet interface {
    // The number of rows returned by the query.
    RowCount() uint64

    // Get the next row of results.
    Next() bool

    // Scan the current row of results.
    Scan(refs ...interface{}) os.Error

    // Get a string value by column name.
    String(name string) (string, os.Error)

    // Get an int8 value by column name.
    Int8(col string) (int8, os.Error)

    // Get an int16 value by column name.
    Int16(col string) (int16, os.Error)

    // Get an int32 value by column name.
    Int32(col string) (int32, os.Error)

    // Get an int64 value by column name.
    Int64(col string) (int64, os.Error)

    // Get a uint8/byte value by column name.
    Uint8(col string) (uint8, os.Error)

    // Get a uint16 value by column name.
    Uint16(col string) (uint16, os.Error)

    // Get a uint32 value by column name.
    Uint32(col string) (uint32, os.Error)

    // Get a uint64 value by column name.
    Uint64(col string) (uint64, os.Error)

    // Get a float32 value by column name.
    Float32(col string) (float32, os.Error)

    // Get a float64 value by column name.
    Float64(col string) (float64, os.Error)

    // Get a value of arbitrary type by column name.
    Value(col string) (interface{}, os.Error)

    // Release the resources held by a ResultSet.
    Close() os.Error
}

