Todo CLI
============


![prompt](https://raw.githubusercontent.com/HxX2/todocli/main/.github/assets/todo.png)


**Table of Contents**

<!-- toc -->

- [About](#about)
  * [Installing](#installing)
  * [Uninstalling](#uninstalling)
  * [Build From Source](#build-from-source)
  * [Usage](#usage)

<!-- tocstop -->

## About

Todo CLI is a simple to do list to manage your tasks.
Written in GO and styled with [Nerd Fonts](https://www.nerdfonts.com/)

### Installing

```console
GOBIN=<your install path> go install ./cmd/todo
```

### Uninstalling

```console
rm -rf <your install path>/todo
```

### Build From Source

Install Go and build with this command:

```console
go build ./cmd/todo 
```

### Usage

To add a task to the list

```console
todo -a <Task String>
```
Toggle a task as done or undone

```console
todo -t <Task number>
```
Remove a Task from the list

```console
todo -r <Task Number>
```
Opens editor to edite the raw file of the list (it uses the $EDITOR env var)

```console
todo -e 
```

List done tasks

```console
todo -ld
```

List undone tasks

```console
todo -lu
```

Hide Progress bar (can be used with other options)

```console
todo -hp
```
