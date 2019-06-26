# Cloud Exec

Handle HTTP request and execute script

## Usage

1. Write script
2. Serve script by Cloud Exec

### 1. Write script

```shell
#!/bin/sh

echo Hello World from shellscript
```

### 2. Serve script by Cloud Exec

```shell
$ cloud-exec ./script.sh
```

```shell
$ curl localhost:8080 -i
HTTP/1.1 200 OK
Date: Wed, 26 Jun 2019 05:38:58 GMT
Content-Length: 29
Content-Type: text/plain; charset=utf-8

Hello World from shellscript
```
