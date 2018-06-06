package main

import (
	"fmt"
)

type Config struct {
	abool   bool
	astring string
}

type ConfigOption func(*Config)

func setTrue() ConfigOption {
	return func(c *Config) {
		c.abool = true
	}
}

func setString(s string) ConfigOption {
	return func(c *Config) {
		c.astring = s
	}
}

func doSomething(name string, opts ...ConfigOption) {
	options := &Config{}
	options.astring = "default"

	for _, o := range opts {
		o(options)
	}

	fmt.Printf("%s: %#v\n", name, options)

}

func main() {
	doSomething("no opts")
	doSomething("a bool", setTrue())
	doSomething("a string", setString("hello"))

	doSomething("together", setTrue(), setString("hello"))
}
