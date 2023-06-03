gengrpc - Generate gRPC
=======================

Copyright (c) 2022, 2023, Geert JM Vanderkelen

gengrpc is a package and command line tool to generate gRPC code from
Protocol Buffer contract files. It is based on the tool found in the
project [pxmysql][1] but focusing on getting the contracts through Go packages.
The idea is to eventually use `genrpc` in `pxmysql`, merging both
functionalities.

Quick start
-----------

This project wraps around the [`protoc`][2] tool. You will need to
install this first. Actually, you can use `protoc` directly, but it can become
painful with lots of contract files (hence this package).

You have:
* contract files (*.proto) in a shared Go package, for example, 
  `example.com/golistic/shared` and files are located in that repository under
  `grpc`
* you have another project `example.com/golistic/eventing` which needs to
  generate using the contract files found in the `shared` project
* you want the generate code in the folder `internal/grpc`

Using the above example, you can the execute the following:

```shell
$ gorpc --module example.com/golistic/eventing \
 --source example.com/golistic/shared \
 --source-path grpc \
 --dest-path internal/grpc
```

The `--source-path` flag is optional. If not provided, all `.proto` files will
be included.

The source-package must be loaded through Go modules. It works using Go
workspaces.

License
-------

Distributed under the MIT license. See `LICENSE.md` for more information.

[1]: https://github.com/golistica/pxmysql

[2]: https://grpc.io/docs/languages/go/quickstart/
