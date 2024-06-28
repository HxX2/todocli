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

Todo CLI is a simple todo list to manage your tasks writen in GO and styled with Nerd Fonts

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
