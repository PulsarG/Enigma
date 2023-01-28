# Enigma - Encryption system

Work in progress.

Currently available:

> Russian alphabet in lowercase

> all numbers

> some symbols

> space 

> line break (Enter key)

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

* password too long message,

+ another error.

> bool - false

### Decryption is no different from encryption and proceeds in the same way, using the same method, 
but the encrypted text is passed in the parameters, and the decrypted text is returned.
