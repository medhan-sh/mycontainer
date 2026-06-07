# MyContainer 🐳

A minimalist Linux container runtime written from scratch in Go following LIZ Rice's video on "Building a container from Scratch" and building upon that.

This project explores the underlying mechanics of containerization engines (like Docker) by directly interfacing with Linux system calls. It creates isolated environments using Linux Namespaces, `chroot`, and custom filesystem mounts without relying on external container libraries.

## Features :

* **Process Isolation:** Utilizes `CLONE_NEWPID` to isolate the process ID number space. The containerized process runs as PID 1 inside its environment.
* **Network & Hostname Isolation:** Uses `CLONE_NEWUTS` to allow the container to have its own independent hostname (`inside`).
* **Mount Namespace Isolation:** Leverages `CLONE_NEWNS` to ensure mount and unmount operations do not affect the host system.
* **Filesystem Jailing:** Implements `chroot` to restrict the container's file system access to an isolated `/container` directory.
* **Procfs Mounting:** Dynamically mounts and unmounts the pseudo-filesystem `/proc` so utilities like `ps` work correctly inside the container.

## Prerequisites

* **Linux Operating System:** This project relies on Linux-specific system calls (Namespaces). It will not run natively on macOS or Windows.
* **Go 1.13+** installed.
* **Root Privileges:** Required to execute system calls like `chroot` and `mount`.
* A target directory named `/container` at the root of your host filesystem containing a basic rootfs (or dynamically adjust the path in the code).

## Usage

Build the Go binary:
```bash
go build -o mycontainer main.go
