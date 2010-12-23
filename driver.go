package dbi

import (
    "os"
    "http"
)

// API to be implemented by all drivers (i.e. go-dbd-*)
type Driver interface {
    // Get a connection to the database identified by the given values.
    Connect(url *http.URL) (Connection, os.Error)
}

// Maps driver names to their implementation.
var drivers map[string]Driver

// Register a driver implementation.
func AddDriver(name string, driver Driver) {
    checkDriversReady()
    drivers[name] = driver
}

// Connect to the database specified by dsn.
func Connect(dsn string) (Connection, os.Error) {
    url, err := http.ParseURLReference(dsn)
    if err != nil {
        return nil, os.NewError("Invalid DSN URL: " + dsn)
    }
    driver, found := drivers[url.Scheme]
    if !found {
        return nil, os.NewError("No driver found: " + url.Scheme)
    }
    return driver.Connect(url)
}

func checkDriversReady() {
    if drivers == nil {
        drivers = make(map[string]Driver)
    }
}

func init() {
    //
    // Previous calls to AddDriver() by imported drivers may have
    // initialized drivers already.
    //
    checkDriversReady()
}

