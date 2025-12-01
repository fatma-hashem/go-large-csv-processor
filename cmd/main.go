package main

import (
    "fmt"
    "github.com/fatma-hashem/go-large-csv-processor/internal"
    "log"
    "os"
    "encoding/csv"
    "strconv"
)

func main() {
    // find max value of rank which is second column (needed for normalization)
    var max float64 //initilizes max value to 0
    internal.ProcessCSV("data/charts.csv", func(row []string) { //calls callback for each row
        if len(row) < 2 { //so no error when accessing row[1]
            return
        }
        v, err := strconv.ParseFloat(row[1], 64)
        if err != nil { //could be empty cell or not a number so return
            return
        }
        if v > max {
            max = v
        }
    })

    fmt.Println("Max value in column 1:", max)

    // Now process and normalize row by row
    //open new file and overwrite if it exists
    outFile, err := os.Create("data/processed_normalized.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer outFile.Close() //ensure file closes at the end even if there is an error

    writer := csv.NewWriter(outFile) //allows writing rows as csv lines
    defer writer.Flush()

    internal.ProcessCSV("data/charts.csv", func(row []string) {
        normalizedRow := internal.NormalizeRow(row, max)
        writer.Write(normalizedRow) // write normalized row directly
    })

    fmt.Println("CSV normalized and written to processed_normalized.csv")
}
