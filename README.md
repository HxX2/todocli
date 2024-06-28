Todo CLI
============


![prompt](https://raw.githubusercontent.com/HxX2/todocli/main/images/todo.png)


**Table of Contents**

<!-- toc -->

- [About](#about)
  * [Installing](#installing)
  * [Build From Source](#post-installation)
  * [Usage](#uninstalling)

<!-- tocstop -->

## About

### Installing

Move the binary in the bin dir to your desired binary's dir

### Build From Source

Install Go and build with this command:

```console
go build
```

### Usage

To add a task to the list

```console
todo -a <>
```
To add a task to the list

```console
todo -t <Number of task>
```
To add a task to the list

```console
todo -r <Number of task>
```
opens editor to edite the raw file of the list (it uses the $EDITOR env var)

```console
todo -e 
```
