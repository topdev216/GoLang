package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var ipAddrs chan string = make(chan string)
var result chan string = make(chan string)
var thread chan int = make(chan int)
var nowThread int
var clo chan bool = make(chan bool)

func writeResult() {
	fileName := "scan.txt"
	fout, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Errorf("%s was failed to create\n", fileName)
	}
	defer fout.Close()
	s, ok := <-result
	for ok {
		log.Printf("\tFOUND TARGET: %s", s)
		fout.WriteString(s + "\n")
		s, ok = <-result
	}
	clo <- true
}

func runScan() {
	t, ok := <-thread
	nowThread = t
	if ok {
		for i := 0; i < nowThread; i++ {
			go scan(strconv.Itoa(i))
		}
	}
	for <-thread == 0 {
		nowThread--
		if nowThread == 0 {
			close(result)
			break
		}
	}
}

func scan(threadId string) {
	s, ok := <-ipAddrs
	for ok {
		if ip2num(s[0:strings.Index(s, ":")])&0xFF == 0 {
			log.Printf("GO ROUTINE: %s, IP BLOCK: %s", threadId, s)
		}
		_, err := net.Dial("tcp", s)
		if err == nil {
			result <- s
		}
		s, ok = <-ipAddrs
	}
	thread <- 0
}

func num2ip(ip int) string {
	return strconv.FormatInt((int64(ip)&0xFF000000)>>24, 10) + "." +
		strconv.FormatInt((int64(ip)&0x00FF0000)>>16, 10) + "." +
		strconv.FormatInt((int64(ip)&0x0000FF00)>>8, 10) + "." +
		strconv.FormatInt((int64(ip)&0x000000FF), 10)
}

func ip2num(ip string) int {
	ips := strings.Split(ip, ".")
	n0, err := strconv.Atoi(ips[0])
	if err != nil {
		return 0
	}
	n1, err := strconv.Atoi(ips[1])
	if err != nil {
		return 0
	}
	n2, err := strconv.Atoi(ips[2])
	if err != nil {
		return 0
	}
	n3, err := strconv.Atoi(ips[3])
	if err != nil {
		return 0
	}
	return n0<<24 | n1<<16 | n2<<8 | n3
}

func processIp(startIp, endIp string) []string {
	nStartIp := ip2num(startIp)
	nEndIp := ip2num(endIp)

	resultIp := make([]string, 0) // nEndIp-nStartIp+1)
	for i := 0; i <= nEndIp-nStartIp; i++ {
		resultIp = append(resultIp, num2ip(nStartIp+i))
	}
	log.Printf("THERE ARE %d HOSTS BETWEEN %s AND %s.", len(resultIp), startIp, endIp)
	return resultIp
}

func processFlag(arg []string) {
	var startIp, endIp string
	var ports []int = make([]int, 0)
	index := 0
	startIp = arg[index]
	si := net.ParseIP(startIp)
	if si == nil {
		fmt.Println("startip configure error")
		os.Exit(1)
	}

	// 起止IP范围
	index++
	endIp = arg[index]
	ei := net.ParseIP(endIp)
	if ei == nil {
		endIp = startIp
	} else {
		index++
	}

	// 端口范围列表
	tmpPort := arg[index]
	if strings.Index(tmpPort, "-") != -1 {
		tmpPorts := strings.Split(tmpPort, "-")
		var startPort, endPort int
		var err error
		startPort, err = strconv.Atoi(tmpPorts[0])
		if err != nil || startPort < 1 || startPort > 65535 {
			fmt.Println("start port number error")
			os.Exit(1)
		}
		if len(tmpPorts) >= 2 {
			endPort, err = strconv.Atoi(tmpPorts[1])
			if err != nil || endPort < 1 || endPort > 65535 || endPort < startPort {
				fmt.Println("endPort configure error")
				os.Exit(1)
			}
		} else {
			endPort = 65535
		}
		for i := 0; startPort+i <= endPort; i++ {
			ports = append(ports, startPort+i)
		}
	} else {
		ps := strings.Split(tmpPort, ",")
		for i := 0; i < len(ps); i++ {
			p, err := strconv.Atoi(ps[i])
			if err != nil {
				fmt.Println("port configure error")
				os.Exit(1)
			}
			ports = append(ports, p)
		}
	}

	// go routine 数量
	index++
	t, err := strconv.Atoi(arg[index])
	if err != nil {
		fmt.Println("thread configure error")
		os.Exit(1)
	}

	if t < 1 {
		t = 1
	}

	if t > 512 {
		log.Println("TOO MANY GO ROUTINE MAY CAUSE A CRASH.")
	}

	thread <- t
	ips := processIp(startIp, endIp)
	il := len(ips)
	for i := 0; i < il; i++ {
		pl := len(ports)
		for j := 0; j < pl; j++ {
			ipAddrs <- ips[i] + ":" + strconv.Itoa(ports[j])
		}
	}
	close(ipAddrs)
}

func main() {
	flag.Parse()
	if flag.NArg() != 3 && flag.NArg() != 4 {
		fmt.Println("Usage:\n./scan startip endip port thread_num\n192.168.0.1 192.168.1.255 80,21-25 128")
		os.Exit(1)
	}

	args := make([]string, 0, 4)
	for i := 0; i < flag.NArg(); i++ {
		args = append(args, flag.Arg(i))
	}

	go runScan()
	go writeResult()
	processFlag(args)
	<-clo
}
