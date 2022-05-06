# luhn
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.png?v=103)](https://opensource.org/licenses/mit-license.php)

Luhn is a go module that uses the luhn aglorithm also known as Modulus 10 Algorithm to verify if a number is a luhn number

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
