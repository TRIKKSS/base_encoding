# BASE ENCODING

**DESCRIPTION**

a simple golang package written to understand how base encodings works

**USAGE**
```bash
go get github.com/TRIKKSS/base_encoding
```

```golang
package main

import(
	"fmt"
	base "github.com/TRIKKSS/base_encoding"
)


func main() {
	//encode
	fmt.Println(base.EncodeB64("mystring"))

	fmt.Println(base.EncodeB64url("mystring"))

	fmt.Println(base.EncodeB32("mystring"))

	fmt.Println(base.EncodeB32hex("mystring"))

	fmt.Println(base.EncodeB16("mystring"))

	//decode
	decodeStringB64, err := base.DecodeB64("bXlzdHJpbmc=")
	checkErr(err)
	decodeStringB64url, err := base.DecodeB64url("bXlzdHJpbmc")
	checkErr(err)
	decodeStringB32, err := base.DecodeB32("NV4XG5DSNFXGO===")
	checkErr(err)
	decodeStringB32hex, err := base.DecodeB32hex("DLSN6T3ID5N6E===")
	checkErr(err)
	decodeStringB16, err := base.DecodeB16("6D79737472696E67")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
```