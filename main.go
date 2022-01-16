package main

import (
	"errors"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {

	flags := processCommandArgs()
	if os.Args[1] == ENCRYPT {

		password, ok := flags["-p"]
		if !ok {
			fmt.Println("Missing password. Please add -p")
			displayHelp()
			return
		}

		filename, ok := flags["-file"]
		if !ok {
			fmt.Println("Missing encryption file. Please add -file")
			displayHelp()
			return
		}

		// optional -o flag for output file name
		outputFilename, ok := flags["-o"]
		if !ok {
			outputFilename = filename + "_encrypted"
		}

		res, success := EncryptFile(filename, password, outputFilename)

		fmt.Println(res)
		if success {
			fmt.Println("Completed.")
		} else {
			fmt.Println("Please try again")
		}
	} else if os.Args[1] == DECRYPT {
		password, ok := flags["-p"]
		if !ok {
			fmt.Println("Missing password. Please add -p")
			displayHelp()
			return
		}

		cipherFilename, ok := flags["-cipherfile"]
		if !ok {
			fmt.Println("Missing cipher filename. please add -cipherfile")
			displayHelp()
			return
		}

		// optional -o flag for output file name
		outputFilename := flags["-o"]

		res, success := DecryptFile(cipherFilename, password, outputFilename)

		fmt.Println(res)
		if success {
			fmt.Println("Completed without errors")
		} else {
			fmt.Println("Please try again")
		}
	} else if os.Args[1] == "ui" {

		showUI()

	} else {
		fmt.Println("Invalid flags")
		displayHelp()
	}
}

func displayHelp() {
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------")
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

func showUI() {

	fyneApp := app.New()
	fyneWindow := fyneApp.NewWindow("File Encryptor")
	var filePath fyne.URIReadCloser
	var fileButton *widget.Button

	fileDialog := dialog.NewFileOpen(func(path fyne.URIReadCloser, err error) {
		if path == nil {
			return
		}
		filePath = path
		fileButton.SetText(filePath.URI().Path())
	}, fyneWindow)
	fileDialog.Resize(fyne.NewSize(760, 420))

	fileButton = widget.NewButton("Select File", func() {
		fileDialog.Show()
	})

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	outputFilename := widget.NewEntry()

	emptyValidator := func(s string) (err error) {
		if len(s) <= 0 {
			return errors.New("Cannot be empty")
		}

		return nil
	}
	outputFilename.Validator = emptyValidator
	password.Validator = emptyValidator

	options := widget.NewRadioGroup([]string{ENCRYPT, DECRYPT}, func(string) {})
	options.Horizontal = true
	options.SetSelected(ENCRYPT)
	resultText := widget.NewMultiLineEntry()
	resultText.Disable()
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Input File", Widget: fileButton, HintText: "Input file to be encrypted/decrypted"},
			{Text: "Password", Widget: password, HintText: "Password to be used for encryption/decryption"},
			{Text: "Output File", Widget: outputFilename, HintText: "Output filename"},
		},
		OnSubmit: func() {
			selection := options.Selected
			resultText.SetText("Start " + selection + "ing " + filePath.URI().Path())
			var res string
			var success bool
			if selection == ENCRYPT {
				res, success = EncryptFile(filePath.URI().Path(), password.Text, outputFilename.Text)
			} else {
				res, success = DecryptFile(filePath.URI().Path(), password.Text, outputFilename.Text)
			}
			resultText.SetText(resultText.Text + "\n" + res)
			if success {
				resultText.SetText(resultText.Text + "\nCompleted without errors")
			} else {
				resultText.SetText(resultText.Text + "\nPlease try again")
			}
		},
	}
	form.SubmitText = "Start"
	form.Append("Options", options)
	form.Append("Result", resultText)

	fyneWindow.SetContent(container.NewVBox(
		widget.NewLabel("File Encryptor"),
		form,
	))
	fyneWindow.Resize(fyne.NewSize(860, 460))
	fyneWindow.SetFixedSize(true)
	fyneWindow.ShowAndRun()
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
