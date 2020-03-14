package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	Code int
	Msg  string
}

func (r Response) IsOK() bool { return 200 <= r.Code && r.Code < 400 }

func Parse(data []byte) (*Response, error) {
	subs := strings.SplitN(string(data), " ", 2)
	if len(subs) < 2 {
		return nil, fmt.Errorf("could not parse response %q", string(data))
	}
	code, err := strconv.Atoi(subs[0])
	if err != nil {
		return nil, fmt.Errorf("could not parse response code %q", subs[0])
	}
	return &Response{Code: code, Msg: subs[1]}, nil
}

func Exchange(ctx context.Context, c net.Conn, data []byte) ([]byte, error) {

	done := make(chan struct{})
	buf := make([]byte, 1024)
	var n int
	var err error

	go func() {
		defer close(done)
		_, err = c.Write(data)
		if err != nil {
			return
		}
		n, err = c.Read(buf)
	}()
	select {
	case <-done:
		if err != nil {
			return nil, err
		}
		return buf[:n], nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func runTest(ctx context.Context, conn net.Conn) error {
	date := fmt.Sprintf("Date: %s\r\n", time.Now().Format(time.RFC1123Z))
	mail := (date + "Subject: Hello, World!\r\nFrom: no-reply@test.burgerdev.de\r\n\r\nHi there!\r\n.\r\n")
	msg := []string{
		"",
		"HELO test.burgerdev.de\r\n",
		"MAIL FROM: no-reply@test.bur.invalid\r\n",
		"RCPT TO: webmaster@burgerdev.de\r\n",
		"DATA\r\n",
		mail,
		"QUIT\r\n",
	}

	for _, m := range msg {
		fmt.Printf("client> %s", m)
		if m == "" {
			fmt.Printf("\n")
		}
		ans, err := Exchange(ctx, conn, []byte(m))
		if err != nil {
			return err
		}
		fmt.Printf("server> %s", ans)
		resp, err := Parse(ans)
		if err != nil {
			return err
		}
		if !resp.IsOK() {
			return fmt.Errorf("%+v", *resp)
		}
	}
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	address := os.Getenv("SMTP_SERVER")
	if address == "" {
		address = "mail.burgerdev.de:smtp"
	}
	dialer := &net.Dialer{}
	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}

	err = runTest(ctx, conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
}
