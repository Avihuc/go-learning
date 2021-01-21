package talk

import "rsc.io/quote"

const (
	hello = "Hello go!!"
	curse = "Fuck off!!"
)

func SayHello() string {
	return quote.Hello()
}

func Curse() string {
	return curse
}
