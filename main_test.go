// main_test.go
package main

import "testing"

func TestHelloWorld(t *testing.T) {
    result := HelloWorld()
    expected := "Hello, World!"

    if result != expected {
        t.Errorf("Expected '%s' but got '%s'", expected, result)
    }
}
