This repo shows how to install tools used your project and track their version in your go.mod.

1. Clone this repository
2. Run make
3. Add the bin directory to your PATH, with 
    ```bash
    export PATH=`pwd`/bin:$PATH
    ```
4. Run `mockgen --version` or `golangci-lint` and confirm that they match the versions in our go.mod file.
