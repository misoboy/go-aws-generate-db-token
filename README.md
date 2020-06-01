# misoboy/go-aws-generate-db-token

AWS Generate IAM DB Token in Go GUI.

## Installation

```
go get -u github.com/misoboy/go-aws-generate-db-token
```

## Examples

| [Demo](examples)|
| --- |
| ![Screenshot](examples/screenshot.png)|

## Overview

AWS IAM DB Token creation also has an AWS CLI, but it provides a GUI to help you create it easily.

In the PRD environment, it is assumed that MFA is activated and the OTP CODE is input.

We need to modify conf / env.conf before using it.

```conf
[_GLOBAL_SECTION_]

[dev-eu]
PROFILE=AWS CLI Profile Name
RDS_HOSTNAME=AWS RDS Service Endpoint
RDS_PORT=AWS RDS Service Port
RDS_USERNAME=AWS RDS Username
REDSHIFT_CLUSTER_ID=AWS Redshift Cluster ID
REDSHIFT_USERNAME=AWS Redshift Username
```

Enjoy!
