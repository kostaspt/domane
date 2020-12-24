package whois

import (
	"time"

	"github.com/likexian/whois-go"
	"github.com/likexian/whois-parser-go"
)

type Parser interface {
	Parse(domain string) (info whoisparser.WhoisInfo, err error)
	ExpirationDate(domain string) (t time.Time, err error)
}

type ParserImpl struct{}

func NewParser() ParserImpl {
	return ParserImpl{}
}

func (ParserImpl) Parse(domain string) (info whoisparser.WhoisInfo, err error) {
	response, err := whois.Whois(domain)
	if err != nil {
		return
	}

	return whoisparser.Parse(response)
}

func (p ParserImpl) ExpirationDate(domain string) (t time.Time, err error) {
	w, err := p.Parse(domain)
	if err != nil {
		return
	}

	return time.Parse(time.RFC3339, w.Domain.ExpirationDate)
}
