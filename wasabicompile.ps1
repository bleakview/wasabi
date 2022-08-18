$Env:GOOS = "windows"; $Env:GOARCH = "amd64"
mkdir bin\windows -ErrorAction SilentlyContinue
go build -o bin\windows\wasabi.exe .\cmd\wasabi\main.go
$Env:GOOS = "freebsd"; $Env:GOARCH = "amd64"
mkdir bin\freebsd -ErrorAction SilentlyContinue
go build -o bin\freebsd\wasabi .\cmd\wasabi\main.go
$Env:GOOS = "linux"; $Env:GOARCH = "ppc64le"
mkdir bin\linuxppc64le -ErrorAction SilentlyContinue
go build -o bin\linuxppc64le\wasabi .\cmd\wasabi\main.go
$Env:GOOS = "linux"; $Env:GOARCH = "riscv64"
mkdir bin\linuxrv64 -ErrorAction SilentlyContinue
go build -o bin\linuxrv64\wasabi .\cmd\wasabi\main.go