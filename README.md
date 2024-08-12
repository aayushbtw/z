# z

**z** - A simple CLI tool for encoding and decoding strings.

## Installation

To use **z**, you need Go installed on your system. If Go is not already installed, follow the [Go installation page](https://golang.org/doc/install).

Clone the repository and build the CLI tool:

```bash
git clone https://github.com/aayushbtw/z
cd z
go build -o z
```

## Usage

### Encode

To encode a string, use:

```bash
./z -t <algorithm> -e <string>
```

- `-t <algorithm>`: Encoding algorithm. Choices: `hex`, `binary`, `bin`, `base64`, `b64`.
- `-e <string>`: String to encode.

**Example:**

```bash
./z -t base64 -e "Hello, World!"
```

### Decode

To decode a string, use:

```bash
./z -t <algorithm> -d <string>
```

- `-t <algorithm>`: Decoding algorithm. Choices: `hex`, `binary`, `bin`, `base64`, `b64`.
- `-d <string>`: String to decode.

**Example:**

```bash
./z -t base64 -d "SGVsbG8sIFdvcmxkIQ=="
```

### Help

For help and usage information:

```bash
./z -h
```
