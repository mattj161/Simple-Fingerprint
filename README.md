# Simple-Fingerprint
A very simple system fingerprinter in golang using WMI

Example

```
package main

import (
  "github.com/mattj161/Fingerprint"
  "fmt"
)

func main() {
	fp := Fingerprint.FingerPrint()
	fmt.Println(fp)
}
```
