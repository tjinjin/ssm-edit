# ssm-edit

Edit a parameter  on Systems Manager parameter store.

Inspiread by https://github.com/tsub/s3-edit


# Install

Support homebrew.

```
$ brew install tjinjin/ssm-edit/ssm-edit
```
# AWS credentials

Support `~/.aws/credential`.

```
cat ~/.aws/credentials
[default]
aws_access_key_id = hogehoge
aws_secret_access_key = fugafuga

[test]
role_arn = arn:aws:iam::123456789012:root
source_profile = default
```

# Usage

```
$ ssm-edit list
$ ssm-edit edit
$ ssm-edit version
```


# Todo
- [x] Support assumeRole
- [x] Support Pagination to list command
- [ ] Add `create` command
- [ ] Add `show` command
- [ ] Add Makefile
- [ ] Add test
- [ ] Add CI
- [ ] Refactoring!!!
- [ ] Support Secure String
- [ ] Add history
