# File Encryptor
[![Go](https://github.com/gongchen161/file-encryptor/actions/workflows/go.yml/badge.svg)](https://github.com/gongchen161/file-encryptor/actions/workflows/go.yml)

This is a command-line tool to securely encrypt and decrypt a file using AES-GCM.


## Setup and Usage
* Run `go build -ldflags "-s -w" -o file-encryptor` to build the binary executable
* Run `./file-encryptor encrypt -p [password] -file [file_to_encrypt]` to encrypt the file
* Run `./file-encryptor decrypt -p [password] -cipherfile [cipherfile_to_decrypt]` to decrypt the file and get back the original file content

## Supported Flags (<em>More to be added...</em>)
### Flags in Encryption
* `-p` Password (hashed using SHA256) to be used for encryption. <em>Required.</em>
* `-file` File to be encrypted. <em>Required.</em>
* `-o` Output filename containing the encrypted content. <em>Optional. If not supplied, target_filename.encrypted will be used by default</em>

### Flags in Decryption
* `-p` Password (hashed using SHA256) to be used for decryption. Must be the same as the password used in encryption. <em>Required.</em>
* `-cipherfile` File to be decrypted <em>Required.</em>
* `-o` Output filename containing the original/decrypted content. <em>Optional. If not supplied, the original filename will be used by default</em>
