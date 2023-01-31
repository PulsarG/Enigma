# Enigma - Encryption system

Work in progress. ver 1.1.0

> Implemented the generation of a new (individual) rotor for encryption. Provides additional protection.

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

### Generate New Rotor

> The method must be called BEFORE the Start Encryption ( StartCrypt() ).

>If the new Rotor is used in encryption, but not saved in a file, then it is not possible to decrypt the data in the future.

`NewRotor()`

> return array []int: for save in file

> and return boolean: true - if new rotor generate without errors, else return false

### Decryption is no different from encryption and proceeds in the same way, using the same method, 
but the encrypted text is passed in the parameters, and the decrypted text is returned.
