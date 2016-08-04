// Validate strings with the luhn (mod-10) algorithm
package main

import(
    "fmt"
    "os"
    "github.com/atarantini/luhn/lib"
)

var usage string = `Luhn can validate string for luhn algorithm.

Usage:

        ./luhn 4242424242424242
        4242424242424242 is valid

        ./luhn 4242424242424241
        4242424242424242 is NOT valid

`

func main() {
    // Get PANs from command line arguments
    args := os.Args[1:]

    // Validate PANs
    for _,pan := range args {
        if luhn.Validate(pan) {
            fmt.Println(pan, "is valid")
        } else {
            fmt.Println(pan, "is NOT valid")

        }
    }
}
