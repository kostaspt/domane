package http

import (
	"net/http"
	"regexp"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"

	"github.com/kostaspt/domane/api/internal/app/generator"
)

type SearchRequest struct {
	Query string `json:"query" valid:"required,minstringlength(1)"`
}

func (Handler) DomainsExtensions(ctx echo.Context) error {
	req := SearchRequest{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	results := make(generator.Results, 0, len(generator.TopExtensions()))

	// Add common extensions, e.g. foo.com
	for _, r := range generator.CommonExtensions(req.Query) {
		results = append(results, r)
	}

	// If there is a space add common extensions with a hyphen, e.g. foo-bar.com
	if regexp.MustCompile(`\W+`).MatchString(req.Query) {
		for _, r := range generator.CommonExtensionsHyphenated(req.Query) {
			results = append(results, r)
		}
	}

	// Add short extensions e.g. foo.bar
	for _, r := range generator.ShortExtensions(req.Query) {
		results = append(results, r)
	}

	sortResults(results)

	return ctx.JSON(http.StatusOK, results)
}

func (Handler) DomainsSimilars(ctx echo.Context) error {
	req := SearchRequest{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	results := generator.Synonyms(req.Query)

	sortResults(results)

	return ctx.JSON(http.StatusOK, results)
}
