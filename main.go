package main

import (
	"fmt"
	"os"
)

func main() {

	flags := processCommandArgs()
	if os.Args[1] == "encrypt" {

		password, ok := flags["-p"]
		if !ok {
			return
		}

		filename, ok := flags["-file"]
		if !ok {
			return
		}

		// optional -o flag for output file name
		outputFilename, ok := flags["-o"]
		if !ok {
			outputFilename = filename + "_encrypted"
		}

		EncryptFile(filename, password, outputFilename)

	} else if os.Args[1] == "decrypt" {
		password, ok := flags["-p"]
		if !ok {
			return
		}

		cipherFilename, ok := flags["-cipherfile"]
		if !ok {
			return
		}

		// optional -o flag for output file name
		outputFilename := flags["-o"]

		DecryptFile(cipherFilename, password, outputFilename)
	} else {
		displayHelp()
	}
}

func displayHelp() {
	fmt.Println("This is a tool to encrypt & decrypt a file")
	fmt.Println()

	fmt.Println("Usage: ")
	fmt.Println("         encrypt [arguments]")
	fmt.Println("         decrypt [arguments]")
	fmt.Println()

	fmt.Println("Valid arguments are: ")
	fmt.Println("         -p                password")
	fmt.Println("         -file             file to be encrypted")
	fmt.Println("         -cipherfile       cipher file to be decrypted")
	fmt.Println("         -o                output filename (optional)")
	fmt.Println()
}

func processCommandArgs() map[string]string {
	flags := make(map[string]string)

	for i := 0; i < len(os.Args)-1; i++ {
		arg := os.Args[i]
		if len(arg) > 0 && arg[0] == '-' {
			flags[arg] = os.Args[i+1]
		}
	}
	return flags
}
