# go-dbi

go-dbi is a database independent interface for accessing databases from the
Go language.

## Status

Alpha.

## Overview

Perl's DBI interface is a popular model for database connectivity in modern
programming languages. go-dbi brings a flavour of this convenient API to the
Go programming language.

If you're at all familiar with Perl DBI, the PEAR::DB or MDB2 PHP
libraries or Ruby/DBI, you should feel right at home with go-dbi.

## Installation

1. Make sure you have a working Go environment. See the
[install instructions](http://golang.org/doc/install.html).
2. Ensure goinstall is on your $PATH.
3. Install go-dbi: $ goinstall github.com/thomaslee/go-dbi

Alternatively, you can clone the git repository & build it using make.

## Usage

This example assumes you have both go-dbi and the
[go-dbd-mysql](http://github.com/thomaslee/go-dbd-mysql) driver installed.

    package main

    import (
        "fmt"
        dbi "github.com/thomaslee/go-dbi"
        //
        // important: import all drivers that need to be supported by your
        //            application!
        //
        _ "github.com/thomaslee/go-dbd-mysql"
    )

    func main() {
        conn, err := dbi.Connect("mysql://root@localhost/somedatabase")
        if err != nil {
            fmt.Printf("error: unable to connect to the database: %s\n", err.String())
            return
        }
        defer conn.Close()

        err = conn.Execute("INSERT INTO users (username) VALUES ('tom')")
        if err != nil {
            fmt.Printf("error: unable to insert user: %s\n", err.String())
            return
        }

        rs, err := conn.Query("SELECT username FROM users WHERE username='tom'")
        if err != nil {
            fmt.Printf("error: unable to fetch users: %s\n", err.String())
            return
        }
        defer rs.Close()

        for rs.Next() {
            // You can scan the current row of values like so ...
            var username string
            err = rs.Scan(&username)
            if err != nil {
                fmt.Printf("error: %s\n", err.String())
            } else {
                fmt.Printf("rs.Scan(&username): %s\n", username)
            }

            // ... or scan by column name ...
            err = rs.Scan("username", &username)
            if err != nil {
                fmt.Printf("error: %s\n", err.String())
            } else {
                fmt.Printf("rs.Scan(\"username\", &username): %s\n", username)
            }

            // ... or by using more traditional accessors.
            username, err = rs.String("username")
            if err != nil {
                fmt.Printf("error: %s\n", err.String())
            } else {
                fmt.Printf("GetString(\"username\"): %s\n", username)
            }
        }
    }


## License

This software is licensed under the terms of the [MIT License](http://en.wikipedia.org/wiki/MIT_License).

## Support

Please log defects and feature requests using the issue tracker on github.

## About

go-dbi was written by [Tom Lee](http://tomlee.co).

Follow me on [Twitter](http://www.twitter.com/tglee) or
[LinkedIn](http://au.linkedin.com/pub/thomas-lee/2/386/629).

