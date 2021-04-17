# Go App Skeleton

Originally the work of the [AdhocTeam](https://github.com/adhocteam/GolangAppSkeleton) (github). This skeleton Golang Web Application allows developers to quickly get started with a web application build in Golang.

This version upgrades the original with embed.FS usage for both Templates and Static Files, and removes the `STATIC_BASE` environment variable.

## Prerequisites

* [Go](https://golang.org) 1.16 or greater
* Optionally, [modd](https://github.com/cortesi/modd#install) for live reload

## How To Use

* Fork this repository
* Change the package name in `go.mod` to reflect your own repository
* run `modd`
    * if you chose not to install modd, build the app with `go build -o your_program_name`
* modify code to see the results

## Environment Variables

* `PORT`: the port for the app to listen on. Defaults to 8080
