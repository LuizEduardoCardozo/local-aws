# WIP - local-aws

A local service that emulates the AWS workspace.

## how to use

Launch the service by using

```
PORT=9000 make start
```

Use the usual aws commands, but changing the "aws" command for "laws"

```
after:
aws s3 mb s3://bucket-name
before:
laws s3 mb s3://bucket-name
```

## Roadmap

- SQS
- S3
- Cognito
