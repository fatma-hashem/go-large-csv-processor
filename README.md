# Go Large CSV Processor

A memory-efficient CSV processor written in Go. Designed to handle very large CSV files (hundreds of MBs to gigabytes) without loading the entire file into memory.

## Overview

This project demonstrates:
- **Streaming CSV processing** (constant memory usage)
- **Two-pass processing** (find max value → normalize rows)
- **On-the-fly row transformation**
- **Writing output to a new CSV file**
- **Basic memory profiling** using `pprof`

---

## Features

### ✅ 1. Stream Processing (O(1) Memory)

The CSV is processed row by row using a callback mechanism:

```go
internal.ProcessCSV(path, callback)
```

This prevents Go from loading the entire dataset into memory.

### ✅ 2. Max-Value Scan

The program first scans the dataset to compute the maximum value in column 2 (index 1), which is required for normalization.

### ✅ 3. Normalization

A second pass normalizes each row using:

```
normalized_value = original_value / max
```

The normalized data is written directly to:

```
data/processed_normalized.csv
```

### ✅ 4. Memory Profiling

After processing, the program writes a heap profile:

```
memprofile.prof
```

Analyze it with:

```bash
go tool pprof memprofile.prof
```

---

## Project Structure

```
go-large-csv-processor/
│
├── cmd/
│   └── main.go                # Main entry point
│
├── internal/
│   ├── process.go             # Stream CSV reader + callback
│   └── normalize.go           # Normalization logic
│
├── data/
│   ├── charts.csv             # Input CSV (original)
│   ├── processed_normalized.csv
│   └── (large files ignored via .gitignore)
│
└── README.md                  # Documentation
```

---

## How It Works

### Step 1 — Find Max Value
The program makes a first pass to calculate the maximum numeric value in column 2.

### Step 2 — Normalize Each Row
Each row is divided by the max value and written directly to a new CSV.

### Step 3 — Memory Profile
A heap snapshot is generated so you can inspect memory usage during processing.

---

## Running the Program

### 1. Place your CSV

Place your input file at:

```
data/charts.csv
```

### 2. Run the processor

```bash
go run cmd/main.go
```

**Example output:**

```
Max value in column 1: 200
CSV normalized and written to processed_normalized.csv
Memory profile saved to memprofile.prof
```

---

## Memory Profiling

Generate heap report:

```bash
go tool pprof memprofile.prof
```

View top allocations:

```
(pprof) top
```

Graph view (requires Graphviz):

```
(pprof) web
```

---

## Why This Project Matters

This project demonstrates **production-grade large file handling**:

- ✅ No exponential RAM growth
- ✅ Efficient disk streaming
- ✅ Suitable for large datasets
- ✅ Easy to extend for ETL or database ingestion

It showcases Go's strengths in:

- ✔️ Streaming I/O
- ✔️ Low memory footprint
- ✔️ High-performance data pipelines

---

## License

MIT

---

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.
