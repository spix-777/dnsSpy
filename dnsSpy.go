package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	baseURL := flag.String("u", "url", "Domain. ex: nrk.no")
	wordlist := flag.String("w", "wordlist", "File wordlist")
	out := flag.String("o", "dnsSpoy.txt", "Output file")
	flag.Parse()
	lineString := "-----------------------------------"
	banner := `
·▄▄▄▄   ▐ ▄ .▄▄ · .▄▄ ·  ▄▄▄· ▄· ▄▌
██▪ ██ •█▌▐█▐█ ▀. ▐█ ▀. ▐█ ▄█▐█▪██▌
▐█· ▐█▌▐█▐▐▌▄▀▀▀█▄▄▀▀▀█▄ ██▀·▐█▌▐█▪
██. ██ ██▐█▌▐█▄▪▐█▐█▄▪▐█▐█▪·• ▐█▀·.
▀▀▀▀▀• ▀▀ █▪ ▀▀▀▀  ▀▀▀▀ .▀     ▀ • 

 		... The internet spy`
	fmt.Println(banner)
	fmt.Println(lineString)
	fmt.Println()

	// wordlist file
	file, err := os.Open(*wordlist)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	https := "https://"

	for scanner.Scan() {
		word := scanner.Text()
		url := https + word + "." + *baseURL
		_, err := http.Get(url)
		if err != nil {
			continue
		}

		ip, _ := net.LookupIP(word + "." + *baseURL)

		fmt.Println("\033[32m", "[ + ] ", word+"."+*baseURL)
		fmt.Println("\033[32m", "[ + ]  IP: ", ip)
		fmt.Println("\033[0m")

		var ipStrings []string
		for _, ip := range ip {
			ipStrings = append(ipStrings, ip.String())
		}
		// Output: "192.168.1.1,10.0.0.1"
		vaildIp := strings.Join(ipStrings, " ")
		vaildUrl := word + "." + *baseURL + "\n" + vaildIp + "\n\n"

		// file write
		file, err := os.OpenFile(*out, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		if _, err := file.WriteString(vaildUrl); err != nil {
			log.Fatalln(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(lineString)
}
