# GoSorter
[![Go](https://github.com/avazquezcode/govetryx/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/avazquezcode/govetryx/actions/workflows/ci.yml)
<a href="https://goreportcard.com/report/github.com/avazquezcode/gosorter"><img src="https://goreportcard.com/badge/github.com/avazquezcode/gosorter" alt="Go Report Card" /></a>
[![codecov](https://codecov.io/gh/avazquezcode/gosorter/graph/badge.svg?token=GCULNO7W0Q)](https://codecov.io/gh/avazquezcode/gosorter)

GoSorter is a golang implementation of a command line tool to sort lines in files.
It is largely inspired in the ["sort"](https://man7.org/linux/man-pages/man1/sort.1.html) linux command, but implemented in Golang. Also the functionality is not exactly the same (the **linux** tool is much more extensive and supports more options).

*Important*: This tool doesn't support binary files at the moment.

# Important

⚠️ This tool was developed very quickly and just for fun :). It is **NOT** battle tested. So please use it carefully and at your own risk!

# Getting started

- Just install this tool: `go install github.com/avazquezcode/gosorter@latest`

- Execute: `gosorter sort <FILE_PATH>`

- To check the different flags (options) you can use, please run `gosorter sort --help`

# Documentation
For a complete documentation of the tool, please refer to [this](./docs/gosorter.md).

