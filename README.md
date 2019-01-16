# wechat-jssdk

### usage

```GO
package main

import (
	"fmt"

	jssdk "github.com/limtech/wechat-jssdk"
)

const (
	WECHAT_APP_ID     string = "xxxxxxxxx"
	WECHAT_APP_SECRET string = "xxxxxxxxx"
)

func main() {
	sdk := jssdk.New(WECHAT_APP_ID, WECHAT_APP_SECRET)
}
```