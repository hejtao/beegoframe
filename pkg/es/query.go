package es

import "beegoframe/pkg/es/internal"

func NewBoolQuery() internal.BoolQuery {
	return internal.NewBoolQuery()
}

// [SELECT * FROM book WHERE title LIKE ? OR author LIKE ?] data
func NewMultiMatchQuery(text interface{}, fields ...string) internal.MultiMatchQuery {
	return internal.NewMultiMatchQuery(text, fields...)
}

// [SELECT * FROM book WHERE title = ?] data
func NewMatchQuery(name string, text interface{}) internal.MatchQuery {
	return internal.NewMatchQuery(name, text)
}

func NewQueryStringQuery(query string) internal.QueryStringQuery {
	return internal.NewQueryStringQuery(query)
}

func NewTermQuery(name string, value interface{}) internal.TermQuery {
	return internal.NewTermQuery(name, value)
}

func NewTermsQuery(name string, value interface{}) internal.TermsQuery {
	return internal.NewTermsQuery(name, value)
}
