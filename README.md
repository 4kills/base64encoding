# base64encoding

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A small base64 encoding/decoding library for go that offers a variety of default charsets but also the option 
to use your very own, user-defined charset.

Please be aware that there is still a lot of potential for perfomance improvements. 
Feel free to contribute :yellow_heart:. 

# Usage 

An encoder is used for both, encoding and decoding. Every encoder you create is tied to its very own, immutable charset. 
This ensures consistency when encoding/decoding with the same encoder. 

### Creating an Encoder:

You may create an encoder in one of the following ways:

- encoder with the default charset that is **url-safe** but **non-standard**:
  ```go 
  encoder := base64encoding.New()
  ```
- encoder with the **standard** base64 charset (e.g. used for encoding pictures in base64 in HTML) that is **not url-safe**:
  ```go 
  encoder := base64encoding.NewWeb()
  ```
- encoder with a custom charset (or a constant of this library that wasn't given its own constructor) that is defined by the user. 
  Note that the charset must be comprised of exactly 64 pairwise distinct (extended) ASCII characters:
  ```go 
  myCharset := base64encoding.EasilyReadableCodeSet  
  // or myCharset := "my_custoM-CharSeT+0123456789..." (omitting the rest)
  
  encoder, err := base64encoding.NewCustom(myCharset)
  if err != nil {
    // handle error caused by illegal charset (e.g. not pairwise distinct)
  }
  ```
  
Once you have created your encoder you may use it like that: 

### Encode:

```go 
data := []byte("some data: could be a picture, a UUID or whatever") 
stringB64encoded := encoder.Encode(data) 
```

### Decode:

```go 
data, err := encoder.Decode(stringB64encoded)
if err != nil {
  // handle error (e.g. string not encoded with the encoder's charset)
}
```

# Notes

- **encoders are thread-safe**. You may use the same encoder accross mutliple threads simultaneously, as the charset is read-only.
