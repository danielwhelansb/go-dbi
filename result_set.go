package dbi

import "os"

// API to be implemented for query resultsets.
type ResultSet interface {
    // The number of rows returned by the query.
    Len() uint64

    // Get the next row of results.
    Next() bool

    // Get a string value by column name or zero-based index.
    GetString(col interface{}) (string, os.Error)

    // Get an int32 value by column name or zero-based index.
    GetInt32(col interface{}) (int, os.Error)

    // Get a value of arbitrary type by column name or zero-based index.
    GetValue(col interface{}) (interface{}, os.Error)

    // Release the resources held by a ResultSet.
    Close() os.Error
}

