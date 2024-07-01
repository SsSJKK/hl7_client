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
	"strings"
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

var vsdata = `MSH|^~\\&|||||20240701145039||ORU^R01|35|P|2.3.1||||||UNICODE
PID|1||^^^^MR
PV1|1
OBR|1||16|00001^Automated Count^99MRC|||20240701104947|||||||||||||||||HM||||||||1
OBX|1|IS|08001^Take Mode^99MRC||O||||||F
OBX|2|IS|08002^Blood Mode^99MRC||W||||||F
OBX|3|IS|01002^Ref Group^99MRC||Общая||||||F
OBX|4|NM|6690-2^WBC^LN||5.5|10*9/L|4.0-10.0|N|||F
OBX|5|NM|731-0^LYM#^LN||2.7|10*9/L|0.8-4.0|N|||F
OBX|6|NM|736-9^LYM%^LN||48.3|%|20.0-40.0|H~N|||F
OBX|7|NM|789-8^RBC^LN||4.29|10*12/L|3.50-5.50|N|||F
OBX|8|NM|718-7^HGB^LN||118|g/L|110-160|N|||F
OBX|9|NM|787-2^MCV^LN||73.3|fL|80.0-100.0|L~N|||F
OBX|10|NM|785-6^MCH^LN||27.5|pg|27.0-34.0|N|||F
OBX|11|NM|786-4^MCHC^LN||375|g/L|320-360|H~N|||F
OBX|12|NM|788-0^RDW-CV^LN||11.6|%|11.0-16.0|N|||F
OBX|13|NM|21000-5^RDW-SD^LN||31.8|fL|35.0-56.0|L~N|||F
OBX|14|NM|4544-3^HCT^LN||31.4|%|37.0-54.0|L~N|||F
OBX|15|NM|777-3^PLT^LN||310|10*9/L|180-320|N|||F
OBX|16|NM|32623-1^MPV^LN||9.2|fL|6.5-12.0|N|||F
OBX|17|NM|32207-3^PDW^LN||15.8||15.0-17.0|N|||F
OBX|18|NM|10002^PCT^99MRC||0.286|%|0.108-0.282|H~N|||F
OBX|19|NM|10027^MID#^99MRC||0.6|10*9/L|0.1-1.5|N|||F
OBX|20|NM|10029^MID%^99MRC||10.6|%|3.0-15.0|N|||F
OBX|21|NM|10028^GRAN#^99MRC||2.2|10*9/L|2.0-7.0|N|||F
OBX|22|NM|10030^GRAN%^99MRC||41.1|%|50.0-70.0|L~N|||F`

// var vsdata = `MSH|^~\\&|||||20240628210706||ORU^R01|804|P|2.3.1||||||UNICODE
// PID|1||ppid^^^^MR||фамишия^имя|||Муж
// PV1|1|Амбул. больной|^^место
// OBR|1||4|00001^Automated Count^99MRC||20240526123400|20240529133328|||||||20240526124500||||||||||HM||||||||1
// OBX|1|IS|08001^Take Mode^99MRC||O||||||F
// OBX|2|IS|08002^Blood Mode^99MRC||W||||||F
// OBX|3|IS|08003^Test Mode^99MRC||CBC+DIFF||||||F
// OBX|4|IS|01002^Ref Group^99MRC||Взрос.муж||||||F
// OBX|5|NM|30525-0^Age^LN||66|yr|||||F
// OBX|6|ST|01001^Remark^99MRC||комменнт||||||F
// OBX|7|NM|6690-2^WBC^LN||6.27|10*9/L|4.00-10.00|A|||F
// OBX|8|NM|704-7^BAS#^LN||0.02|10*9/L|0.00-0.10|A|||F
// OBX|9|NM|706-2^BAS%^LN||0.3|%|0.0-1.0|A|||F
// OBX|10|NM|751-8^NEU#^LN||3.94|10*9/L|2.00-7.00|A|||F
// OBX|11|NM|770-8^NEU%^LN||62.8|%|50.0-70.0|A|||F
// OBX|12|NM|711-2^EOS#^LN||0.39|10*9/L|0.02-0.50|A|||F
// OBX|13|NM|713-8^EOS%^LN||6.1|%|0.5-5.0|H~A|||F
// OBX|14|NM|731-0^LYM#^LN||1.61|10*9/L|0.80-4.00|A|||F
// OBX|15|NM|736-9^LYM%^LN||25.7|%|20.0-40.0|A|||F
// OBX|16|NM|742-7^MON#^LN||0.31|10*9/L|0.12-1.20|A|||F
// OBX|17|NM|5905-5^MON%^LN||5.1|%|3.0-12.0|A|||F
// OBX|18|NM|789-8^RBC^LN||4.35|10*12/L|4.00-5.50|A|||F
// OBX|19|NM|718-7^HGB^LN||125|g/L|120-160|A|||F
// OBX|20|NM|787-2^MCV^LN||84.4|fL|80.0-100.0|A|||F
// OBX|21|NM|785-6^MCH^LN||28.6|pg|27.0-34.0|A|||F
// OBX|22|NM|786-4^MCHC^LN||339|g/L|320-360|A|||F
// OBX|23|NM|788-0^RDW-CV^LN||12.5|%|11.0-16.0|A|||F
// OBX|24|NM|21000-5^RDW-SD^LN||41.7|fL|35.0-56.0|A|||F
// OBX|25|NM|4544-3^HCT^LN||36.8|%|40.0-54.0|L~A|||F
// OBX|26|NM|777-3^PLT^LN||259|10*9/L|100-300|A|||F
// OBX|27|NM|32623-1^MPV^LN||9.0|fL|6.5-12.0|A|||F
// OBX|28|NM|32207-3^PDW^LN||15.9||9.0-17.0|A|||F
// OBX|29|NM|10002^PCT^99MRC||0.234|%|0.108-0.282|A|||F
// OBX|30|NM|10013^PLCC^99MRC||52|10*9/L|30-90|A|||F
// OBX|31|NM|10014^PLCR^99MRC||20.2|%|11.0-45.0|A|||F`

