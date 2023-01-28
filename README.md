# Enigma - Encryption system

Work in progress.

Currently available:

> Russian letters in upper and lower case

> English letters in upper and lower case.

> All numbers

> All symbols and characters of the standard keyboard

> Space 

> Line break (Enter key)

> There are 162 letters, signs and symbols in total.

## DOCS

Download and install:
```
go get github.com/PulsarG/Enigma
```

### Start crypt:

```
StartCrypt(text, key string) (string, bool)
```

> text: text for crypt

> key: word password

### Return:

If encryption was successful:

> string: finished ciphertext

> bool: true

Else:

> string: Error messages:

- invalid character message and character,

* message that the password is too long,

+ message that the password is missing (password is required)

> bool - false

### Decryption is no different from encryption and proceeds in the same way, using the same method, 
but the encrypted text is passed in the parameters, and the decrypted text is returned.
