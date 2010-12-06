package dbi

import (
    "os"
    "http"
    "strings"
)

type Driver interface {
    GetConnection(host, username, password, db string, options map[string]interface{}) (Connection, os.Error)
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

    return driver.GetConnection(url.Host, username, password, db, nil)
}

func init() {
    // drivers = make(map[string]Driver)
}

