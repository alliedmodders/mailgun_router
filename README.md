Mailgun Router
==============

Mailgun Router is an SMTP server that routes requests to Mailgun.

Building
--------

1. Make sure you have Golang installed, (see: http://golang.org/)
2. Make sure your Go environment is set up. Example:
```
mkdir -p ~/go
export GOPATH=/home/$USER/go
```
3. Then, grab the source code:
```
cd $GOPATH
go get github.com/alliedmodders/mailgun_router
cd github.com/alliedmodders/mailgun_router
git submodule update --init --recursive
go install
```
4. The binary will be in `$GOPATH/pkg/`.

Usage
-----

```mailgun_router -config_file=/path/to/config.yaml```

Configuration
-------------

See `extra/sample_config.yaml` for a sample configuration file. You will need your Mailgun API keys
as well as domain.
