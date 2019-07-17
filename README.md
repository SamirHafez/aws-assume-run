# Motivation
`aws sts assume-role` is a pain because it requires parsing the output and passing the parsed credentials to the subsequent command.

`aws-assume-run` launches a subprocess with AWS IAM assumed role credentials. Think [envconsul](https://github.com/hashicorp/envconsul) for AWS IAM Roles.

# Install
```sh
go get -u github.com/fundingcircle/aws-assume-run
```

# Usage
```sh
aws-assume-run [role-arn-to-assume] [cmd] [params...]
```

# Examples

**Access s3 bucket via role**
```sh
aws-assume-run arn:aws:iam::121634321:role/bucket-reader aws s3 ls bucket-role-has-access-to
```

**Update a kubeconfig for a cluster in another account.**
```sh
aws-assume-run arn:aws:iam::121634321:role/eks-service-role aws eks update-kubeconfig --name dev-eks
```

# Building

Build for release.

```sh
make VERSION=v1.0.1 release -j2
```
