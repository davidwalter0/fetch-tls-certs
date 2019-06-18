# fetch-tls-certs

[![Build Status](https://travis-ci.org/davidwalter0/fetch-tls-certs.svg?branch=master)](https://travis-ci.org/davidwalter0/fetch-tls-certs)

Get server's certificate information.

## Installation

For Mac it provide Homebrew integration.
Please install like following.

```sh
$ brew tap davidwalter0/homebrew-cert
$ brew install cert
```

For other platforms, Precompiled binaries for released versions are available in the [releases](https://github.com/davidwalter0/fetch-tls-certs/releases) page.

Or `go get`.

```sh
$ go get github.com/davidwalter0/fetch-tls-certs/...
```

## Usage

Give domain names as arguments.

```sh
$ cert github.com google.co.jp
DomainName: github.com
IP:         192.30.255.113
Issuer:     DigiCert SHA2 Extended Validation Server CA
NotBefore:  2016-03-10 09:00:00 +0900 JST
NotAfter:   2018-05-17 21:00:00 +0900 JST
CommonName: github.com
SANs:       [github.com www.github.com]
Error:

DomainName: google.co.jp
IP:         216.58.196.227
Issuer:     Google Internet Authority G3
NotBefore:  2017-10-17 19:59:51 +0900 JST
NotAfter:   2018-01-09 19:00:00 +0900 JST
CommonName: *.google.co.jp
SANs:       [*.google.co.jp google.co.jp]
Error:

```

You can specify port number.
So you can get server certificate information of not only web server but also *mail server and others*.

```sh
$ cert github.com google.co.jp:443 imap.gmail.com:993
DomainName: github.com
IP:         192.30.255.113
Issuer:     DigiCert SHA2 Extended Validation Server CA
NotBefore:  2016-03-10 09:00:00 +0900 JST
NotAfter:   2018-05-17 21:00:00 +0900 JST
CommonName: github.com
SANs:       [github.com www.github.com]
Error:

DomainName: google.co.jp
IP:         172.217.27.163
Issuer:     Google Internet Authority G3
NotBefore:  2017-10-17 19:59:51 +0900 JST
NotAfter:   2018-01-09 19:00:00 +0900 JST
CommonName: *.google.co.jp
SANs:       [*.google.co.jp google.co.jp]
Error:

DomainName: imap.gmail.com
IP:         64.233.188.108
Issuer:     Google Internet Authority G2
NotBefore:  2017-10-17 19:10:29 +0900 JST
NotAfter:   2017-12-29 09:00:00 +0900 JST
CommonName: imap.gmail.com
SANs:       [imap.gmail.com]
Error:

```

## Options

```sh
$ cert --help
Usage of cert:
  -f string
        Output format. md: as markdown, json: as JSON.  (default "simple table")
  -format string
        Output format. md: as markdown, json: as JSON.  (default "simple table")
  -k    Skip verification of server's certificate chain and host name.
  -s int
        Timeout seconds. (default 3)
  -skip-verify
        Skip verification of server's certificate chain and host name.
  -t string
        Output format as Go template string or Go template file path.
  -template string
        Output format as Go template string or Go template file path.
  -timeout int
        Timeout seconds. (default 3)
  -u    Use UTC to represent NotBefore and NotAfter.
  -utc
        Use UTC to represent NotBefore and NotAfter.
  -v    Show version.
  -version
        Show version.
```

### Output as JSON

Use `cert -f json`.

```sh
$ cert -f json github.com | jq .
[
  {
    "DomainName": "github.com",
    "IP": "192.30.255.112",
    "Issuer": "DigiCert SHA2 Extended Validation Server CA",
    "CommonName": "github.com",
    "SANs": [
      "github.com",
      "www.github.com"
    ],
    "NotBefore": "2016-03-10 09:00:00 +0900 JST",
    "NotAfter": "2018-05-17 21:00:00 +0900 JST",
    "Error": ""
  }
]
```

### Output as Markdown

Use `cert -f md`.

```sh
$ cert -f md github.com
DomainName | IP | Issuer | NotBefore | NotAfter | CN | SANs | Error
--- | --- | --- | --- | --- | --- | --- | ---
github.com | 192.30.255.113 | DigiCert SHA2 Extended Validation Server CA | 2016-03-10 09:00:00 +0900 JST | 2018-05-17 21:00:00 +0900 JST | github.com | github.com<br/>www.github.com<br/> |
```

DomainName | IP | Issuer | NotBefore | NotAfter | CN | SANs | Error
--- | --- | --- | --- | --- | --- | --- | ---
github.com | 192.30.255.113 | DigiCert SHA2 Extended Validation Server CA | 2016-03-10 09:00:00 +0900 JST | 2018-05-17 21:00:00 +0900 JST | github.com | github.com<br/>www.github.com<br/> |

### Specify output format by Go template

Use `cert -t`.

By direct string.

```sh
$ cert -t '{{range .}}{{.Issuer}}{{end}}' github.com
DigiCert SHA2 Extended Validation Server CA
```

By template file.

```sh
$ cat /tmp/cert_templ
{{range .}}{{range .CertChain}}Issuer: {{.Issuer.CommonName}}
{{end}}{{end}}
$
$ cert -t /tmp/cert_templ github.com
Issuer: DigiCert SHA2 Extended Validation Server CA
Issuer: DigiCert High Assurance EV Root CA

```

## License

[MIT](https://github.com/davidwalter0/fetch-tls-certs/blob/master/LICENSE)

## Author

[davidwalter0](https://github.com/davidwalter0)

