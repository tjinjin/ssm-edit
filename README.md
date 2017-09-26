# ssm-edit

Edit a parameter  on SystemsManger parameter store.

Inspiread by https://github.com/tsub/s3-edit

# Install

```
$ make install
```

# Usage

```
$ ssm-edit list
$ ssm-edit edit
$ ssm-edit show
```

```
# all
global flag
--profile

# edit show
local flag
--name
```


# Todo
- [x] Support assumeRole
- [ ] Support Pagination to list command
- [ ] Add `create` command
- [ ] Add `show` command
- [ ] Add Makefile
- [ ] Add test
- [ ] Add CI
- [ ] Refactoring!!!
