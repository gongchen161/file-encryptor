# File Encryptor
[![Go](https://github.com/gongchen161/file-encryptor/actions/workflows/go.yml/badge.svg)](https://github.com/gongchen161/file-encryptor/actions/workflows/go.yml)

This is a tool to securely encrypt and decrypt a file using AES-GCM. 

You may open this tool as a desktop application (cross-platform, created using *[Fyne](https://github.com/fyne-io/fyne)*), or you may use it via the terminal.

## Setup and Usage
* Run `go build -ldflags "-s -w" -o file-encryptor` to build the binary executable

### To launch the Cross-Platform UI
* Run `./file-encryptor ui`
<img width="858" alt="Screen Shot 2022-01-15 at 9 25 02 PM" src="https://user-images.githubusercontent.com/31603060/149644902-015b8200-f1df-4f3d-9463-3a4957383931.png">


### To Encrypt/Decrypt the File in Terminal
* Run `./file-encryptor encrypt -p [password] -file [file_to_encrypt]` to encrypt the file
* Run `./file-encryptor decrypt -p [password] -cipherfile [cipherfile_to_decrypt]` to decrypt the file and get back the original file content


## Supported Flags in Terminal (<em>More to be added...</em>)
### Flags in Encryption
* `-p` Password (hashed using SHA256) to be used for encryption. <em>Required.</em>
* `-file` File to be encrypted. <em>Required.</em>
* `-o` Output filename containing the encrypted content. <em>Optional. If not supplied, target_filename.encrypted will be used by default</em>

### Flags in Decryption
* `-p` Password (hashed using SHA256) to be used for decryption. Must be the same as the password used in encryption. <em>Required.</em>
* `-cipherfile` File to be decrypted <em>Required.</em>
* `-o` Output filename containing the original/decrypted content. <em>Optional. If not supplied, the original filename will be used by default</em>
