# Apple Identifier
A utility library used to help identify product names for Apple devices.

## Installation

In your Golang project, please run:

```
go get github.com/bartmika/appleidentifier
```

## Usage

```go
import (
    "fmt"

    "github.com/bartmika/appleidentifier"
)

deviceName := appleidentifier.GetProductNameByDeviceCode("iPhone1,1")
fmt.Println(deviceName) // "iPhone"
```

## License
Made with ❤️ by [Bartlomiej Mika](https://bartlomiejmika.com).   
The project is licensed under the [Unlicense](LICENSE).

Resource used:

* [adamawolf/Apple_mobile_device_types.txt](https://gist.github.com/adamawolf/3048717) is used for the data source that powers this library.
