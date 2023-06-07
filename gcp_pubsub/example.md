# gcp_pubsub #

gcp_pubsub é um wrapper do [GCP PubSub](https://pkg.go.dev/cloud.google.com/go/pubsub) com abstração das funções

**Váriaveis de ambiente**

```code
GOOGLE_APPLICATION_CREDENTIALS=config/google_credentials.json
PUBSUB_PROJECT_ID=project-id
```

**Váriaveis de ambiente de sugestão para os exmplos**

```code
PUBSUB_TOPIC=my.topic
PUBSUB_SUBSCRIPTION_NAME=my.topic.subscription
```

## Exemplo ##

```go
package main

import (
	"fmt"
	"os"

	"github.com/yggbrazil/go-toolbox/gcp_pubsub"
)

func main() {
  // Services
  pubusb := gcp_pubsub.NewPubSub()

  topic := os.Getenv("PUBSUB_TOPIC")

  msg := struct{}{
    Name: "Andre",
    Age: 32
  }

  // To Publish
  pubsub.Publish(topic, msg)


  // To Subscribre
  subscription_name := os.Getenv("PUBSUB_SUBSCRIPTION_NAME")

  pubsub.Subscribre(topic, subscription_name, func(message string) error {
    fmt.Println(message)
    return nil
  })
}
```
