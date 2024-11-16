# tcp_3wayAnd4left
## 環境配置

Go version 1.23.3

Linux 需安裝 libpcap-dev

Windows 需安裝 [Npcap](https://npcap.com/#download)

## Go package
```
go get github.com/google/gopacket
```

## C package
因為 gopacket 是建立在 libpcap-dev 上的
所以要使用 gopacket 需要安裝 libpcap-dev 不然可能會有找不到檔案導致無法使用的問題
或是安裝 Npcap
### Linux install libpcap-dev
```
sudo apt-get install libpcap-dev
```

### Windows install Npcap
[Npcap 官網連結](https://npcap.com/#download)

* Npcap 1.80 installer for Windows 7/2008R2, 8/2012, 8.1/2012R2, 10/2016, 2019, 11 (x86, x64, and ARM64).


## 執行測試
需要三個 session
分別執行 server.go capture.go client.go
```
cd tcp_3wayAnd4left/server
go run server.go

cd tcp_3wayAnd4left/pcap
go run capture.go

cd tcp_3wayAnd4left/client
go run client.go
```
