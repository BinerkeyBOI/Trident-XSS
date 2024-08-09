package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

var extraline strings.Builder
var lines []string
var contnt string
var request string
var finalused int
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"

func main() {
	welcome := `
	/¯¯¯¯¯¯¯¯¯¯¯¯¯\
        ⎢‹—————┐      │
	⎢      ⏐      │
	⎢‹—————┼——————┤
	⎢      ⏐      │
	⎢‹—————┘      │
	\_____________/
	    Trident
		v1.0.3
	`
	fmt.Println(Blue + welcome + Reset)

	// Ports & Targets or Flags
	target := flag.String("t", "localhost", "Define the target to attack.")
	port := flag.String("p", "8080", "Define the port.")
	data := flag.String("d", "", "Gives certain data to the script running.")
	Datafile := flag.String("datafile", "", "Unlike -d --datafile provides a path for a certain file for the script.")
	script := flag.String("script", "null", "Runs script provided.")
	subplace := flag.String("location", "index.php", "Locates where the attack takes place.\n Example: http://127.0.0.1/home/")
	req := flag.Bool("r", false, "To show request that is used.")
	timeout := flag.Int("o", 10000, "Set timeout for request package.")
	//response := flag.Bool("s", false, "Check response from target.")
	debugMode := flag.Bool("debug", false, "Activate Debug mode.")
	flag.Parse()
	if *debugMode {
		fmt.Println(Yellow, "[!] You've entered Debug Mode.")
		fmt.Println(Magenta, "[*] ", Reset, "Data.d and request.dat will not be removed.")
	}
	datafile, err := os.Create("./data.d")
	if err != nil {
		fmt.Println(err)
		return
	}
	datafile.Write([]byte(*data))
	datafile.WriteString("\n" + *Datafile)
	datafile.Close()

	cmd := exec.Command("python3", *script)
	err = cmd.Run()
	if err != nil {
		fmt.Println(Red+"Error running script: "+Reset, err, Yellow+" Did you put the right path?"+Reset)
		return
	}

	// Prettify
	fmt.Println("	" + Cyan + "Method" + Reset + ": " + "POST")
	fmt.Println("	" + Cyan + "Target" + Reset + ": " + *target)
	fmt.Println("	" + Cyan + "Port" + Reset + ": " + *port)
	fmt.Println("	" + Cyan + "Location" + Reset + ": " + *subplace)

	// Connect
	conn, err := net.DialTimeout("tcp", *target+":"+*port, time.Duration(*timeout*1000000))
	if err != nil {
		fmt.Println(Red+"[ERROR] Connection error: "+Reset, err)
		cmd = exec.Command("rm", "./data.d")
		err = cmd.Run()
		cmd = exec.Command("rm", "./request.dat")
		err = cmd.Run()
		if err != nil {
			fmt.Println(Red+"Error finishing script: "+Reset, err, Yellow+" Did yo run in sudo?"+Reset)
			return
		}
		return
	}

	// Encode
	file, err := os.Open("./request.dat")
	if err != nil {
		fmt.Println(Red+"Error opening file:"+Reset, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linenum := -1
	for scanner.Scan() {
		linenum += 1
		line := scanner.Text()
		request = `
POST !
HTTP/2
Host: @
Content-Type: #
`
		if linenum == 0 {
			contnt = line
			request = strings.Replace(request, "#", line, 1)
		} else {
			extraline.WriteString(line)
			extraline.WriteString("\n")
		}
	}

	request = strings.Replace(request, "!", *subplace, 1)
	request = strings.Replace(request, "@", *target, 1)
	request = strings.Replace(request, "#", contnt, 1)
	request += extraline.String()
	if *req {
		fmt.Println(Green, "Request:\n", request)
	}

	if *debugMode {

	} else {
		cmd = exec.Command("rm", "./data.d")
		err = cmd.Run()
		cmd = exec.Command("rm", "./request.dat")
		err = cmd.Run()
		if err != nil {
			fmt.Println(Red+"Error finishing script: "+Reset, err, Yellow+" Did yo run in sudo?"+Reset)
			return
		}
	}

	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println(Red+"Error sending package: ", err)
	}
}