var (
	header = "ID пробы,Имя,Фамилия,Режим,Дата,Время,Состоян.пробы,WBC (10^9/L),Neu# (10^9/L),Lym# (10^9/L),Mon# (10^9/L),Eos# (10^9/L),Bas# (10^9/L),Neu% (%),Lym% (%),Mon% (%),Eos% (%),Bas% (%),RBC (10^12/L),HGB (g/L),HCT (%),MCV (fL),MCH (pg),MCHC (g/L),RDW-CV (%),RDW-SD (fL),PLT (10^9/L),MPV (fL),PDW ( ),PCT (%),P-LCC (10^9/L),P-LCR (%),ID пациента,Пол,Тип пациента,Конт.группа,Дата рождения,Возраст,Отделение,Место №,Дата отбора,Время отбора,Дата доставки,Время доставки,Врач,Оператор,Проверил(а),Комментарии,Сообщ.о WBC,Сообщ.о RBC,Сообщ.о PLT,Группа крови,РОЭ,Микроскопич.парам."
	layout = `"EntityIdentifier","LNAME",WB-CD,"FNAME","DATE",,"TIME","WBC","NEU#","LYM#","MON#","EOS#","BAS#","NEU%","LYM%","MON%","EOS%","BAS%","RBC","HGB","HCT","MCV","MCH","MCHC","RDW-CV","RDW-SD","PLT","MPV","PDW","PCT","PLCC","PLCR","PIID","SEX","PCLASS","Ref Group","BDATE","Age",,,,,,,,,,,,,,,,`
)

func newStorFile() (*StorFile, error) {
	f, err := os.OpenFile(fmt.Sprintf("%d.csv", time.Now().Unix()), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logging.Error("Error:", slog.String("->", err.Error()))
		return nil, err
	}
	f.WriteString(header + "\r\n")
	return &StorFile{
		F:     f,
		Close: false,
	}, nil
}

func init() {
	flag.StringVar(&startType, "type", "", "server or client")
	flag.StringVar(&host, "host", "", "server or client")
	flag.Parse()

	// if startType == "" {
	// 	fmt.Println("type is required")
	// 	os.Exit(1)
	// }

	// if host == "" {
	// 	fmt.Println("host is required")
	// 	os.Exit(1)
	// }

	logging = *slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		// AddSource: true,

	}))
	fileStore = &StorFile{
		Close: true,
	}
	confRead()
}

func confRead() {
	fConf, err := os.Open("config")
	if err != nil {
		logging.Error("Error:", slog.String("->", err.Error()))
		return
	}
	defer fConf.Close()
	dataConf, err := io.ReadAll(fConf)
	if err != nil {
		logging.Error("Error:", slog.String("->", err.Error()))
		return
	}
	dataConf = bytes.ReplaceAll(dataConf, []byte("\r\n"), []byte("\n"))
	rData := bytes.Split(dataConf, []byte("\n"))
	vals := make(map[string]string)
	for _, v := range rData {
		vv := bytes.Split(v, []byte("="))
		if len(vv) != 2 {
			continue
		}
		vals[string(vv[0])] = string(vv[1])
	}

	if v, ok := vals["header"]; ok && v != "" {
		header = v
	}
	if v, ok := vals["layout"]; ok && v != "" {
		layout = v
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
		decoder([]byte(vsdata))
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
	// data := fmt.Sprintf(`"%s","%s","%s",WB-CD,"%s","%s",,"%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s",,,,,,,,,,,,,,,,`,
	// 	vals["EntityIdentifier"],
	// 	vals["LNAME"],
	// 	vals["FNAME"],
	// 	vals["DATE"],
	// 	vals["TIME"],
	// 	vals["WBC"],
	// 	vals["NEU#"],
	// 	vals["LYM#"],
	// 	vals["MON#"],
	// 	vals["EOS#"],
	// 	vals["BAS#"],
	// 	vals["NEU%"],
	// 	vals["LYM%"],
	// 	vals["MON%"],
	// 	vals["EOS%"],
	// 	vals["BAS%"],
	// 	vals["RBC"],
	// 	vals["HGB"],
	// 	vals["HCT"],
	// 	vals["MCV"],
	// 	vals["MCH"],
	// 	vals["MCHC"],
	// 	vals["RDW-CV"],
	// 	vals["RDW-SD"],
	// 	vals["PLT"],
	// 	vals["MPV"],
	// 	vals["PDW"],
	// 	vals["PCT"],
	// 	vals["PLCC"],
	// 	vals["PLCR"],
	// 	vals["PIID"],
	// 	vals["SEX"],
	// 	vals["PCLASS"],
	// 	vals["Ref Group"],
	// 	vals["BDATE"],
	// 	vals["Age"],
	// )
	data := layout
	layoutC := strings.ReplaceAll(layout, `"`, "")
	for _, v := range strings.Split(layoutC, ",") {
		if v == "" {
			continue
		}
		if val, ok := vals[v]; ok {
			data = strings.Replace(data, v, val, 1)
		}
	}
	if fileStore.Close {
		fileStore.Close = false
		fileStore, err = newStorFile()
		if err != nil {
			logging.Error("Error:", slog.String("->", err.Error()))
			return
		}
	}
	fileStore.F.Write([]byte(data))
	fileStore.F.Write([]byte("\r\n"))
}
