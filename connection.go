package dbi

import "os"

// Operations that can be performed on a database connection.
type Connection interface {
    // Executes a query, returning a ResultSet.
    Query(sql string, params ...interface{}) (ResultSet, os.Error)

    // Executes a query, discarding the ResultSet.
    Execute(sql string, params ...interface{}) os.Error

    // Get the value of the first field of the first row returned by a query.
    GetOne(sql string, params ...interface{}) (interface{}, os.Error)

    // Disconnect from the database.
    Close() os.Error
}

