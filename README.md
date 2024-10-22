# GoSorter
[![Go](https://github.com/avazquezcode/govetryx/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/avazquezcode/govetryx/actions/workflows/ci.yml)
<a href="https://goreportcard.com/report/github.com/avazquezcode/gosorter"><img src="https://goreportcard.com/badge/github.com/avazquezcode/gosorter" alt="Go Report Card" /></a>
[![codecov](https://codecov.io/gh/avazquezcode/gosorter/graph/badge.svg?token=GCULNO7W0Q)](https://codecov.io/gh/avazquezcode/gosorter)

GoSorter is a command line tool written in Golang, to sort **lines** of files.

It is largely inspired by the ["sort"](https://man7.org/linux/man-pages/man1/sort.1.html) linux command, but implemented in Golang. Also the functionality is not exactly the same (the **linux** tool is much more extensive and supports more options at the moment).

# Important

⚠️ This tool was developed very quickly and just for fun :). It is **NOT** battle tested. So please use it carefully and at your own risk!

# Getting started

## Installation

- Just install the tool: `go install github.com/avazquezcode/gosorter@latest`

## Usage

__Basic usage__:

- Execute: `gosorter sort <FILE_PATH>`

__For more than one file just add multiple names :) For example__:

- Execute: `gosorter sort <FILE1_PATH> <FILE2_PATH> <FILE3_PATH>`

__To check the different flags (options) you can use, please run__: 

- Execute: `gosorter sort --help`

# Documentation

For a complete documentation of the tool, please refer to [this](./docs/gosorter.md).
