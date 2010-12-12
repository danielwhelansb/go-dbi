package dbi

import "os"

// Operations that can be performed on a database connection.
type Connection interface {
    // Executes a query, returning a ResultSet.
    Query(sql string, params ...interface{}) (ResultSet, os.Error)

    // Executes a query, discarding the ResultSet.
    Execute(sql string, params ...interface{}) os.Error

    // Begins a transaction.
    BeginTransaction() os.Error

    // Rolls back the current transaction.
    Rollback() os.Error

    // Commits the current transaction.
    Commit() os.Error

    // Disconnect from the database.
    Close() os.Error
}

