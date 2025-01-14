package Converters

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ConvertShellcode2Hex
func ConvertShellcode2Hex(shellcode string, language string) (string, int) {
	// Convert raw shellcode to hexadecimal
	hexShellcode := hex.EncodeToString([]byte(shellcode))

	// Split hex shellcode into individual hex values
	hexValues := strings.Split(hexShellcode, "")

	formattedHexShellcode := ""

	// Format and add "0x" in front of each pair of hex characters
	for i := 0; i < len(hexValues); i += 2 {
		formattedHexShellcode += "0x" + hexValues[i] + hexValues[i+1]
		if i < len(hexValues)-2 {
			formattedHexShellcode += ", "
		}
	}

	// Calculate shellcode size in bytes
	shellcodeSize := len(shellcode)

	return formattedHexShellcode, shellcodeSize
}

// ConvertShellcode2Template function
func ConvertShellcode2Template(shellcode string, language string, length int, variable string) string {
	switch language {
	case "c":
		template := fmt.Sprintf(`unsigned char %s[] = "%s";`, variable, shellcode)
		return template
	case "csharp":
		template := fmt.Sprintf(`byte[] %s = new byte[%d] {%s};`, variable, length, shellcode)
		return template
	case "nim":
		template := fmt.Sprintf(`var %s: array[%d, byte] = [byte %s]`, variable, length, shellcode)
		return template
	case "rust":
		template := fmt.Sprintf(`let %s: [u8; %d] = [%s];`, variable, length, shellcode)
		return template
	default:
		fmt.Println("[!] Unsupported programming language:", language)
		os.Exit(1)
		return ""
	}
}

// ConvertShellcode2String function
func ConvertShellcode2String(shellcodePath string) (string, error) {
	// Read the contents of the file into a byte slice
	fileContent, err := ioutil.ReadFile(shellcodePath)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string
	rawShellcode := strings.TrimSpace(string(fileContent))

	return rawShellcode, nil
}

// FormatKeysToHex function
func FormatKeysToHex(byteArray []byte) string {
	var hexBytes []string
	for _, byteVal := range byteArray[:len(byteArray)-1] {
		hexBytes = append(hexBytes, fmt.Sprintf("0x%02x", byteVal))
	}
	hexBytes = append(hexBytes, fmt.Sprintf("0x%02x", byteArray[len(byteArray)-1]))

	return strings.Join(hexBytes, ", ")
}

// FormatShellcode function
func FormatShellcode(encryptedShellcode []byte) string {
	var formattedShellcode []string
	for _, b := range encryptedShellcode {
		formattedShellcode = append(formattedShellcode, fmt.Sprintf("0x%02x", b))
	}

	shellcodeFormatted := strings.Join(formattedShellcode, ", ")

	return shellcodeFormatted
}
