package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	hl7svc "hl7_client/hl7_svc"
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
	host      string
	logging   slog.Logger
	fileStore *StorFile
)

type StorFile struct {
	F     *os.File
	Close bool
}

func newStorFile() (*StorFile, error) {
	f, err := os.OpenFile(fmt.Sprintf("%d.csv", time.Now().Unix()), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logging.Error("Error:", slog.String("->", err.Error()))
		return nil, err
	}
	return &StorFile{
		F:     f,
		Close: false,
	}, nil
}

func init() {
	flag.StringVar(&startType, "type", "", "server or client")
	flag.StringVar(&host, "host", "", "server or client")
	flag.Parse()
	if startType == "" {
		fmt.Println("type is required")
		os.Exit(1)
	}

	if host == "" {
		fmt.Println("host is required")
		os.Exit(1)
	}

	logging = *slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		// AddSource: true,

	}))
	fileStore = &StorFile{
		Close: true,
	}
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
	listener, err := net.Listen("tcp", host)
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
	conn, err := net.Dial("tcp", host)
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
	buffer := make([]byte, 1024*1024*10)
	ni := 0
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			logging.Error("Error:", slog.String("->", err.Error()))
			return
		}
		if bytes.Equal(buffer[:n], []byte{2}) {
			ni++
			if !fileStore.Close && ni > 3 {
				fileStore.Close = true
				fileStore.F.Close()
			}
			continue
		}
		ni = 0
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
	if lind != -1 {
		// logging.Debug("Data is not HL7")
		data = data[ind:lind]
		// return
	} else {
		data = data[ind:]

	}
	// fmt.Printf("data: "%s"\n", data) //
	// parceData, err := hl7Decoder.Decode(data)
	logging.Debug("Data:", slog.String("->", string(data)))
	parceData, err := hl7Decoder.Decode(data)
	if err != nil {
		// logging.Debug("sdata:", slog.String("->", string(data)))
		logging.Error("Error:", slog.String("->", err.Error()))
		return
	}
	_ = parceData

	// fmt.Printf("err: %+v\n", err)/

	jData, err := json.Marshal(parceData)
	if err != nil {
		logging.Error("Error:", slog.String("->", err.Error()))
		return
	}
	// logging.Debug("JSON data:", slog.Any("->", jData))
	ob := &hl7svc.Object{}
	err = json.Unmarshal(jData, ob)
	if err != nil {
		// t.Fatal(err)
		logging.Error("Error:", slog.String("->", err.Error()))
		return
	}
	// fmt.Printf("ob: %+v\n", ob)
	// ob.PatientResult[0].OrderObservation[0].Observation[0].OBX
	vals := make(map[string]string)
	for _, v := range ob.PatientResult {
		if len(v.Patient.PID.PatientName) != 0 {
			vals["FNAME"] = v.Patient.PID.PatientName[0].FamilyNameLastNamePrefix
			vals["LNAME"] = v.Patient.PID.PatientName[0].GivenName
		}
		vals["SEX"] = v.Patient.PID.Sex
		vals["BDATE"] = v.Patient.PID.DateTimeOfBirth.Format("02-01-2006")
		if len(v.Patient.PID.PatientIdentifierList) != 0 {
			vals["PIID"] = v.Patient.PID.PatientIdentifierList[0].ID
		}
		vals["PV1"] = v.Patient.Visit.PV1.SetID
		vals["PCLASS"] = v.Patient.Visit.PV1.PatientClass
		for _, v := range v.OrderObservation {
			vals["DATE"] = v.OBR.ObservationDateTime.Format("02-01-2006")
			vals["TIME"] = v.OBR.ObservationDateTime.Format("15:04")
			vals["EntityIdentifier"] = v.OBR.FillerOrderNumber.EntityIdentifier
			for _, v := range v.Observation {
				if len(v.OBX.ObservationValue) != 0 {
					vals[v.OBX.ObservationIdentifier.Text] = fmt.Sprintf("%v", v.OBX.ObservationValue[0])
				}
			}
		}
	}
	logging.Debug("Values:", slog.Any("->", vals))
	writeToFile(vals)
}

func writeToFile(vals map[string]string) {
	var err error
	// fmt.Printf("vals: %v\n", vals)
	daata := fmt.Sprintf(`"%s","%s","%s",WB-CD,"%s","%s",,"%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s",,,,,,,,,,,,,,,,`,
		vals["EntityIdentifier"],
		vals["LNAME"],
		vals["FNAME"],
		vals["DATE"],
		vals["TIME"],
		vals["WBC"],
		vals["NEU#"],
		vals["LYM#"],
		vals["MON#"],
		vals["EOS#"],
		vals["BAS#"],
		vals["NEU%"],
		vals["LYM%"],
		vals["MON%"],
		vals["EOS%"],
		vals["BAS%"],
		vals["RBC"],
		vals["HGB"],
		vals["HCT"],
		vals["MCV"],
		vals["MCH"],
		vals["MCHC"],
		vals["RDW-CV"],
		vals["RDW-SD"],
		vals["PLT"],
		vals["MPV"],
		vals["PDW"],
		vals["PCT"],
		vals["PLCC"],
		vals["PLCR"],
		vals["PIID"],
		vals["SEX"],
		vals["PCLASS"],
		vals["Ref Group"],
		vals["BDATE"],
		vals["Age"],
	)
	if fileStore.Close {
		fileStore.Close = false
		fileStore, err = newStorFile()
		if err != nil {
			logging.Error("Error:", slog.String("->", err.Error()))
			return
		}
	}
	fileStore.F.Write([]byte(daata))
	fileStore.F.Write([]byte("\r\n"))
}
