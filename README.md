# examples

以下のトーク内で使用する、あるいは関連する例を集めたリポジトリ

* [TinyGo で組込ソフトウェア開発と自作キーボードに入門しよう](https://docs.google.com/presentation/d/1MUp2QZ98v8KEhcZ9PBw_qj3TjiMKg1u3Vchx3vYD-uY/edit?usp=sharing)

基本的には Wio Terminal 用になっています。

```
$ tinygo flash --target wioterminal --size short --monitor ./01_helloworld/
   code    data     bss |   flash     ram
  10832     148    6680 |   10980    6828
Connected to COM14. Press Ctrl-C to exit.
hello world!
hello world!
hello world!
```

`04_net` および `05_net_server` については `--stack-size 4KB` の設定が必要です。

```
$ tinygo flash --target wioterminal --size short --stack-size 4KB --monitor ./04_net/
   code    data     bss |   flash     ram
 293040    4432   11552 |  297472   15984
Connected to COM14. Press Ctrl-C to exit.

Realtek rtl8720dn Wifi network device driver (rtl8720dn)

Driver version           : 0.28.0
RTL8720 firmware version : 2.1.3
MAC address              : 2c:f7:f1:1b:3d:c3

Connecting to Wifi SSID 'sgmobile'...CONNECTED

DHCP-assigned IP         : 192.168.45.47
DHCP-assigned subnet     : 255.255.255.0
DHCP-assigned gateway    : 192.168.45.195

Getting https://httpbin.org/get?name=John+Doe&occupation=gardener

HTTP/1.1 200 OK
Server: gunicorn/19.9.0
Access-Control-Allow-Origin: *
Access-Control-Allow-Credentials: true
Date: Sat, 14 Sep 2024 00:01:56 GMT
Content-Type: application/json
Content-Length: 331
Connection: keep-alive

{
  "args": {
    "name": "John Doe",
    "occupation": "gardener"
  },
  "headers": {
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/1.1",
    "X-Amzn-Trace-Id": "Root=1-66e4d274-5825d2387d824d9c2b961d7d"
  },
  "origin": "150.31.50.108",
  "url": "https://httpbin.org/get?name=John+Doe&occupation=gardener"
}
```

# LICENSE

MIT
