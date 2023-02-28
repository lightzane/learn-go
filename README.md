# Learn Go

Learning GO language (https://go.dev/)

# Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Create GO Module](#create-go-module)
- [Hello World](#hello-go)
- [Notes on Packages and Imports](#notes-on-package-and-imports)
- [GO language Fundamentals](#next)

## Prerequisites

- vscode (https://code.visualstudio.com/)
- **Go** (vscode extension) by _Go Team at Google_
  - this extension may require another tools like `gopls` which helps in auto-suggestion for the vscode IDE and also helps identify compilation errors, unused imports and variables, etc.
    - **For Windows**, it will be saved under `%USERPROFILE%/go/bin` (e.g. **go/bin/gopls.exe**)

## Getting Started

Download and install GO (https://go.dev/dl/)

After installing, run in `cli` to confirm

```
go version
```

For Windows (will get the following result):

> go version go1.20.1 windows/amd64

## Create GO module

Create a folder for your GO project and execute the following command in `cli`

```
go mod init <module_name>
```

This will generate a `go.mod` file with following content:

```
module learn-go

go 1.20
```

## Hello GO

Create `main.go` file with the following content:

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello GO")
}
```

Run GO using the following command:

```
go run .
```

This will execute a `.go` file including all its other dependencies/imports

```
go run <filename>
```

This will only execute a single file but it will fail if the file is depending on other packages

`fmt` - is a package already exist in `Go` language.

## Notes on Package and Imports

- All `.go` file should start with `package <package_name>`

- Exporting all `variables` and `functions` should be written in _PascalCase_

```go
// shared/utils/my-package-name.go
package mypackagename

func PublicFunction() {}

func privateFunction() {}
```

```go
// main.go
package main

import (
    mypackagename "shared/utils"
)

func main() {
    mypackagename.PublicFunction()
}
```

# Next

Learn the fundamentals in GO language

- Go to: https://github.com/lightzane/learn-go/tree/booking-app#readme
