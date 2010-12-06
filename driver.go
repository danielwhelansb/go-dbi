package dbi

import (
    "os"
    "http"
    "strings"
)

type Driver interface {
    GetConnection(host, username, password, db string, options map[string][]string) (Connection, os.Error)
}

var drivers map[string]Driver

func AddDriver(name string, driver Driver) {
    if drivers == nil {
        drivers = make(map[string]Driver)
    }
    drivers[name] = driver
}

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

func init() {
    // drivers = make(map[string]Driver)
}

