// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package sample_test

import (
	"fmt"
	"os"

	"codeberg.org/kurth4cker/go-sample"
)

func ExampleFhelloln() {
	sample.Fhelloln(os.Stdout, "kurth4cker")
	// Output: hello kurth4cker
}

func ExampleHelloln() {
	sample.Helloln("Gopher")
	// Output: hello Gopher
}

func ExampleShello() {
	greeting := sample.Shello("world")
	fmt.Println(greeting)
	// Output: hello world
}
