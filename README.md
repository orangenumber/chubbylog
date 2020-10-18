# ChubbyLog

(c) 2020 Gon Yi.
Written by Gon Yi. <https://gonyyi.com/copyright.txt>
Last update: 9/30/2020  

ChubbyLog is a leveled logger based on `alog` <https://github.com/gonyyi/alog>.
ChubbyLog has 4 different print types. (info, warn, error, fatal)
And each has its own output writer.


## Usage

Without creating an instance

```go
package main

import "github.com/orangenumber/chubbylog"

func main() {
	l := chubbylog.New(os.Stdout, chubbylog.F_STD)  // chubbylog.F_STD is a standard flag (F_TIME | F_DATE | F_PREFIX)
}
```


## Example


```go
package main

import "github.com/orangenumber/chubbylog"

func main() {
	out, err := os.Create("test.log")
	if err != nil {
		println(err.Error())
	}

	x := chubbylog.New(out, chubbylog.F_STD)

	x.Infof("OK: %s", "blah blah")
	x.Warnf("HMM: %s", "blah blah")
	x.Errorf("Oops: %s", "blah blah")
	x.Fatalf("Dang it: %s", "blah blah")
}
```
