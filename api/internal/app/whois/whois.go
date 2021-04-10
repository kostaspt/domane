package whois

import (
	"errors"
	"time"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"github.com/rs/zerolog/log"
)

func Get(domain string) (wd Info, err error) {
	raw, err := whois.Whois(domain)
	if err != nil {
		return
	}

	wd.SetRaw(&raw)

	p, err := whoisparser.Parse(raw)
	if errors.Is(err, whoisparser.ErrDomainNotFound) {
		wd.SetIsAvailable(true)
		return wd, nil
	} else if err != nil {
		return
	}

	expiresAt, err := time.Parse(time.RFC3339, p.Domain.ExpirationDate)
	if err != nil {
		log.Err(err).Send()
	}

	wd.SetIsAvailable(time.Now().After(expiresAt))
	wd.ExpiresAt = &expiresAt

	return
}
