# Section 2: Getting Started

## Go project structure & Hello World Exmaple
* [Go Project Structure](project-structure)
* [Hello World Example Code](code)

## Code Organization
* Modules provide the ability to not have to strictly organize code.
* Organize code in a workspace. Code directory. GOPATH 
  * MAC OS: `/Users/bvandercar/go`

## Code Structure
* Application_structure
  |_ main.go
* main.go must have a "package main". 4


## Compiling (go build) & Running (go run) Go Code

### Run go Code
* You can run it before compiling you can use `go run main.go`
    * Also `go run -x main.go` gives you more verbose output.
    * This doesn't produce a binary.
* In vscode you `Rclick main.go >> open in integrated terminal >> go run main.go`
* `Ctrl + f5` in vscode
* To build go binaries without running the binary is to use `go build`.
    * To give the binary a different name by using `go build -O <name_of_binary>`
* To compile a binary for any operating system.
    * `GOOS=<windows/linux/darwin> GOARCH=<amd64/arm64/x86_64> go build -o <binary_name>`
    * Windows binary name would be `<binary_name>.exe`
* `go install` along with `go build` will place the resulting binary in the current directory and move the binary to `GOPATH/bin`.
* When running `go install` you use paths relative to `GOPATH/src`
* 




