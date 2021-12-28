package internal

import "github.com/olivere/elastic/v7"

type Query interface {
	elastic.Query
}

type BoolQuery struct {
	*elastic.BoolQuery
}

type MatchQuery struct {
	*elastic.MatchQuery
}

type MultiMatchQuery struct {
	*elastic.MultiMatchQuery
}

type QueryStringQuery struct {
	*elastic.QueryStringQuery
}

type TermQuery struct {
	*elastic.TermQuery
}

type TermsQuery struct {
	*elastic.TermsQuery
}

func NewBoolQuery() BoolQuery {
	return BoolQuery{elastic.NewBoolQuery()}
}

func NewMultiMatchQuery(text interface{}, fields ...string) MultiMatchQuery {
	return MultiMatchQuery{elastic.NewMultiMatchQuery(text, fields...)}
}

func NewMatchQuery(name string, text interface{}) MatchQuery {
	return MatchQuery{elastic.NewMatchQuery(name, text)}
}

func NewQueryStringQuery(query string) QueryStringQuery {
	return QueryStringQuery{elastic.NewQueryStringQuery(query)}
}

func NewTermQuery(name string, value interface{}) TermQuery {
	return TermQuery{elastic.NewTermQuery(name, value)}
}

func NewTermsQuery(name string, value interface{}) TermsQuery {
	return TermsQuery{elastic.NewTermsQuery(name, value)}
}
