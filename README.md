# go-mail

Very basic mail implementation in `Go`. Created to do simple tests as an unprivileged 
user from a single binary.

## Usage

Note that go-mail comes with a full set of defaults. All you will probably need is to
set the `--smtp` and provide a `to-addr`.

```
go-mail --smtp smtp://127.0.0.1:10025 test@test.be
cat somefile.txt | go-mail --smtp smtp://127.0.0.1:10025 test@test.be
```

## Install

```shell
# homebrew
brew install stenic/tap/go-mail

# gofish
gofish rig add https://github.com/stenic/fish-food
gofish install github.com/stenic/fish-food/go-mail

# scoop
scoop bucket add go-mail https://github.com/stenic/scoop-bucket.git
scoop install go-mail

# go
go install github.com/stenic/go-mail@latest

# docker 
docker pull ghcr.io/stenic/go-mail:latest

# dockerfile
COPY --from=ghcr.io/stenic/go-mail:latest /go-mail /usr/local/bin/
```

> For even more options, check the [releases page](https://github.com/stenic/go-mail/releases).

## Run

```shell
# Installed
go-mail -h
# Docker
docker run -ti ghcr.io/stenic/go-mail:latest -h
# Kubernetes
kubectl run go-mail --image=ghcr.io/stenic/go-mail:latest --restart=Never -ti --rm -- -h
```

## Documentation

```shell
$ go-mail -h                                                                                                
go-mail

Usage:
  go-mail [-s subject] to-addr ... [flags]

Flags:
  -b, --bcc-addr stringArray   Send a blink carbon copy to this address
      --body string            Define the message body (default "Hello from go-mail")
  -c, --cc-addr stringArray    Send a carbon copy to this address
  -f, --from string            Send the mail from this address (default "go-mail@example.com")
  -h, --help                   help for go-mail
      --smtp string            Define the smtp server (default "smtp://127.0.0.1:25")
  -s, --subject string         Subject of the mail (default "Test mail from go-mail")

```

## Badges

[![Release](https://img.shields.io/github/release/stenic/go-mail.svg?style=for-the-badge)](https://github.com/stenic/go-mail/releases/latest)
[![Software License](https://img.shields.io/github/license/stenic/go-mail?style=for-the-badge)](./LICENSE.md)
[![Build status](https://img.shields.io/github/workflow/status/stenic/go-mail/Release?style=for-the-badge)](https://github.com/stenic/go-mail/actions?workflow=build)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge)](https://conventionalcommits.org)

## License

[License](./LICENSE)
