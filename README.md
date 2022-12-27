Parse otp secret from the migration QR code of Google authenticator

## Install
```shell
go install github.com/dchaofei/gauth-exporter@latest
```

## Usage

- The first step is to parse the QR code of Google authenticator into `otpauth-migration:offline?data=xxxxxxxx`
- The second step is to run the command line program

```shell
$ gauth-exporter -uri "otpauth-migration://offline?data=xxxxx"

#output
[
  {
    "secret": "xxxxx",
    "name": "xxxxx",
    "issuer": "xxxx",
    "algorithm": "ALGO_SHA1",
    "digits": 1,
    "type": "OTP_TOTP"
  }
]

```

#### Refer to other language implementations

- https://github.com/krissrex/google-authenticator-exporter