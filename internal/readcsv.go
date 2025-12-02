package internal

import (
    "bufio"       // for reading large files efficiently
    "encoding/csv" // parses CSV rows into slices of strings
    "io"
    "log"
    "os"
)

// ProcessCSV opens a CSV file and calls a callback function for each row.
// This allows processing rows one by one without loading the entire file into memory.
func ProcessCSV(filePath string, callback func([]string)) {
    // Open the CSV file
    file, err := os.Open(filePath) // returns the file and an error
    if err != nil {                // if there is an error, log it and stop execution
        log.Fatal(err)
    }
    defer file.Close() // ensure the file is closed later, no matter what

    // Create a CSV reader with buffering for efficiency
    reader := csv.NewReader(bufio.NewReader(file))
    // bufio divides the file into chunks
    //parsing is converting comma separated values to string go slice like ["alice", "25"] instead of "alice,25" is this better
    for {
        // Read one CSV row at a time
        record, err := reader.Read()
        if err == io.EOF { // end of file reached
            break
        }
        if err != nil { // if an error occurs while reading, log it and skip that row
            log.Println("Error reading row:", err)
            continue
        }

        // Call the callback function for processing the row
        // This is where you can normalize the data, print it, or write it to another file
        callback(record)
    }
}
