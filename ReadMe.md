# Barcode sequence generation for use with barcode font Barcode.ttf

## Usage
### Install font Barcode.ttf
### Use golang to generate sequences.
```go
import(
    "fmt"
    "github.com/dronm/fbarcode"
)

    barcode := "0722134016756"
    //this sequence should be placed in documents (excel, doc etc) with font set to Barcode
    fmt.Println(fbarcode.Ean13(barcode))

```
