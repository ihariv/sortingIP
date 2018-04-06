package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
)

// main2 generate list of ip for test
func main2() {
	start := ip2Long("192.168.10.1")
	f, _ := os.Create("input.txt")
	for i := 100000; i > 0; i-- {
		start--
		f.WriteString(backtoIP4(int64(start)) + "\n")
	}
	f.Close()
}

func main() {

	// get parameters
	inPtr := flag.String("in", "input.txt", "select input file. example: -in=input.txt")
	outPtr := flag.String("out", "", "select output file. example: -out=output.txt")
	sepPtr := flag.String("sep", ",", "select separator string. example: -sep=,")
	flag.Parse()

	realIPs := []uint32{}

	file, err := os.Open(*inPtr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read file by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ipString := scanner.Text()

		//search separator
		if sep := strings.Index(ipString, *sepPtr); sep > -1 {
			// add ip for range
			for i := ip2Long(ipString[:sep]); i <= ip2Long(ipString[sep+1:]); i++ {
				realIPs = append(realIPs, i)
			}
		} else {
			// add ip for line
			realIPs = append(realIPs, ip2Long(scanner.Text()))
		}

		// clear memory
		ipString = ""
	}

	// sorting by ip
	sort.Slice(realIPs, func(i, j int) bool {
		return realIPs[i] < realIPs[j]
	})

	// lastIP is previous ip from work range
	lastIP := uint32(0)

	// deltaIP is previous ip from emitated range
	deltaIP := uint32(0)

	for _, ip := range realIPs {
		if lastIP == 0 {
			// init first IP
			lastIP = ip
			deltaIP = ip
			continue
		} else {

			// ignore dublicates
			if ip == deltaIP {
				continue
			}

			// ip not is next of delta
			if ip != deltaIP+1 {

				// ip is virtual range
				if deltaIP != lastIP {
					saveOutput(backtoIP4(int64(lastIP))+"-"+backtoIP4(int64(deltaIP)), *outPtr)
					deltaIP = ip
				} else {
					//ip not is virtual range
					saveOutput(backtoIP4(int64(lastIP)), *outPtr)
				}

				// member lastIP
				lastIP = ip
			}

		}
		// member deltaIP
		deltaIP = ip

	}

	// save last data
	if deltaIP != lastIP {
		saveOutput(backtoIP4(int64(lastIP))+"-"+backtoIP4(int64(deltaIP)), *outPtr)
	} else {
		saveOutput(backtoIP4(int64(lastIP)), *outPtr)
	}
}

// saveOutput -- save data in file or stdOut
func saveOutput(text, file string) {

	if file == "" {
		fmt.Println(text)
	} else {
		f, _ := os.Create(file)

		n3, _ := f.WriteString(text + "\n")
		fmt.Printf("wrote %d bytes\n", n3)
	}
}

// ip2Long is faster convert from string to uint32
func ip2Long(ip string) uint32 {
	var long uint32
	binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	return long
}

// backtoIP4 is faster convert from int64 to string
func backtoIP4(ipInt int64) string {

	// need to do two bit shifting and “0xff” masking
	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt((ipInt & 0xff), 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}
