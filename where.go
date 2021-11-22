// package where declares where the data from
package where

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "io/ioutil"
    "strings"
)

const (
    defa = "default.cfg"
)

func FromArgs() []string {
    var newargs []string
    for _, arg := range os.Args[1:] {
        if checkExist(arg) {
            processFile(arg)
        } else {
            newargs = append(newargs, arg)
        }
    }
    return newargs
}

func checkExist(filename string) bool {
    _, err := os.Stat(filename)
    if err == nil {
        return true
    }
    return false
}

func processFile(filename string) []string {
    var result []string
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalf("File read failed!")
    }

    for _, line := range strings.Split(string(data), "\n") {
        result = append(result, line)
    }
    return result
}

func FromFile(filename ...string) {
    if len(filename) == 0 {
        searchLocal(defa)
    } else {
        files := os.Args[1:]
        if len(files) == 1 {
            searchLocal(files[0])
        } else if len(files) > 1 {
            for _, file := range files {
                f, err := os.Open(file)
                defer f.Close()

                // TODO: replace with noerr
                if err != nil {
                    fmt.Fprintf(os.Stderr, "failed: %v\n", err)
                    continue
                }
            }
        }
    }
}

func searchLocal(defa string) []string {
    var result []string
    if checkExist(defa) {
        result = processFile(defa)
    } else {
        result = append(result, fmt.Sprintf("No such file named default."))
    }
    return result
}

func FromStdin() string {
    var result string
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        result += input.Text()
    }
    return result
}
