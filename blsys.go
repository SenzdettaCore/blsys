// https://github.com/SenzdettaCore/blsys

package main

import (
    "os"
    "github.com/SenzdettaCore/blsys/utils/invinput"
    "github.com/SenzdettaCore/blsys/module/help"
    "github.com/SenzdettaCore/blsys/module/copy"
    "github.com/SenzdettaCore/blsys/module/move"
    "github.com/SenzdettaCore/blsys/module/remove"
    "github.com/SenzdettaCore/blsys/module/create"
    "github.com/SenzdettaCore/blsys/module/fsclear"
    "github.com/SenzdettaCore/blsys/module/chmod"
    "github.com/SenzdettaCore/blsys/module/symlink"
    "github.com/SenzdettaCore/blsys/module/lookpath"
    "github.com/SenzdettaCore/blsys/module/codegen"
    "github.com/SenzdettaCore/blsys/module/tmpdir"
    "github.com/SenzdettaCore/blsys/module/bindir"
    "github.com/SenzdettaCore/blsys/module/platform"
    "github.com/SenzdettaCore/blsys/module/geteuid"
    "github.com/SenzdettaCore/blsys/module/manifest"
    "github.com/SenzdettaCore/blsys/module/find"
    "github.com/SenzdettaCore/blsys/module/runcmd"
    "github.com/SenzdettaCore/blsys/module/optdir"
    "github.com/SenzdettaCore/blsys/module/install"
)

type CommandFunc func(
    args []string,
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    commands := map[string]CommandFunc{
        "--help": help.Execute,
        "--copy": copy.Execute,
        "--move": move.Execute,
        "--remove": remove.Execute,
        "--create": create.Execute,
        "--fsclear": fsclear.Execute,
        "--chmod": chmod.Execute,
        "--symlink": symlink.Execute,
        "--lookpath": lookpath.Execute,
        "--codegen": codegen.Execute,
        "--tmpdir": tmpdir.Execute,
        "--bindir": bindir.Execute,
        "--platform": platform.Execute,
        "--geteuid": geteuid.Execute,
        "--manifest": manifest.Execute,
        "--find": find.Execute,
        "--runcmd": runcmd.Execute,
        "--optdir": optdir.Execute,
        "--install": install.Execute,
    }

    cmdKey := args[0]
    subArgs := args[1:]

    if fn, exists := commands[cmdKey]; exists {
        fn(subArgs)
    } else {
        invinput.InvalidOption(cmdKey)
        os.Exit(1)
    }
}

// Copyright (c) 2026 Senzdetta