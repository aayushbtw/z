package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func encodeBase64(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func decodeBase64(value string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func encodeHex(value string) string {
	return hex.EncodeToString([]byte(value))
}

func decodeHex(value string) (string, error) {
	decoded, err := hex.DecodeString(value)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func encodeBinary(value string) string {
	var binaryStr strings.Builder
	for _, b := range []byte(value) {
		binaryStr.WriteString(fmt.Sprintf("%08b", b))
	}
	return binaryStr.String()
}

func decodeBinary(value string) (string, error) {
	var decoded []byte
	for i := 0; i < len(value); i += 8 {
		var byteValue byte
		_, err := fmt.Sscanf(value[i:i+8], "%08b", &byteValue)
		if err != nil {
			return "", err
		}
		decoded = append(decoded, byteValue)
	}
	return string(decoded), nil
}

func main() {
	app := &cli.App{
		Name:  "z",
		Usage: "A simple CLI tool for encoding and decoding strings.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "t",
				Usage:    "Specify the algorithm for encoding/decoding. Choices are: hex, binary, bin, base64, b64",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "d",
				Usage: "Specify the value to decode. Use with -t for decoding.",
			},
			&cli.StringFlag{
				Name:  "e",
				Usage: "Specify the value to encode. Use with -t for encoding.",
			},
		},
		Action: func(c *cli.Context) error {
			tValue := c.String("t")
			dValue := c.String("d")
			eValue := c.String("e")

			validAlgorithms := map[string]struct{}{
				"hex":    {},
				"binary": {},
				"bin":    {},
				"base64": {},
				"b64":    {},
			}

			if _, valid := validAlgorithms[tValue]; !valid {
				return cli.Exit(fmt.Sprintf("Error: Invalid value for -t flag. Must be one of: hex, binary, bin, base64, b64"), 1)
			}

			if dValue == "" && eValue == "" {
				return cli.Exit("Error: One of -d or -e must be provided", 1)
			}

			var result string
			var err error

			switch tValue {
			case "base64", "b64":
				if dValue != "" {
					result, err = decodeBase64(dValue)
				} else {
					result = encodeBase64(eValue)
				}

			case "hex":
				if dValue != "" {
					result, err = decodeHex(dValue)
				} else {
					result = encodeHex(eValue)
				}

			case "binary", "bin":
				if dValue != "" {
					result, err = decodeBinary(dValue)
				} else {
					result = encodeBinary(eValue)
				}

			default:
				return cli.Exit("Error: Algorithm not implemented", 1)
			}

			if err != nil {
				return cli.Exit(fmt.Sprintf("Error: %v", err), 1)
			}

			fmt.Println(result)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
