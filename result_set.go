package dbi

import "os"

// API to be implemented for query resultsets.
type ResultSet interface {
    // The number of rows returned by the query.
    RowCount() uint64

    // Get the next row of results.
    Next() bool

    // Scan the current row of results by column index.
    Scan(refs ...interface{}) os.Error

    // Scan the current row of results by column name.
    NamedScan(refs ...interface{}) os.Error

    // Release the resources held by a ResultSet.
    Close() os.Error
}

