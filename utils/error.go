// https://github.com/SenzdettaCore/blsys

package utils

import (
    "os"
    "fmt"
    "github.com/SenzdettaCore/blsys/utils/color"
)

func Error(err error) {
    fmt.Fprintf(
        os.Stderr,
        "%s[!] %sError: %s%v%s\n",
        color.R, color.N, color.GG, err, color.N,
    )
    os.Exit(1)
}

// Copyright (c) 2026 Senzdetta