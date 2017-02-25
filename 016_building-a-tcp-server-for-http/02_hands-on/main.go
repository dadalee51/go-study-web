package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	c := make(chan string)
	fmt.Println("chan made")
	// read request
	request(conn, c)
	fmt.Println("requested")
	// write response
	respond(conn, c)
	fmt.Println("responded")
}

func request(conn net.Conn, c chan string) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[1]
			fmt.Println("***URL", m)
			//a conventional way for the channel to kick off
			go func(){
				c<-m
				fmt.Println("sending through.")
			}()
			fmt.Println("are we there yet.")
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func respond(conn net.Conn, c chan string) {
	fmt.Println("responding...")
	d:= <-c
	fmt.Println("read!")
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>The URL you requested was:`+ d + `</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
