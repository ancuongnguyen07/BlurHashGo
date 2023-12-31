# BlurHashGo
BlurHashGo is a Go implementation of the BlurHash algorithm, which allows you to create compact representations of
images that can be used as placeholders with a low-res preview. Blurhash is originally created by **[Wolt](https://github.com/woltapp/)** folks.

Reference:
- https://github.com/woltapp/blurhash
- https://github.com/bbrks/go-blurhash
- https://github.com/buckket/go-blurhash

### Overview
BlurHash is a clever algorithm that encodes the average color of each 8x8 pixel block of an image into a short string.
This string can then be used as a placeholder for the image, providing a blurry version of it before the actual image is loaded.

This Go implementation, BlurHashGo, allows you to easily encode and decode BlurHash strings in your Go applications.

### Features
- Encoding: Convert an image into a BlurHash string.
- Decoding: Decode a BlurHash string into representative image.

## Getting Started BlurHash Package
### Installation
```bash
go get -u github.com/ancuongnguyen07/BlurHashGo
```

### Usage
```go
import (
    "fmt"
    "image"
    "os"

    "github.com/ancuongnguyen07/BlurHashGo"
)

func main() {
    // Load an image
    imgFile, err := os.Open("example.jpg")
    if err != nil {
        fmt.Println("Error opening image:", err)
        return
    }
    defer imgFile.Close()

    img, _, err := image.Decode(imgFile)
    if err != nil {
        fmt.Println("Error decoding image:", err)
        return
    }

    // Encode the image to BlurHash
    hash, err := BlurHashGo.Encode(img, 4, 3)
    if err != nil {
        fmt.Println("Error encoding image:", err)
        return
    }

    // Print the BlurHash string
    fmt.Println("BlurHash:", hash)
}
```
For more detailed documentation, refer to [Godoc](https://pkg.go.dev/github.com/ancuongnguyen07/BlurHashGo)

## Getting Started BlurHash CLI
### Installation
You can build locally by running the following commands:
```bash
git clone https://github.com/ancuongnguyen07/BlurHashGo & cd BlurHashGo
cd cli
make
./build/blurhash-cli # run the executable binary file
```
The executable file `blurhash-cli` should be available in the folder `build`.

Or using a more convenient way:
```bash
go install github.com/ancuongnguyen07/BlurHashGo/cli/blurhash-cli
```

### Usage
The CLI tool provides subcommands for encoding and decoding BlurHash strings.

#### Encode
The `encode` subcommand allows you to encode an image to a BlurHash string.
Example:
```bash
blurhash-cli encode --file example.jpg --xcomp 4 --ycomp 3
```
This command encodes the image located in the file "example.jpg" using a grid of 4x3 components.

#### Decode
The `decode` subcommnd allows you to decode a BlurHash string into an image.
Example:
```bash
blurhash-cli decode --hash LKO2?U%2Tw=w]~RBVZRi};RPxuwH --width 400 --height 800 --punch 1 --dest ./result.png
```
This command decodes the BlurHash string `LKO2?U%2Tw=w]~RBVZRi};RPxuwH` and saves the resulting `result.png` image.

#### Help
To show the help menu of `blurhash-cli`, run:
```bash
$ blurhash-cli --help
Blurhash command line tool

Usage:
  blurhash [flags]
  blurhash [command]

Examples:

        blurhash encode --filepath ./smile.png --xcomponent 4 --ycomponent 3


Available Commands:
  completion  Generate the autocompletion script for the specified shell
  decode      Decode an image from local file or downloaded url
  encode      Encode an image from local file or downloaded url
  help        Help about any command

Flags:
  -h, --help      help for blurhash
  -v, --version   version for blurhash

Use "blurhash [command] --help" for more information about a command.
```

### License
This project is licensed under the MIT Licence -- see the [LICENCE](./LICENCE) file for details.