# go-get-org
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

Overview

## :memo: Description

`go get` all repositories of the specified organization.

## :package: Installation

```
$ go get github.com/micnncim/go-get-org
```

## :mag: Usage

```sh
$ go-get-org <ORGANIZATION> <GITHUB_ACCESS_TOKEN>
$ # recommend to use itchyny/fillin
$ fillin go-get-org <ORGANIZATION> {{token}}
# execute the following commands inside
# go get -u <ORGANIZATION>/<REPOSITORY1>/...
# go get -u <ORGANIZATION>/<REPOSITORY2>/...
# ...
```

## :bulb: Example

```
go-get-org uber-go $GITHUB_TOKEN
installing the following repositories (22)
uber-go/zap
uber-go/gwr
uber-go/atomic
uber-go/sally
uber-go/hackeroni
uber-go/flagoverride
uber-go/tally
uber-go/ratelimit
uber-go/fx
uber-go/dosa
uber-go/cadence-client
uber-go/tools
uber-go/multierr
uber-go/dig
uber-go/mapdecode
uber-go/automaxprocs
uber-go/icu4go
uber-go/config
uber-go/goleak
uber-go/protoidl
uber-go/go-helix
uber-go/kafka-client
Installing uber-go/zap SUCCEEDED
Installing uber-go/gwr SUCCEEDED
Installing uber-go/atomic SUCCEEDED
Installing uber-go/sally timeout
Installing uber-go/hackeroni SUCCEEDED
Installing uber-go/flagoverride SUCCEEDED
Installing uber-go/tally SUCCEEDED
Installing uber-go/ratelimit SUCCEEDED
Installing uber-go/fx SUCCEEDED
Installing uber-go/dosa SUCCEEDED
Installing uber-go/cadence-client SUCCEEDED
Installing uber-go/tools SUCCEEDED
Installing uber-go/multierr SUCCEEDED
Installing uber-go/dig SUCCEEDED
Installing uber-go/mapdecode SUCCEEDED
Installing uber-go/automaxprocs timeout
Installing uber-go/icu4go SUCCEEDED
Installing uber-go/config timeout
Installing uber-go/goleak SUCCEEDED
Installing uber-go/protoidl SUCCEEDED
Installing uber-go/go-helix SUCCEEDED
Installing uber-go/kafka-client timeout

installed repositories: 18
not installed repositories: 4

the following repositories not installed
uber-go/sally
uber-go/config
uber-go/automaxprocs
uber-go/kafka-client
```

## :bust_in_silhouette: Author

[@micnncim](https://twitter.com/micnncim)

## :credit_card: License

[MIT](./LICENSE)