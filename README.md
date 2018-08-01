# kit

EyeCueLab's `kit` is a repository of open-source golang libraries developed by EyeCueLab for use in our various projects.

## philosophy / notes

EyeCueLab tries to follow the philosphy of the [twelve-factor app](https://12factor.net).  All non-business logic should be open & accessible to the public, and configuration details should be in the environment.

Accordingly, EyeCueLab uses the [cobra command-line interface](github.com/spf13/cobra) and the [viper configuration library](https://github.com/spf13/viper). Many parts of `kit` either rely on these libraries for configuration or hook into them during init() so that cobra/viper can access them.

## Testing

Navigate to the root of the source directory for kit

```bash
go test ./... -cover
```

## address

`address` contains tools for storing and comparing street addreses.

## assets

(TODO)

## brake

`brake` contains tools for setting up and working with the [airbrake](https://airbrake.io/) error monitoring software.

## branch

(TODO)

## cmd

`cmd` contains helper fucntions for the [cobra command-line interface](https://github.com/spf13/cobra)

## coerce

`coerce` contains functions which use reflection to coerce various numeric types using c-style numeric promotion. This is mostly for testing.

## config

`config` contains helper functions to help set up [viper configuration library](https://github.com/spf13/viper)

## counter

`counter` implements various counter types, similar to python's collections.counter

## db

`db` contains tools for various databases. See `db/mongo` or `db/psql`.

## db/mongo

`db/mongo` helps connect to our mongoDB instances and provides various helper functions 

## db/psql

`db/psql` helps connect to psql instances and provides various helper functions

## dockerfiles

(TODO)

## env

`env` contains tools for accessing and manipulating environment variables

## errorlib

`errorlib` contains tools to deal with errors, including the heavily-used `ErrorString` type (for errors that are compile-time constants) and `LoggedChannel` function (to report non-fatal errors during concurrent execution)

## flect

`flect` is meant to work alongside go's `reflect` library, containing additional runtime reflection tools

## geojson

`geojson` contains an interface and various structs for 2d and 3d geometry that correspond to the [geojson](http://geojson.org/) format.

## goenv

(TODO)

## imath

`imath` is contains tools for signed integer math. it largely corresponds with go's built in `math` library for float64s

### imath/operator

`imath/operator` contains functions which represent all of golang's built in operators,as well as bitwise operators

## log

`log` is an extension of the [logrus](https://github.com/sirupsen/logrus) package. It contains various tools for logging information and errors during runtime.

## mailman

(TODO)

## maputil

`maputil` contains various helper functions for maps with string keys.

the maputil package itself covers `map[string]interface{}`

### maputil/string

`maputil` contains helper functions for the type `map[string]string`

## oauth

(TODO)

## pretty

`pretty`  provides pretty-printing for go values. It is a fork of the [kr/pretty](https://github.com/kr/pretty) package, obtained under the MIT license.

## random

`random` provides tools for generating cryptographically secure random elements. it uses golang's built in `crypto/rand` for it's RNG.

### random/shuffle

`random/shuffle` provides tools for creating shuffled _copies_ of various golang slice types using the fisher-yates shuffle. It uses `crypto/rand` for it's RNG.

## retry

(TODO)

## runeset

`runeset` implements a set of runes and various methods for dealing with strings accordingly. it should probably be folded into the `set` package.

## s3

`s3` contains helper functions for dealing with amazon's aws/s3 and integrating it with the cobra CLI.

## set

`set` implements various set types and associated methods; union, intersection, containment, etc. see `runeset` for a set of runes (this will be folded into set.)

## str

`str` contains various tools for manipulation of strings beyond those available in golang's `strings` library. it also contains a wide variety of string constants.

## stringslice

`stringslice` contains various functions to work with slices of strings and the strings contained within.

### set/int

`set/int` implements a set type for `int`

### set/string

`set/string` implements a set type for `string`

## sortlib

`sortlib` contains tools for producing sorted _copies_ of various golang types

## tickertape

`tickertape` provides an implenetation of a concurrency-safe 'ticker tape' of current information during  a running program - that is, repeatedly updating the same line with new information.

## tsv

`tsv` contains tools for dealing with tab-separated values.

## umath

`umath` contains various math functions for unsigned integers, roughly corresponding with `imath` and golang's built in `math` package.

## web

`web` is the bones of our web framework, built on top of google's [jsonapi framework](https://github.com/eyecuelab/jsonapi) and labstack's [echo framework](https://github.com/labstack/echo).

### web/middleware

(TODO)

### web/server

(TODO)

### web/testing

(TODO)

### web/webtest

(TODO)