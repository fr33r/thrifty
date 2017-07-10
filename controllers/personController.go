package controllers

import(
  "thrifty/gen-go/thrifty"
)

type PersonController interface {
  Get() (*thrifty.Person, error)
}
