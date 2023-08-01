# CSV Compare

[![Build and Test](https://github.com/Climberdav/csv-compare/actions/workflows/go-test-and-build.yml/badge.svg)](https://github.com/Climberdav/csv-compare/actions/workflows/go-test-and-build.yml)


Package **csv-compare** helper provide an API to process CSV files with the same structure (columns).

The process results in a list of unique rows, in reverse order.

You can compare 1 (deduplication) or more files.

## Current capabilities
- deduplication (self-comparing) : top row has the precedence
- deduplication for n files
- not revert: on most case we want to have a diff file in FIFO style, this is the default behaviour,
    but sometimes you do not want.

## Installation
```bash
go get github.com/Climberdav/csv-compare
```
