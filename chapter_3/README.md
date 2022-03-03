## Deployment

```
docker-compose up
```

## Executing

### Standard curl
- client

```
$ curl localhost:18888
<html><body>hello</body></html>
```

- server log

```
simple-server_1  | GET / HTTP/1.1
simple-server_1  | Host: localhost:18888
simple-server_1  | Accept: */*
simple-server_1  | User-Agent: curl/7.68.0
```

### Curl with HTTP/1.0

- client

```
$ curl --http1.0 http://localhost:18888
<html><body>hello</body></html>
```

- server log

```
simple-server_1  | GET / HTTP/1.0
simple-server_1  | Host: localhost:18888
simple-server_1  | Connection: close
simple-server_1  | Accept: */*
simple-server_1  | User-Agent: curl/7.68.0
```

### Request by netcat
#### HTTP/0.9

- client

```
$ echo -e "GET / HTTP/0.9\r\n" | nc localhost 18888
HTTP/1.1 505 HTTP Version Not Supported: unsupported protocol version
Content-Type: text/plain; charset=utf-8
Connection: close

505 HTTP Version Not Supported: unsupported protocol version
```

#### HTTP/1.0

- client

```
$ echo -e "GET / HTTP/1.0\r\nHost: localhost:18888\r\n" | nc localhost 18888
HTTP/1.0 200 OK
Date: Wed, 05 Jan 2022 22:34:41 GMT
Content-Length: 32
Content-Type: text/html; charset=utf-8

<html><body>hello</body></html>
```

- server log

```
simple-server_1  | GET / HTTP/1.0
simple-server_1  | Connection: close
```