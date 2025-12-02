package internal

import (
    "strconv" //package that converts strings to floats
)


func NormalizeRow(row []string, maxValue float64) []string { 
	//function takes a string slice from csv file and largest rank value in csv
	//returns []row after normalization
	//function modifies row in place (it is memory efficient)

	//if row has less than 2 columns then nothing to normalize bc row[1] does not exist
    if len(row) < 2 {
        return row
    }

    // convert 2nd column to float
    num, err := strconv.ParseFloat(row[1], 64)
    if err != nil {
        // if parsing fails then its probably an empty cell or not a number so skip the row
        return row
    }

    // normalize between 0 and 1
    normalized := num / maxValue

    // replace original value with normalized string
    row[1] = strconv.FormatFloat(normalized, 'f', 6, 64)

    return row
}
