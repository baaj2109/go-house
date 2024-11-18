package model

import "github.com/PuerkitoBio/goquery"

// Document is the representation goquery.Document.
type Document struct {
	doc *goquery.Document
}

// NewDocument create a `Document` with default value.
func NewDocument() *Document {
	return &Document{
		doc: &goquery.Document{},
	}
}
