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

### Binary

```shell
wget https://github.com/stenic/go-mail/releases/latest/download/go-mail_$GOOS_$GOARCH.gz
gunzip go-mail_*.gz
chmod +x go-mail_*
mv go-mail_* /usr/local/bin/go-mail
```

### Docker

```dockerfile
FROM alpine
# ...
COPY --from=ghcr.io/stenic/go-mail:latest /go-mail /usr/local/bin/
# ...
```

### Source

```shell
go install github.com/stenic/go-mail@latest
```

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
