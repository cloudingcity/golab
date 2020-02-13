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
- Shows project which depend on a certain package (Golang: `go.mod`, PHP: `composer.json`)
- Lint `.gtilab-ci.yml`

## Installation

### Homebrew

```shell script
brew install cloudingcity/tap/golab
```

### Binary

Download the pre-built binaries from the [Releases](https://github.com/cloudingcity/golab/releases) page. Extract them, move it to your `$PATH`.

```shell script
curl -OL https://github.com/cloudingcity/golab/releases/download/v0.6.1/golab_0.6.1_Linux_x86_64.tar.gz
tar -xzvf golab_0.6.1_Linux_x86_64.tar.gz
mv golab /usr/local/bin/golab
golab version
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
  clone       Clone a repository from GitLab
  config      Init or list golab CLI options
  depend      Shows project which depend on a certain package
  mr          Manage merge requests
  own         Manage own resources
  project     Manage projects
  version     Print version number of golab

Flags:
  -h, --help   help for golab

Use "golab [command] --help" for more information about a command.
```

### Initial config

```shell script
$ golab config init                                                                                                    master ↓ 1 ↑ 1 ✚ 1 
Gitlab Host [https://gitlab.com]: <INPUT>
Create a token here: https://gitlab.com/profile/personal_access_tokens
Gitlab Token (scope: api) [None]: <INPUT>

Config saved to /Users/<USER>/.config/golab.yaml
```

### Clone repository from GitLab
```shell script
$ golab clone pokemon/eevee
# git clone git@gitlab.com:pokemon/eevee.git
```

### Show current repository merge requests
```shell script
$ golab mr list
  MRID   TITLE                        URL                                                                   
  1      Catch your first Pokémon     https://example.com/pokemon/trainer/merge_requests/1  
  2      To become a Pokémon Master   https://example.com/pokemon/trainer/merge_requests/2  
```

### Show all merge requests that assigned to you
```shell script
$ golab own mr list --review
  PID    MRID   PROJECT           TITLE                        URL                                                                   
  4255   1      pokemon/trainer   Catch your first Pokémon     https://example.com/pokemon/trainer/merge_requests/1  
  4255   2      pokemon/trainer   To become a Pokémon Master   https://example.com/pokemon/trainer/merge_requests/2  
```

### Open merge requests page in browser
```shell script
$ golab mr open <MR-ID>
```

### Show `pokemon/eevee` [composer](https://getcomposer.org/) package which project depend on
```shell script
$ golab depend php pokemon/eevee --group pokemon
  PROJECT    VERSION   BRANCH    URL
  vaporeon   v0.1.2    master    https://example.com/pokemon/vaporeon
  jolteon    v1.2.0    staging   https://example.com/pokemon/jolteon
  flareon    v3.0.0    staging   https://example.com/pokemon/flareon
```

### Show `example.com/pokemon/eevee` [go modules](https://github.com/golang/go/wiki/Modules) which project depend on
```shell script
$ golab depend go example.com/pokemon/eevee --group pokemon
  PROJECT    VERSION   BRANCH    URL
  vaporeon   v0.1.2    master    https://example.com/pokemon/vaporeon
  jolteon    v1.2.0    staging   https://example.com/pokemon/jolteon
  flareon    v3.0.0    staging   https://example.com/pokemon/flareon
```

### Check `.gitlab-ci.yml` is valid

```shell script
$ golab ci lint .gitlab-ci.yml
Valid!
```

## Development

### Running gitlab on a container

```shell script
docker-compose up -d
```
