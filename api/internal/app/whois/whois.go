package whois

import (
	"time"

	whoisparser "github.com/likexian/whois-parser-go"
)

type Whois struct {
	parser Parser
}

func NewWhois(parser Parser) Whois {
	return Whois{
		parser: parser,
	}
}

func (w *Whois) IsAvailable(domain string) (bool, error) {
	t, err := w.parser.ExpirationDate(domain)
	if err == whoisparser.ErrDomainNotFound {
		return true, nil
	}
	if err != nil {
		return false, err
	}

	return t.Before(time.Now()), nil
}
