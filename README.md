# whop-go-sdk

## Installation
```
go get github.com/whopio/whop-go-sdk
```

## Usage
```
package main

import (
	"context"
	"fmt"

	"github.com/whopio/whop-go-sdk"
)

func main() {
	client := whop.NewWhop("<Bearer Token>")
	ctx := context.Background()

	resp, _ := client.GetLinks(ctx)
	fmt.Printf("%+v", resp)
}
```