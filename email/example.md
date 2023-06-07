# email #

## Exemplo ##

```go
package main

import (
  "os"

  "github.com/yggbrazil/go-toolbox/env"
	"github.com/yggbrazil/go-toolbox/email"
)

func main() {
	emailService, err := email.New(
		os.Getenv("PROVIDER_EMAIL_SMTP_HOST"),
		env.MustInt("PROVIDER_EMAIL_SMTP_PORT"),
		os.Getenv("PROVIDER_EMAIL_USER"),
		os.Getenv("PROVIDER_EMAIL_PASSWORD"),
	)
	if err != nil {
		return
	}

  m := email.Mail{
    Subject: "Title",
    To:      []string{"andre@yggbrazil.com"},
    Body:    "Body",
  }

  err = emailService.SendEmail(m)
  if err != nil {
    return err
  }
}
```
