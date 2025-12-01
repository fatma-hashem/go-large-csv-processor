package main

import (
    "fmt"
    "github.com/fatma-hashem/go-large-csv-processor/internal"
)

func main() {
    internal.ProcessCSV("data/charts.csv", func(row []string) {
        fmt.Println(row)
    })
}
