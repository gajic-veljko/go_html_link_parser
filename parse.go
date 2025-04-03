package link

import "io"

// Link represents a link(<a herf="...") in an HTML
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return a
// slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
