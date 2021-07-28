package base_encoding


import(
	"fmt"
	"strconv"
	"strings"
	"encoding/hex"
)

var EncodingValueBase64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
var EncodingValueBase32 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
var EncodingValueBase32Hex = "0123456789ABCDEFGHIJKLMNOPQRSTUV"

func Binary(s string) string {
    res := ""
    for _, c := range s {
        res = fmt.Sprintf("%s%.8b", res, c)
    }
    return res
}

func Padding(data string, multiple int, char string) string {
	dataLenght := len(data) % multiple
	if dataLenght % multiple != 0 {
		for i := 0; i < multiple - (dataLenght % multiple); i++ {
			data = data + char
		}
	}
	return data
}

func ChunkString(s string, chunkSize int) []string {
    var chunks []string
    runes := []rune(s)

    if len(runes) == 0 {
        return []string{s}
    }

    for i := 0; i < len(runes); i += chunkSize {
        nn := i + chunkSize
        if nn > len(runes) {
            nn = len(runes)
        }
        chunks = append(chunks, string(runes[i:nn]))
    }
    return chunks
}

func EncodeB64(plainText string) string {
	data := Padding(Binary(plainText), 6, "0")
	char := ChunkString(data, 6)
	var EncodedValueArray []int64
	for _, i := range char {
		a , _ := strconv.ParseInt(i, 2, 64)
		EncodedValueArray = append(EncodedValueArray, a)
	}
	var result string
	for _, u := range EncodedValueArray {
		result = result + string(EncodingValueBase64[u])
	}
	result = Padding(result, 4, "=")
	return result
}

func EncodeB64url(plainText string) string {
	b64 := EncodeB64(plainText)
	result := strings.ReplaceAll(b64, "/", "_")
	result = strings.ReplaceAll(b64, "+", "-")
	result = strings.ReplaceAll(b64, "=", "")
	return result
}

func EncodeB32(plainText string) string {
	data := Padding(Binary(plainText), 5, "0")
	char := ChunkString(data, 5)
	var EncodedValueArray []int64
	for _, i := range char {
		a , _ := strconv.ParseInt(i, 2, 64)
		EncodedValueArray = append(EncodedValueArray, a)
	}
	var result string
	for _, u := range EncodedValueArray {
		result = result + string(EncodingValueBase32[u])
	}
	result = Padding(result, 4, "=")
	return result
}

func EncodeB32hex(plainText string) string {
	data := Padding(Binary(plainText), 5, "0")
	char := ChunkString(data, 5)
	var EncodedValueArray []int64
	for _, i := range char {
		a , _ := strconv.ParseInt(i, 2, 64)
		EncodedValueArray = append(EncodedValueArray, a)
	}
	var result string
	for _, u := range EncodedValueArray {
		result = result + string(EncodingValueBase32Hex[u])
	}
	result = Padding(result, 4, "=")
	return result
}

func EncodeB16(plainText string) string {
	data := Padding(Binary(plainText), 4, "0")
	char := ChunkString(data, 4)
	var EncodedValueArray []int64
	for _, i := range char {
		a , _ := strconv.ParseInt(i, 2, 64)
		EncodedValueArray = append(EncodedValueArray, a)
	}
	var result string
	for _, u := range EncodedValueArray {
		result = result + string(EncodingValueBase32Hex[u])
	}
	return result
}

func GetChar(i string, charList string) (int, error) {
	for a, u := range charList {
		if i == string(u) {
			return a, nil
		}
	}
	return 0, fmt.Errorf("can't decode this string, invalid characters.")
}

func BasetoBinary(encodeString string, charList string, baseType string) (string, error) {
	var decodeB64Value string
	for _, i := range encodeString {
		char, err := GetChar(string(i), charList)
		if err != nil {
			return "", err
		}
		switch baseType {
			case "B64":
				decodeB64Value = decodeB64Value + fmt.Sprintf("%06b", char)
			case "B32":
				decodeB64Value = decodeB64Value + fmt.Sprintf("%05b", char)
			default:
				return "", fmt.Errorf("type %s is an invalid base type.", baseType)
		}
	}
	return decodeB64Value, nil
}

func BaseTostring(binData string) (string, error) {
	result := binData[:len(binData) - len(binData) % 8]
	var numbers []string
	for _, o := range ChunkString(result, 8) {
		add, err := strconv.ParseInt(o, 2, 64)
		if err != nil {
			return "", err
		}  
		numbers = append(numbers, string(rune(add)))
	}
	result = strings.Join(numbers[:], "")
	return result, nil
}

func DecodeB64(encodeString string) (string, error) {
	encodeString = strings.ReplaceAll(encodeString, "=", "")
	binData, err := BasetoBinary(encodeString, EncodingValueBase64, "B64")
	if err != nil {
		return "", err
	}
	result, err := BaseTostring(binData)
	if err != nil {
		return "", err
	}
	return result, nil
}

func DecodeB64url(encodeString string) (string, error) {
	encodeString = strings.ReplaceAll(encodeString, "/", "_")
	encodeString = strings.ReplaceAll(encodeString, "+", "-")
	result, err := DecodeB64(encodeString)
	if err != nil {
		return "", err
	}
	return result, nil
}

func DecodeB32(encodeString string) (string, error) {
	encodeString = strings.ReplaceAll(encodeString, "=", "")
	binData, err := BasetoBinary(encodeString, EncodingValueBase32, "B32")
	if err != nil {
		return "", err
	}
	result, err := BaseTostring(binData)
	if err != nil {
		return "", err
	}
	return result, nil
}

func DecodeB32hex(encodeString string) (string, error) {
	encodeString = strings.ReplaceAll(encodeString, "=", "")
	binData, err := BasetoBinary(encodeString, EncodingValueBase32Hex, "B32")
	if err != nil {
		return "", err
	}
	result, err := BaseTostring(binData)
	if err != nil {
		return "", err
	}
	return result, nil
}

func DecodeB16(encodeString string) (string, error) {
	decodeStringResult, err := hex.DecodeString(encodeString)
	if err != nil {
		return "", fmt.Errorf("invalid base 16 string")
	}
	return string(decodeStringResult), nil
}
