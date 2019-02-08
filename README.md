# Mongodb Hooks for [Logrus](https://github.com/sirupsen/logrus) <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:"/>

## Install

```shell
$ go get github.com/bregydoc/mgorus
```

## Usage

```go
package main

import (
	"fmt"
	
	"github.com/sirupsen/logrus"
	"github.com/bregydoc/mgorus"
	
)

func main() {
	log := logrus.New()
	hooker, err := mgorus.NewHooker("localhost:27017", "db", "collection")
	if err == nil {
	    log.Hooks.Add(hooker)
	} else {
		fmt.Print(err)
	}

	log.WithFields(logrus.Fields{
		"name": "zhangsan",
		"age":  28,
	}).Error("Hello world!")
}
```

With a pre-existing collection
```go
package main

// Work in progress

```

## License
*MIT*
