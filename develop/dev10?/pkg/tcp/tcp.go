package tcpconnector

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"dev10/pkg/arguments"
)

type TcpConnector struct {
	conn net.Conn
}

func MakeTcpConnector(a arguments.Arguments) *TcpConnector {
	var connector TcpConnector
	var err error
	connector.conn, err = net.DialTimeout("tcp", a.Host+":"+a.Port, a.Timeout)
	if err != nil {
		<-time.After(a.Timeout)
		log.Fatal(err)
	}
	return &connector
}

func (c *TcpConnector) RunTelnet(ctx context.Context, wg *sync.WaitGroup) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			err := c.conn.Close()
			if err != nil {
				log.Fatal(err)
			}
			return
		default:
			fmt.Fprintf(c.conn, scanner.Text()+"\n")
			fmt.Print("Message from server: ")
			for {
				message, err := bufio.NewReader(c.conn).ReadString('\n')
				if err == io.EOF {
					break
				} else if err != nil {
					err := c.conn.Close()
					if err != nil {
						log.Fatal(err)
					}
					os.Exit(0)
				}
				fmt.Println(message)
			}
			fmt.Println()
		}
	}
	wg.Done()
}
