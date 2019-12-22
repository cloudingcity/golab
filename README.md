# Golab

[![Build Status](https://travis-ci.com/cloudingcity/golab.svg?branch=master)](https://travis-ci.com/cloudingcity/golab)
[![codecov](https://codecov.io/gh/cloudingcity/golab/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudingcity/golab)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudingcity/golab)](https://goreportcard.com/report/github.com/cloudingcity/golab)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](http://godoc.org/github.com/cloudingcity/golab)

`golab` is a command line tool that make working with GitLab easier.

## Features

- Easy to configure gitlab settings
- Show all merge requests that created by you or assigned to you
- Open page in default browser
- Lint `.gtilab-ci.yml`

## Installation

### Homebrew

```shell script
brew install cloudingcity/tap/golab
```

### Source

```shell script
git clone git@github.com:cloudingcity/golab.git
cd golab
make install
```

## Command Usage

```
A CLI tool for gitlab

Usage:
  golab [command]

Available Commands:
  ci          Manage gitlab ci
  config      Init or list golab CLI options
  mr          Manage merge requests
  own         Manage own resources
  version     Print version number of golab

Flags:
  -h, --help   help for golab

Use "golab [command] --help" for more information about a command.
```

### Initial config

```shell script
golab config init                                                                                                    master ↓ 1 ↑ 1 ✚ 1 
```

### Show current repository merge requests
```shell script
golab mr list
```

### Show all merge requests that assigned to you
```shell script
golab own mr list --review
```

### Open merge requests page in browser
```shell script
golab mr open <MR-ID>
```
