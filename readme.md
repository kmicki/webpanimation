# webpanimation  
Package is binding to libwebp providing methods to create webp animations from golang `image.Image` interface.
It also allows decoding of webp animations into single frames that implement image.Image

## Origin
This repo is a fork of github.com/size-of-int/webpanimation. 

It adds decoding capabilities to already existing encoding capabilities.

## Installing
`go get github.com/kmicki/webpanimation`

## Examples
Check out [examples](examples) folder.

No examples of decoding yet.

## Dependencies
The only dependency is libwebp, it source code are embeded in package so no additional installations are needed on machine. It can be updated to latest libwebp via ```./generate_from_libwebp.py /path/to/clean/libwebp/src/```
