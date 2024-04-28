# GoMake

Replace you setup scripts for this format.

# Use

Create a file with a name like this:

api.gomake

Then write the instructions like this:

```bash
$binary_name = "app_api"
$binary_name_win = "app_api.exe"
$build_dir = "../build"
$build_dir_linux = "../build/linux"
$build_dir_linuxarm64 = "../build/arm64"
$build_dir_win = "../build/windows"

// this is comment

setup{
    WORKDIR "./api"
    mkdir data
    go run ./appi
}

// this is comment
build{
    go build -o $binary_name
}

```

to execute run this:

```bash
go run ./gomake [SCRIPTNAME] [COMMAND]
```

```bash
go run ./gomake api setup
```