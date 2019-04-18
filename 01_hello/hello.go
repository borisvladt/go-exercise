// Tutorial https://quii.gitbook.io/learn-go-with-tests/

package main

import "fmt"

const enPrefix = "Hello, "
const french  = "French"
const spanish  = "Spanish"

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {
    switch language {
    case french:
        prefix = "Bonjour, "
    case spanish:
        prefix = "Hola, "
    default:
        prefix = enPrefix
    }

    return
}

func main() {
    fmt.Println(Hello("Boris", ""))
}