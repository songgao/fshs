# fshs (File System HTTP Server)

## Usage

```
fshs [laddr]
       ^-- default to ":8080"
```

Serve current directory on HTTP.

## why?

I needed a one-line thing like `python3 -m http.server` for quickly transferring files between servers when `scp` and `rsync` are too slow. I wanted it to work well with `axel` so it fully utilizes network bandwidth. So I started this simple utility. As a bonus point, it also features a simple interactive access control (which is apparently not scalable to anything other than personal use). Oh and this is easier to type.

I don't see this useful for anything other than quickly transfering files through HTTP though.

## Demo

```shell
$ ./fshs
Listening at :8080 ...
GET Request (/) from 127.0.0.1:51456 - "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11) AppleWebKit/601.1.41 (KHTML, like Gecko) Version/9.0 Safari/601.1.41"
Options:
  [y] accept;
  [n] reject;
  [d] details;
  [2] accept all requests in following 2 seconds
* Accept? y
The request has been served.
GET Request (/main.go) from 127.0.0.1:51456 - "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11) AppleWebKit/601.1.41 (KHTML, like Gecko) Version/9.0 Safari/601.1.41"
Options:
  [y] accept;
  [n] reject;
  [d] details;
  [2] accept all requests in following 2 seconds
* Accept? n
The request has been denied with StatusForbidden.
GET Request (/fshs) from 127.0.0.1:51458 - "Axel 2.4 (Darwin)"
Options:
  [y] accept;
  [n] reject;
  [d] details;
  [2] accept all requests in following 2 seconds
* Accept? d
&http.Request{Method:"GET", URL:(*url.URL)(0xc2080101c0), Proto:"HTTP/1.0", ProtoMajor:1, ProtoMinor:0, Header:http.Header{"User-Agent":[]string{"Axel 2.4 (Darwin)"}, "Range":[]string{"bytes=1-"}}, Body:(*struct { http.eofReaderWithWriteTo; io.Closer })(0x404d50), ContentLength:0, TransferEncoding:[]string(nil), Close:true, Host:"localhost", Form:url.Values(nil), PostForm:url.Values(nil), MultipartForm:(*multipart.Form)(nil), Trailer:http.Header(nil), RemoteAddr:"127.0.0.1:51458", RequestURI:"/fshs", TLS:(*tls.ConnectionState)(nil)}
* Accept? 2
The request has been served. All requests within 2 seconds will be served automatically.
"GET" Request (/fshs) from 127.0.0.1:51464 has been served.
"GET" Request (/fshs) from 127.0.0.1:51463 has been served.
"GET" Request (/fshs) from 127.0.0.1:51462 has been served.
"GET" Request (/fshs) from 127.0.0.1:51460 has been served.
"GET" Request (/fshs) from 127.0.0.1:51469 has been served.
"GET" Request (/fshs) from 127.0.0.1:51468 has been served.
"GET" Request (/fshs) from 127.0.0.1:51467 has been served.
"GET" Request (/fshs) from 127.0.0.1:51466 has been served.
"GET" Request (/fshs) from 127.0.0.1:51465 has been served.
"GET" Request (/fshs) from 127.0.0.1:51461 has been served.
```
