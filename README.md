# go-aws-generate-db-token

> A Go GUI desktop tool for generating **AWS IAM database authentication tokens** for RDS and Redshift — with MFA/OTP support.

[![Go](https://img.shields.io/badge/Go-1.x-00ADD8?logo=go&logoColor=white)](https://golang.org)
[![AWS](https://img.shields.io/badge/AWS-RDS%20%7C%20Redshift-FF9900?logo=amazonaws&logoColor=white)](https://aws.amazon.com)
[![GUI](https://img.shields.io/badge/UI-Desktop_GUI-4A90E2)](https://github.com/misoboy/go-aws-generate-db-token)
[![Windows](https://img.shields.io/badge/Windows-exe-0078D6?logo=windows&logoColor=white)](https://github.com/misoboy/go-aws-generate-db-token)

## Overview

AWS CLI can generate IAM DB tokens but it requires command-line usage which is cumbersome for production access. This tool provides a **native GUI** to:

- Input AWS credentials and region
- Enter MFA OTP code (for MFA-enabled environments)
- Generate an IAM DB authentication token for RDS or Redshift
- Copy the token for use in database clients

| Screenshot |
|-----------|
| <img src="examples/screenshot.png" width="400"> |

## Prerequisites

- Go 1.16+
- Windows (primary target) or cross-platform build

## Installation

```bash
go get -u github.com/misoboy/go-aws-generate-db-token
```

## Build

```bash
# Windows
set GOROOT=<Go installation path>
set GOPATH=<Project workspace path>
go build -o aws-generate-db-token.exe github.com/misoboy/go-aws-generate-db-token

# Linux/macOS
go build -o aws-generate-db-token github.com/misoboy/go-aws-generate-db-token
```

## Project Structure

```
.
├── main.go       # Entry point and GUI setup
├── rds/          # RDS token generation
├── redshift/     # Redshift token generation
├── common/       # Shared AWS utilities
├── conf/         # Configuration
└── examples/     # Screenshots and usage examples
```

## License

MIT
