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
    config := whop.Config{
        Bearer: "<Bearer Token>",
	// ClientID: "<client ID">,
    }
	client := whop.NewWhop(config)
	ctx := context.Background()

	resp, _ := client.GetLinks(ctx)
	fmt.Printf("%+v", resp)
}
```
