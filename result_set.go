package dbi

import "os"

// API to be implemented for query resultsets.
type ResultSet interface {
    // The number of rows returned by the query.
    Len() uint64

    // Get the next row of results.
    Next() bool

    // Get a string value by column index (zero-based).
    GetStringByCol(col int) (string, os.Error)

    // Get a string value by column name.
    GetString(col string) (string, os.Error)

    // Get an arbitrary value by column index (zero-based).
    GetValueByCol(col int) (interface{}, os.Error)

    // Get an arbitrary value by column name.
    GetValue(col string) (interface{}, os.Error)

    // Release the resources held by a ResultSet.
    Close() os.Error
}

