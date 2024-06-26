package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"time"

	"github.com/kardianos/hl7"
	"github.com/kardianos/hl7/h231"
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
		// AddSource: true,
		
	}))
}

func main() {
	switch startType {
	case "server":
		server()
	case "client":
		for {
			client()
			<-time.After(time.Second * 10)
		}
	default:
		fmt.Println("type is invalid")
	}

}
func server() {
	logging.Debug("Server is listening on port 5100")
	listener, err := net.Listen("tcp", "0.0.0.0:5100")
	if err != nil {
		logging.Error("Error:", slog.String("->", err.Error()))
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			logging.Error("Error:", slog.String("->", err.Error()))
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
		logging.Error("Error:", slog.String("->", err.Error()))
		return
	}
	logging.Debug("Received data length:", slog.Int("->", len(data)))
	logging.Debug("Writing data to file")
	f, err := os.OpenFile(fmt.Sprintf("%d.log", time.Now().Unix()), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logging.Error("Error:", slog.String("->", err.Error()))
		return
	}
	defer func() {
		f.Close()
		logging.Debug("File closed")
	}()
	f.Write(data)
}

func client() {
	conn, err := net.Dial("tcp", "192.168.0.2:5100")
	if err != nil {
		logging.Error("Error:", slog.String("->", err.Error()))
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
	buffer := make([]byte, 1024*1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			logging.Error("Error:", slog.String("->", err.Error()))
			return
		}
		if bytes.Equal(buffer[:n], []byte{2}) {
			continue
		}

		decoder(buffer[:n])
	}
}

func decoder(data []byte) {
	hl7Decoder := hl7.NewDecoder(h231.Registry, nil)
	ind := bytes.Index(data, []byte("MSH"))
	if ind == -1 {
		logging.Debug("Data is not HL7")
		return
	}
	lind := bytes.LastIndex(data, []byte("\r\x1c\r"))
	if lind == -1 {
		logging.Debug("Data is not HL7")
		return
	}
	data = data[ind:lind]
	// fmt.Printf("data: %s\n", data) //
	// parceData, err := hl7Decoder.Decode(data)
	logging.Debug("Data:", slog.String("->", string(data)))
	parceData, err := hl7Decoder.Decode(data)
	if err != nil {
		logging.Debug("sdata:", slog.String("->", string(data)))
		logging.Error("Error:", slog.String("->", err.Error()))
		return
	}
	// _ = parceData

	// fmt.Printf("err: %+v\n", err)/
	jData, err := json.Marshal(parceData)
	if err != nil {
		logging.Error("Error:", slog.String("->", err.Error()))
		return
	}
	logging.Debug("JSON data:", slog.Any("->", jData))

}
