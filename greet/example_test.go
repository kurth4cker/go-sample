// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package greet_test

import (
	"fmt"
	"os"

	"github.com/kurth4cker/go-sample/greet"
)

func ExampleFhelloln() {
	greet.Fhelloln(os.Stdout, "kurth4cker")
	// Output: hello kurth4cker
}

func ExampleHelloln() {
	greet.Helloln("Gopher")
	// Output: hello Gopher
}

func ExampleShello() {
	greeting := greet.Shello("world")
	fmt.Println(greeting)
	// Output: hello world
}
