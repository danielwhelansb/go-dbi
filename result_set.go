package dbi

import "os"

// API to be implemented for query resultsets.
type ResultSet interface {
    // The number of rows returned by the query.
    RowCount() uint64

    // Get the next row of results.
    Next() bool

    // Get a string value by column name or zero-based index.
    GetString(col interface{}) (string, os.Error)

    // Get an int8 value by column name or zero-based index.
    GetInt8(col interface{}) (int8, os.Error)

    // Get an int16 value by column name or zero-based index.
    GetInt16(col interface{}) (int16, os.Error)

    // Get an int32 value by column name or zero-based index.
    GetInt32(col interface{}) (int32, os.Error)

    // Get an int64 value by column name or zero-based index.
    GetInt64(col interface{}) (int64, os.Error)

    // Get a uint8/byte value by column name or zero-based index.
    GetUint8(col interface{}) (uint8, os.Error)

    // Get a uint16 value by column name or zero-based index.
    GetUint16(col interface{}) (uint16, os.Error)

    // Get a uint32 value by column name or zero-based index.
    GetUint32(col interface{}) (uint32, os.Error)

    // Get a uint64 value by column name or zero-based index.
    GetUint64(col interface{}) (uint64, os.Error)

    // Get a float32 value by column name or zero-based index.
    GetFloat32(col interface{}) (float32, os.Error)

    // Get a float64 value by column name or zero-based index.
    GetFloat64(col interface{}) (float64, os.Error)

    // Get a value of arbitrary type by column name or zero-based index.
    GetValue(col interface{}) (interface{}, os.Error)

    // Release the resources held by a ResultSet.
    Close() os.Error
}

