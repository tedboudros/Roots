package main

import (
    "fmt"
    "net"
    "os"
    "time"
    "io/ioutil"
    "strings"
    "path/filepath"
)

const (
    ConnHost = "0.0.0.0"
    ConnPort = "3333"
    ConnType = "tcp"
)

func main() {
    connectToDatabase("root","","roots")
    // Listen for incoming connections.
    l, err := net.Listen(ConnType, ConnHost+":"+ConnPort)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + ConnHost + ":" + ConnPort)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) int{

    // Make a buffer to hold incoming data.
    buf := make([]byte, 1024)
    // Read the incoming connection into the buffer.

    t := time.Now().UTC()

    _, error := conn.Read(buf)
    if error != nil {
        fmt.Println("Error reading:", error.Error())
    }

    // Split on comma.
    result := strings.Split(string(buf), "\n")

    // fmt.Println(string(buf)) FOR DEBUGGING
    file := ""
    str := ""
    // Display all elements.
    for i := range result {
        if strings.HasPrefix(result[i], "GET") == true{
            file = strings.TrimPrefix(strings.Split(result[i], " ")[1], "/")
            file = strings.Split(file, "?")[0]
            if file == "" { file = "index.rsl" }
            if strings.HasPrefix(file, "http") {
                str = file
                break
            }
            fmt.Println("FILE REQUEST: " + file)
        }
        
    }

    if str == "" {
        filepathabs, er := filepath.Abs(file)
        if er != nil {fmt.Println(er)}
        b, err := ioutil.ReadFile(filepathabs) // just pass the file name
        if err != nil {
            fmt.Print(err)
            conn.Write([]byte("HTTP/1.1 404 Not Found\n"))
        } else  { conn.Write([]byte("HTTP/1.1 200 OK\n")) }
            str = string(b) // convert content to a 'string'
    }
    
    conn.Write([]byte(t.Format("Date: Mon, 02 Jan 2006 15:04:05 GMT") + "\n"))


    if strings.HasSuffix(file, ".svg") == true { conn.Write([]byte("Content-Type: image/svg+xml \n")) }
    if strings.HasSuffix(file, ".css") == true { conn.Write([]byte("Content-Type: text/css \n")) }
    if strings.HasSuffix(file, ".js") == true { conn.Write([]byte("Content-Type: text/javascript \n")) }
    if strings.HasSuffix(file, ".html") == true { conn.Write([]byte("Content-Type: text/html \n")) }
    if strings.HasSuffix(file, ".rsl") == true { conn.Write([]byte("Content-Type: text/html \n")) }
    if strings.HasSuffix(file, ".mp3") == true { conn.Write([]byte("Content-Type: audio/mp3 \n")) }

    conn.Write([]byte("Server: RlS(GoLang)\nConnection: Closed"))
    conn.Write([]byte("\n\n"))//Time for the html part
    conn.Write([]byte(str)) // HTML
    conn.Close()
    return 0
}