# go-worker-pool
[![Build Status](https://travis-ci.org/jayhuang75/wei-helpers.svg?branch=master)](https://travis-ci.org/jayhuang75/wei-helpers) [![codecov](https://codecov.io/gh/jayhuang75/wei-helpers/branch/master/graph/badge.svg)](https://codecov.io/gh/jayhuang75/wei-helpers)
[![Go Report Card](https://goreportcard.com/badge/github.com/jayhuang75/wei-helpers)](https://goreportcard.com/report/github.com/jayhuang75/wei-helpers)

## How to use this?
#### Install package
```bash
$ go get github.com/jayhuang75/wei-helpers
```

#### In your application main.go, import the package
```go
import (
    "github.com/jayhuang75/wei-helpers"
)
```

#### Example how to use
```go
	wei.Logging("info", fmt.Sprintf("[Main] Successful %d players info", len(players)))
    wei.Logging("error", fmt.Sprintf("[Main] Marshal json failed: %s", err.Error()))
```
