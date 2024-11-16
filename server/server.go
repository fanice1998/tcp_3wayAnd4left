package main

import (
    "fmt"
    "log"
    "net"
)

func main() {
    l, err := net.Listen("tcp", "0.0.0.0:8081")
    if err != nil {
        log.Fatalln(err)
    }
    defer l.Close()
    for {
        conn, err := l.Accept()
        if err != nil {log.Fatalln(err)}
        fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
        go handler(conn)
    }
}

func handler(conn net.Conn) {
    defer conn.Close()

    for {
        buf := make([]byte, 1024)
        num, err := conn.Read(buf)
        if err != nil {
            fmt.Println(err)
            if err.Error() == "EOF" {
                return
            }
            break
        }
        fmt.Printf("Received data: %v\n", string(buf[:num]))
        // 由於在關閉時 tcp 對列內依然有數據會發送 RST 包
        // 所以這裡屏蔽調數據發送直接關閉連接
        // if num, err = conn.Write([]byte("88")); err != nil {
        //     fmt.Println(err)
        // }
        // 設置了服務端主動段開連接，主要為了完整演示 4 次分手
        if err := conn.Close(); err != nil {
            fmt.Println(err)
        }

        // return
    }
}