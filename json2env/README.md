# JSON2Env #

É uma lib que le um arquivo json, e coloca os valores achados as variaveis de ambiente para poder ser lido com o pacote "os" do Go em qualquer arquivo 


## Exemplo ##

```go

//test.json
{
  "json":"env"
}

//exemple.go
package main

import (
	"log"
	"os"

	"github.com/yggbrazil/go-toolbox/json2env"
)

func main() {
	if err := json2env.LoadFile("test.json"); err != nil {
		panic(err)
	}

	value := os.Getenv("json")

	log.Println(value)
}
```