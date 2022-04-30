# luhn
A go module that checks Luhn numbers

## Usage

```
go get github.com/gocrazygh/luhn
```

If you are working outside GOPATH:
```
go mod init <module-name>
```

You can simply use it as:
```go
package main

import (
	"fmt"
	"github.com/gocrazygh/luhn"
)

func main() {
	a := luhn.Check("79927398713")
	b := luhn.Check("1111")
	fmt.Println(a)
	fmt.Println(b)
}
```

The output will be:
```
true
false
```
## For Coffee
BTC: `1F5qqrV9bX8Z1eyvy6MBxyVCKnT8cc4Hpc`
