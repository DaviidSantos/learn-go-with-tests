package hello

import "fmt"

const (
	french               = "French"
	spanish              = "Spanish"
	norwegian            = "Norwegian"
	englishHelloPrefix   = "Hello, "
	frenchHelloPrefix    = "Bonjour, "
	spanishHelloPrefix   = "Hola, "
	norwegianHelloPrefix = "Hei, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case norwegian:
		prefix = norwegianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("", ""))
}
