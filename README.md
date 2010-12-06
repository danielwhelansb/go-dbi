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
        "dbi"
        "log"
    )

    func main() {
        conn, err := dbi.Connect("mysql://root@localhost/database")
        if err != nil {
            log.Printf("error: unable to connect to the database: %s", err.String())
        }
        defer conn.Close()
    }

## License

This software is licensed under the terms of the [MIT License](http://en.wikipedia.org/wiki/MIT_License).

## Support

Please log defects and feature requests using the issue tracker on github.

## About

go-dbd-mysql was written by [Tom Lee](http://tomlee.co).

Follow me on [Twitter](http://www.twitter.com/tglee) or
[LinkedIn](http://au.linkedin.com/pub/thomas-lee/2/386/629).

