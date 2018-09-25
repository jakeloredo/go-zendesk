# go-zendesk
Golang wrapper for Zendesk's REST API

## Getting Started

### Requirements

None.

### Installing

Run this command to install from my Github repo:

```
go get https://github.com/jakeloredo/go-zendesk
```

## Usage

```go
import (
  "fmt"
  "github.com/jakeloredo/go-zendesk/"
)

func main() {
    
    // Setup Zendesk client
    zd := Zendesk{
        Username: "email@domain.com",
        Password: "Password",
        Subdomain: "subdomain",
    }
  
    // Get Zendesk ticket #123456
    ticket, err := zd.GetOneTicket(123456)
    
    if err != nil {
        panic(err)
    }
    fmt.Printf(ticket)
    
    // Assign ticket #123456 to zd.Username and submit the ticket as solved
    _, solveErr := zd.SolveTicket(123456)
}
```

