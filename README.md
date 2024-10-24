# GoSorter
[![Go](https://github.com/avazquezcode/gosorter/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/avazquezcode/gosorter/actions/workflows/ci.yml)
<a href="https://goreportcard.com/report/github.com/avazquezcode/gosorter"><img src="https://goreportcard.com/badge/github.com/avazquezcode/gosorter" alt="Go Report Card" /></a>
[![codecov](https://codecov.io/gh/avazquezcode/gosorter/graph/badge.svg?token=GCULNO7W0Q)](https://codecov.io/gh/avazquezcode/gosorter)

GoSorter is a command line tool written in Golang, to sort **lines** of files.

It is largely inspired by the ["sort"](https://man7.org/linux/man-pages/man1/sort.1.html) linux command. With that being said, the functionality is not exactly the same (the **linux** tool is much more extensive and supports more options at the moment).

# Important

⚠️ This tool was developed very quickly and just for fun :). It is **NOT** battle tested. So please use it carefully and at your own risk!

# Getting started

## Installation

- Just install the tool: `go install github.com/avazquezcode/gosorter@latest`

## Usage

*Basic usage, please run*:

`gosorter sort <FILE_PATH>`

*For more than one file just add multiple names :). The lines of all the files will be merged together, and then sorted. Example*:

`gosorter sort <FILE1_PATH> <FILE2_PATH> <FILE3_PATH>`

*To check the different flags (options), please run*: 

`gosorter sort --help`

# Documentation

For a complete documentation of the tool, please refer to [this](./docs/gosorter.md).
