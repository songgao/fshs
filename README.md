# fshs (File System HTTP Server)

## Usage

```
fshs [laddr]
       ^-- default to ":8080"
```

Serve current directory on HTTP.

* Requests from loopback addresses are severed without questions.
* Other requests go through a (very unscalable) interactive access control.

## why?

I needed a one-line thing like `python3 -m http.server` for quickly
transferring files between servers when `scp` and `rsync` are too slow. I
wanted it to work well with `axel` so it fully utilizes network bandwidth. So I
started this simple utility. As a bonus point, it also features a simple
interactive access control (which is apparently not scalable to anything other
than personal use). Oh and this is easier to type.

## Demo

```shell
$ fshs
Listening at :8080 ...
"GET" Request (/) from [::1]:63737 has been served.
"GET" Request (/main.go) from 127.0.0.1:63754 has been served.
"GET" Request (/) from 172.17.164.35:63769 - "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.15 Safari/537.36"
Options:
  [y] accept;
  [n] reject;
  [d] details;
  [2] accept all requests in following 2 seconds
* Accept? y
The request has been served.
"GET" Request (/main.go) from 172.17.164.35:63769 - "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.15 Safari/537.36"
Options:
  [y] accept;
  [n] reject;
  [d] details;
  [2] accept all requests in following 2 seconds
* Accept? n
The request has been denied with StatusForbidden.
"GET" Request (/fshs) from 172.17.164.35:63784 - "Axel 2.4 (Darwin)"
Options:
  [y] accept;
  [n] reject;
  [d] details;
  [2] accept all requests in following 2 seconds
* Accept? d
&http.Request{Method:"GET", URL:(*url.URL)(0xc2080107e0), Proto:"HTTP/1.0", ProtoMajor:1, ProtoMinor:0, Header:http.Header{"Range":[]string{"bytes=1-"}, "User-Agent":[]string{"Axel 2.4 (Darwin)"}}, Body:(*struct { http.eofReaderWithWriteTo; io.Closer })(0x404d50), ContentLength:0, TransferEncoding:[]string(nil), Close:true, Host:"172.17.164.35", Form:url.Values(nil), PostForm:url.Values(nil), MultipartForm:(*multipart.Form)(nil), Trailer:http.Header(nil), RemoteAddr:"172.17.164.35:63784", RequestURI:"/fshs", TLS:(*tls.ConnectionState)(nil)}
* Accept? 2
The request has been served. All requests within 2 seconds will be served automatically.
"GET" Request (/fshs) from 172.17.164.35:63793 has been served.
"GET" Request (/fshs) from 172.17.164.35:63787 has been served.
"GET" Request (/fshs) from 172.17.164.35:63786 has been served.
"GET" Request (/fshs) from 172.17.164.35:63789 has been served.
"GET" Request (/fshs) from 172.17.164.35:63788 has been served.
"GET" Request (/fshs) from 172.17.164.35:63794 has been served.
"GET" Request (/fshs) from 172.17.164.35:63792 has been served.
"GET" Request (/fshs) from 172.17.164.35:63791 has been served.
"GET" Request (/fshs) from 172.17.164.35:63790 has been served.
"GET" Request (/fshs) from 172.17.164.35:63785 has been served.
```
