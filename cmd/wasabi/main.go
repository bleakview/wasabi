package main

import (
	"context"
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/sys"
	"github.com/tetratelabs/wazero/wasi_snapshot_preview1"
)

func main() {
	// Choose the context to use for function calls.
	ctx := context.Background()
	// Create a new WebAssembly Runtime.
	var r wazero.Runtime
	// Get Architecture and OS to check if AoT is supported
	goArch := strings.ToLower(runtime.GOARCH)
	goOs := strings.ToLower(runtime.GOOS)
	//If AoT is supported activate it else run in interpreted mode
	if ((goArch == "amd64") || (goArch == "arm64")) && ((goOs == "windows") || (goOs == "linux")) {
		r = wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfigCompiler().
			WithWasmCore2().WithFeatureBulkMemoryOperations(true))
	} else {
		r = wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfigInterpreter().
			WithWasmCore2().WithFeatureBulkMemoryOperations(true))
	}

	defer r.Close(ctx) // This closes everything this Runtime created.

	// Combine the above into our baseline config, overriding defaults.
	config := wazero.NewModuleConfig().
		WithStdout(os.Stdout).
		WithStderr(os.Stderr).
		WithSysWalltime(). //we need these two to have the same time as the host
		WithSysNanotime().
		WithArgs()

	if _, err := wasi_snapshot_preview1.Instantiate(ctx, r); err != nil {
		log.Panicln(err)
	}

	//Get executable wasi as the commandline argument
	arg := os.Args[1]
	if len(os.Args) == 1 {
		print("No commandline variable")
	}
	if arg == "" {
		print("No commandline variable")
	}

	// read wasm file
	var catWasm []byte
	catWasm, err := ioutil.ReadFile(arg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Compile the WebAssembly module using the default configuration.
	code, err := r.CompileModule(ctx, catWasm, wazero.NewCompileConfig())
	if err != nil {
		log.Panicln(err)
	}

	// InstantiateModule runs the "_start" function, WASI's "main".
	if _, err = r.InstantiateModule(ctx, code, config.WithArgs("wasi")); err != nil {

		// Note: Most compilers do not exit the module after running "_start",
		// unless there was an error. This allows you to call exported functions.
		if exitErr, ok := err.(*sys.ExitError); ok && exitErr.ExitCode() != 0 {
			fmt.Fprintf(os.Stderr, "exit_code: %d\n", exitErr.ExitCode())
		} else if !ok {
			log.Panicln(err)
		}
	}
}
