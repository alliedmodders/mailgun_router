package main

import (
	"fmt"
	"io"
	"log"

	"github.com/emersion/go-smtp"
	"github.com/mailgun/mailgun-go"
)

const BASE_URL string = "https://api.mailgun.net/v3"

type Backend struct {
	config *Config
	info   *log.Logger
	err    *log.Logger
	server *smtp.Server
}
type User struct {
	backend *Backend
}

func (this *Backend) LoginAnonymous() (smtp.User, error) {
	return &User{this}, nil
}

func (this *Backend) Login(username, password string) (smtp.User, error) {
	return nil, fmt.Errorf("SMTP authentication not supported")
}

func (this *User) Logout() error {
	return nil
}

type ReadWrapper struct {
	r io.Reader
}

func (this ReadWrapper) Read(p []byte) (n int, err error) {
	return this.r.Read(p)
}
func (this ReadWrapper) Close() error {
	return nil
}

func (this *User) Send(c *smtp.Conn, from string, to []string, r io.Reader) error {
	backend := this.backend

	mg := mailgun.NewMailgun(
		backend.config.Domain,
		backend.config.ApiKey,
		backend.config.PublicApiKey)
	msg := mg.NewMIMEMessage(ReadWrapper{r}, to...)
	response, id, err := mg.Send(msg)
	if err != nil {
		return err
	}

	c.Logf(backend.server.InfoLog, "Message %s: %s", id, response)
	return nil
}

type SmtpRouter struct {
	server *smtp.Server
}

func NewSmtpRouter(config *Config, infoLogger *log.Logger, errLogger *log.Logger) *smtp.Server {
	backend := &Backend{config, infoLogger, errLogger, nil}
	server := smtp.NewServer(backend)
	server.Addr = config.Address
	server.MaxIdleSeconds = config.MaxIdleSeconds
	server.RequireAuth = false
	server.InfoLog = infoLogger
	server.ErrorLog = errLogger
	backend.server = server
	return server
}
