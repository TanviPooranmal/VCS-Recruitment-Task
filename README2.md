# CLI Tool Documentation

## Introduction

This CLI tool provides functionalities for creating backups, sharing directories, and configuring backup settings. It offers a simple and efficient way to manage backups of your important files and directories.

## Binary Installation

This method allows you to install the CLI tool directly from pre-compiled binary executables. Users can download the binaries and install them onto their system without the need for compiling from source code.

### Installation

Before using the tool, ensure you have Go installed on your system. If not, you can download and install it from [here](https://golang.org/dl/).

To install the CLI tool, run the following command:

```bash
go install github.com/TanviPooranmal/VCS-Recruitment-Task
```

## Source Installation

This method involves cloning the repository containing the CLI tool's source code and manually building it on your local machine. Users have access to the source code and can customize or modify it according to their needs.

Following these steps will help you successfully clone the repository, navigate to the directory, and build the CLI tool from source.

### Clone the Repository
Start by cloning the repository that contains the CLI tool to your local machine. You can do this by executing the following command in your terminal:

```bash
git clone https://github.com/TanviPooranmal/VCS-Recruitment-Task 
```

This command will download the repository files to your local machine.

### Navigate to the Directory

Once the repository is cloned, navigate to the directory of the cloned repository. You can do this by executing the following command, replacing `<repository_name>` with the name of the cloned repository:

```bash
cd VCS-Recruitment-Task
```

This command changes your current directory to the directory where the repository is cloned.

### Build the CLI Tool

If the repository contains build instructions (e.g., Makefile or build scripts), follow those instructions to build the CLI tool. Typically, you can build the tool using the `go build` command for Go projects. Execute the following command in the terminal:

```bash
go build
```

This command compiles the source code and generates an executable file for the CLI tool.

## Usage

The CLI tool supports the following commands:

### 1. Backup

Create backups of directories.

```bash
Code backup [flags]
```

#### Flags

- `-src`: Source directory to backup.
- `-encrypt`: Encrypt files during backup.
- `-recursive`: Recursively encrypt files.
- `-selective`: Selectively encrypt files (comma-separated).
- `-share`: Share the backup directory.

Example:

```bash
Code backup -src /path/to/source -encrypt -recursive -share
```

### 2. Share

Share directories or files.

```bash
Code share [flags]
```

#### Flags

- `-dir`: Directory to share.
- `-files`: Files to send (comma-separated).
- `-prev-versions`: Share previous backup versions.

Example:

```bash
Code share -dir /path/to/directory -prev-versions
```

### 3. Config

View or modify configuration settings.

```bash
Code config [flags]
```

#### Flags

- `-root-dir`: Root directory for the backup.
- `-logger-format`: Format of the logger file.

Example:

```bash
Code config -root-dir /new/root/directory -logger-format "2006-01-02T15:04:05"
```

## Examples

### Backup Example

To create a backup of a directory `/path/to/source`, encrypt files recursively, and share the backup directory, run:

```bash
Code backup -src /path/to/source -encrypt -recursive -share
```

### Share Example

To share a directory `/path/to/directory` along with previous backup versions, run:

```bash
Code share -dir /path/to/directory -prev-versions
```

### Config Example

To view or modify configuration settings, run:

```bash
Code config -root-dir /new/root/directory -logger-format "2006-01-02T15:04:05"
```

## Conclusion

With this CLI tool, you can easily manage your backups, share directories or files, and configure backup settings according to your requirements.
