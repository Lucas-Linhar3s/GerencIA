package dtos

import (
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type ParamsDTO struct {
	Offset *int    `form:"offset"`
	Limit  *int    `form:"limit"`
	Search *string `form:"search"`
}

// FromValueRequestPagination decodes a JSON request body into a RequestPagination struct
func FromValueParamsDTO(value *http.Request) (*ParamsDTO, error) {
	var (
		limit  int
		offset int
		err    error
	)
	limit, err = strconv.Atoi(value.FormValue("limit"))
	if err != nil {
		limit = 10
	}

	offset, err = strconv.Atoi(value.FormValue("offset"))
	if err != nil {
		offset = 0
	}

	search := strings.TrimSpace(strings.ReplaceAll(value.FormValue("search"), " ", "%"))
	searchWithoutAccents := removeAccents(search)

	return &ParamsDTO{
		Limit:  &limit,
		Offset: &offset,
		Search: &searchWithoutAccents,
	}, nil
}

// removeAccents removes accents from a string using golang.org/x/text/transform
func removeAccents(input string) string {
	// Create a transformer to remove accents
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)

	// Apply the transformation to the string
	result, _, _ := transform.String(t, input)
	return result
}

// isMn is a helper function to check if a rune is a non-spacing mark
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Non-spacing mark
}
