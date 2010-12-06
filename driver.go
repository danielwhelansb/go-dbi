package dbi

import (
    "os"
    "http"
    "strings"
)

// API to be implemented by all drivers (i.e. go-dbd-*)
type Driver interface {
    // Get a connection to the database identified by the given values.
    GetConnection(host, username, password, db string, options map[string][]string) (Connection, os.Error)
}

// Maps driver names to their implementation.
var drivers map[string]Driver

// Register a driver implementation.
func AddDriver(name string, driver Driver) {
    if drivers == nil {
        drivers = make(map[string]Driver)
    }
    drivers[name] = driver
}

// Connect to the database specified by dsn.
func Connect(dsn string) (Connection, os.Error) {
    url, err := http.ParseURLReference(dsn)
    if err != nil {
        return nil, os.NewError("Invalid URL: " + dsn)
    }
    userpass := strings.Split(url.RawUserinfo, ":", 2)

    username := userpass[0]
    var password string
    if len(userpass) > 1 {
        password = userpass[1]
    } else {
        password = ""
    }

    db := url.Path[1:]

    driver, found := drivers[url.Scheme]
    if !found {
        return nil, os.NewError("No such driver: " + url.Scheme)
    }

    options, err := http.ParseQuery(url.RawQuery)
    if err != nil {
        return nil, os.NewError("Could not parse query string: " + url.RawQuery)
    }

    return driver.GetConnection(url.Host, username, password, db, options)
}

