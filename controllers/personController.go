package controllers

import(
  "thrifty/gen-go/person"
)

type PersonController interface {
  Get() (*person.Person, error)
}
