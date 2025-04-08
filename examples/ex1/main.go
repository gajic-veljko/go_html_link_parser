package main

import (
	"fmt"
	"strings"

	"github.com/gajic-veljko/go_html_link_parser"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
  <a href="/AAAAAAAA">Dojavola ova 
  
  sudbina</a>
  <a href="/ODAKLE SAM JA">Svadbe		 su sa okidacima, sahrane sa trubacima</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
