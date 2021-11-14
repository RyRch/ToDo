package main

import (
    . "fmt"
    . "os"
)

func helpOpt() {
    Printf("USAGE\n\ttodo [OPTIONS] task|ID\n\n")
    Printf("EXAMPLE\n\ttodo -a 'read a book'")
    Printf("\n\nOPTIONS\n\t-a, --add\t\t\tadd a new task\n\t")
    Printf("-b, --begin\t\t\tbegin a task\n\t-c, --check\t\t")
    Printf("\tcheck a task\n\t-d, --delete\t\tdelete a task\n\t-e, --edit\t")
    Printf("\t\tedit a task\n\t-h, --help\t\t\tprint this help\n\n")
}

func gestOpt()bool {
    if len(Args) == 1 {
        readFile()
        Exit(0)
    }
    if Args[1] == "--help" || Args[1] == "-h" {
        helpOpt()
        Exit(0)
    }
    return false
}