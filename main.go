package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"time"
)

var (
	startType string
	logging   slog.Logger
)

func init() {
	flag.StringVar(&startType, "type", "", "server or client")
	flag.Parse()
	if startType == "" {
		fmt.Println("type is required")
		os.Exit(1)
	}
	logging = *slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}

func main() {
	switch startType {
	case "server":
		server()
	case "client":
		client()
	default:
		fmt.Println("type is invalid")
	}

}
func server() {
	logging.Debug("Server is listening on port 5100")
	listener, err := net.Listen("tcp", "0.0.0.0:5100")
	if err != nil {
		logging.Error("Error:", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			logging.Error("Error:", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	logging.Debug("Client connected")
	defer func() {
		conn.Close()
		logging.Debug("Client disconnected")
	}()

	data, err := io.ReadAll(conn)
	if err != nil {
		logging.Error("Error:", err)
		return
	}
	logging.Debug("Received data length:", slog.Int("->", len(data)))
	logging.Debug("Writing data to file")
	f, err := os.OpenFile(fmt.Sprintf("%d.log", time.Now().Unix()), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logging.Error("Error:", err)
		return
	}
	defer func() {
		f.Close()
		logging.Debug("File closed")
	}()
	f.Write(data)
}

func client() {
	conn, err := net.Dial("tcp", "192.168.10.2:5100")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Send message to server
	// message := []byte("Hello from client")
	// _, err = conn.Write(message)
	// if err != nil {
	// 	fmt.Println("Error writing:", err)
	// 	return
	// }
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			logging.Error("Error", err)
			return
		}
		if bytes.Equal(buffer[:n], []byte{2}) {
			continue
		}

		data := string(buffer[:n])
		// if data == "0" {
		// 	continue
		// }
		// sd := strings.Split(data, "\n\n")
		// fmt.Println("->", buffer[:n])
		fmt.Println("->", data)
	}
}
