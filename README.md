[![golangci-lint](https://github.com/bleakview/wasabi/actions/workflows/golanglint.yml/badge.svg)](https://github.com/bleakview/wasabi/actions/workflows/golanglint.yml) [![release build](https://github.com/bleakview/wasabi/actions/workflows/releasebuild.yml/badge.svg)](https://github.com/bleakview/wasabi/actions/workflows/releasebuild.yml)
# wasabi
A multi platform runtime for WASI. It's written to show that its possible to run WebAssembly in multiple platforms. When .Net 7 preview began supported WASI as a target , I want to know if multiple targets which are not supported by .Net can execute files. Answer yes 🙂

Tested with .Net 7 preview and AssemblyScript. Based on [Wazero](https://github.com/tetratelabs/wazero/).

How to use
After downloading from Release or compiling yourself you can run WASI  files by
```
wasabi executable.wasm
```
I write a article on it on [https://www.wasm.builders/bleakview/running-net-where-it-has-never-run-before-risc-v-freebsd-and-beyond-4cd](https://www.wasm.builders/bleakview/running-net-where-it-has-never-run-before-risc-v-freebsd-and-beyond-4cd)

<figure>
<img src="https://bleakview.github.io/git/wasabi/images/linux_riscv64.jpg" alt="example linux with risc-v" width="800"/>
<figcaption>Running WASI on Linux RISC-V 64</figcaption>
</figure>
<figure>
<img src="https://bleakview.github.io/git/wasabi/images/linux_ppc64le.jpg" alt="example linux with powerpc64le" width="800"/>
<figcaption>Running WASI on Linux PowerPC64le</figcaption>
</figure>
<figure>
<img src="https://bleakview.github.io/git/wasabi/images/freebsd_amd64.jpg" alt="example freebsd with amd64" width="800"/>
    <figcaption>Running WASI on FreeBSD Amd64</figcaption>
</figure>

