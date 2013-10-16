# stateful

Prototype of a state machine for Go.


## Installation

With Google's [Go](http://www.golang.org) installed on your machine:

    $ go get -u github.com/chuckpreslar/stateful

## Usage


```go
package main

import (
  "fmt"
)

import (
  "github.com/chuckpreslar/stateful"
)

const (
  StateOne stateful.State = iota
  StateTwo
)

var UserStates = stateful.NewStateMachine(StateOne).
  BeforeTransition(StateOne, StateTwo, func(user *User) {
    // Transitioned from StateOne to StateTwo, alert the media.
    fmt.Println("Transitioned")
  })

type User struct {
  *stateful.StateMachine
  Name string
}

func NewUser(name string) *User {
  user, err := stateful.Process(&User{Name: name})
  
  if nil != err {
    panic(err)
  }
  
  return user.(*User)
}

func main() {
  user := NewUser("John")
  
  // Do somethings with the user in its current state..
  user.Transition(StateTwo)
}

```

## Documentation

View godocs or visit [godoc.org](http://godoc.org/github.com/chuckpreslar/stateful).

    $ godoc stateful
    
## License

> The MIT License (MIT)

> Copyright (c) 2013 Chuck Preslar

> Permission is hereby granted, free of charge, to any person obtaining a copy
> of this software and associated documentation files (the "Software"), to deal
> in the Software without restriction, including without limitation the rights
> to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
> copies of the Software, and to permit persons to whom the Software is
> furnished to do so, subject to the following conditions:

> The above copyright notice and this permission notice shall be included in
> all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
> THE SOFTWARE.
