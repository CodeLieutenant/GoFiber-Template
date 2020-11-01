package email

import (
	"context"
	"crypto/tls"
	"io"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
	"github.com/rs/zerolog"
)

type (
	Config struct {
		Addr     string
		From     string
		Auth     smtp.Auth
		TLS      *tls.Config
		Logger   zerolog.Logger
		PoolSize int
		Senders  int
	}

	smtpService struct {
		ctx     context.Context
		cancel  context.CancelFunc
		from    string
		pool    *email.Pool
		emailCh chan *email.Email
	}

	Interface interface {
		io.Closer
		NewEmail() *email.Email
		Send(to []string, subject string, body []byte) error
		SendEmail(*email.Email) error
	}
)

func NewEmailService(config Config) (Interface, error) {
	if config.Senders == 0 {
		config.Senders = 1
	}

	if config.PoolSize == 0 {
		config.PoolSize = 4
	}

	p, err := email.NewPool(config.Addr, config.PoolSize, config.Auth)

	if err != nil {
		return nil, err
	}

	ch := make(chan *email.Email, 100)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case e := <-ch:
				if err := p.Send(e, time.Duration(5)*time.Minute); err != nil {
					config.Logger.Error().Err(err).Msg("Error while sending email\n")
				}
			}
		}
	}()

	return smtpService{
		ctx:     ctx,
		cancel:  cancel,
		from:    config.From,
		pool:    p,
		emailCh: ch,
	}, nil
}

func (s smtpService) NewEmail() *email.Email {
	return email.NewEmail()
}

func (s smtpService) Send(to []string, subject string, body []byte) error {
	e := &email.Email{
		To:      to,
		Subject: subject,
		HTML:    body,
		From:    s.from,
	}

	s.emailCh <- e
	return nil
}

func (s smtpService) SendEmail(e *email.Email) error {
	s.emailCh <- e
	return nil
}

func (s smtpService) Close() error {
	s.cancel()
	close(s.emailCh)
	s.pool.Close()
	return nil
}
