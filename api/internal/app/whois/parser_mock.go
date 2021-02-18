package whois

import (
	"time"

	whoisparser "github.com/likexian/whois-parser-go"
)

type MockParser struct{}

func NewMockParser() MockParser {
	return MockParser{}
}

func (MockParser) Parse(domain string) (info whoisparser.WhoisInfo, err error) {
	panic("implement me")
}

func (MockParser) ExpirationDate(domain string) (t time.Time, err error) {
	if domain == "available.com" {
		return time.Now().AddDate(0, -1, 0), nil
	}
	if domain == "not-available.com" {
		return time.Now().AddDate(0, 1, 0), nil
	}
	return
}
