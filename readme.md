Cartel
======
[![Build Status](https://travis-ci.org/icambridge/cartel.svg)](https://travis-ci.org/icambridge/cartel)[![Coverage Status](https://img.shields.io/coveralls/icambridge/cartel.svg)](https://coveralls.io/r/icambridge/cartel)

A task process pool to allow for the easy creation of a pool of workers.

Installation
------------

The recommended way to install go-dependency

```go
    get github.com/icambridge/cartel
```

Examples
--------

How import the package

```go
import (
    "github.com/icambridge/cartel"
)
```

Create your own output

```go
type MockOutput struct {
    Name string
}

func (mo MockOutput) Value() interface{} {
    return mo.Name
}
```

Create your own Task


```go
type MockTask struct {
    Name string
}

func (mt MockTask) Execute() cartel.OutputValue {
    return MockOutput{mt.Name}
}
```

Then to use

```go
p := cartel.NewPool(1)

task := MockTask{"Iain"}

p.Do(task)
p.End()
value := <-p.Output
```