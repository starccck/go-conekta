[![Build Status](https://travis-ci.org/sait/go-conekta.svg?branch=master)](https://travis-ci.org/sait/go-conekta)

### Go-Conekta
A Wrapper for use conekta's api v2 in golang

This tutorial assumes the next:

* Have a conekta account
* Have a frontend that tokenize cards
* Knowledge of conekta's api

### Usage

First get the package

```
go get -u github.com/sait/go-conekta
```

Import in your project

```go
import (
    "github.com/sait/go-conekta/conekta"
)
```

### Simple Order Example

```go
conekta.ApiKey = "<My Secret ApiKey>"
order := new(conekta.Order)
item := conekta.LineItem{
    Name:        "Awesome item",
    Description: "Super Awesome item",
    UnitPrice:   20000,
    Quantity:    2,
}
order.LineItems = append(order.LineItems, item)
order.Currency = "MXN"
charge := conekta.Charge{
    PaymentMethod: conekta.PaymentMethod{
        Type:    "card",
        TokenId: "<token generated by frontend>",
    },
}
order.Charges = append(order.Charges, charge)
order.Metadata = conekta.Metadata{
    "test": "extra_info",
    "hola": "mundo",
}
order.CustomerInfo.Name = "Fulanito Pérez"
order.CustomerInfo.Email = "fulanito@conekta.com"
order.CustomerInfo.Phone = "+52181818181"
statusCode, conektaError, conektaResponde := order.Create()
if statusCode != 200 {
    fmt.Println("There's a problem :(")
    panic(conektaError)
}
fmt.Println(conektaResponde)
fmt.Println("Congratulations!!")
```

### Resources

https://developers.conekta.com/libraries/javascript

https://developers.conekta.com/api

https://developers.conekta.com/tutorials/card
