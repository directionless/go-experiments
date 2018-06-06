/* Passing args as a struct

Pros:
 - Super convinient
 - Clear

Cons:
 - Defaults may not be what you want
 - Hard to tell absent vars from set-to-default
*/

package main

import (
	"fmt"
)

type Config struct {
	abool   bool
	astring string
}

func doSomething(name string, opts Config) {
	fmt.Printf("%s: %#v\n", name, opts)
}

func main() {
	doSomething("no opts", Config{})
	doSomething("a bool", Config{abool: true})
	doSomething("a str", Config{astring: "hello"})
}
