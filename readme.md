Cartel
======
[![Build Status](https://travis-ci.org/icambridge/cartel.svg)](https://travis-ci.org/icambridge/cartel)

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
}

func (mo MockOutput) Value() interface{} {
    return nil
}
```

Create your own Task


```go
type MockTask struct {
}

func (mf MockTask) Execute() cartel.OutputValue {
    return 
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