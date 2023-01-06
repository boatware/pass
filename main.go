package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"math/rand"
	"os"
	"time"
)

type Options struct {
	Length           int  `short:"l" long:"length" description:"Length of the generated password" required:"true" default:"24"`
	ExcludeNumbers   bool `short:"n" long:"exclude-numbers" description:"Exclude numbers in the generated password"`
	ExcludeUppercase bool `short:"u" long:"exclude-uppercase" description:"Exclude uppercase letters in the generated password"`
	ExcludeLowercase bool `short:"c" long:"exclude-lowercase" description:"Exclude lowercase letters in the generated password"`
	ExcludeSymbols   bool `short:"s" long:"exclude-symbols" description:"Exclude symbols in the generated password"`
	Iterations       int  `short:"i" long:"iterations" description:"Number of iterations to perform" default:"1"`
}

var opts Options
var parser *flags.Parser

func main() {
	parser = flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

	for i := 0; i < opts.Iterations; i++ {
		fmt.Println(GeneratePassword(opts.Length, opts.ExcludeNumbers, opts.ExcludeUppercase, opts.ExcludeLowercase, opts.ExcludeSymbols))
	}
}

func GeneratePassword(length int, excludeNumbers bool, excludeUppercase bool, excludeLowercase bool, excludeSymbols bool) string {
	base := ""
	if !excludeNumbers {
		base += "0123456789"
	}

	if !excludeUppercase {
		base += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if !excludeLowercase {
		base += "abcdefghijklmnopqrstuvwxyz"
	}

	if !excludeSymbols {
		base += "!@#$%^&*()_+-=[]{}|;':,./<>?`~"
	}

	if base == "" {
		fmt.Println("Congratulations, you just tried to generate a password with no available characters.")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100000)))
	var password string
	for i := 0; i < length; i++ {
		password += string(base[rand.Intn(len(base))])
	}
	return password
}
