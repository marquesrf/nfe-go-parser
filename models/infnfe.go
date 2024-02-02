package models

type InfNFE struct {
	Ide   Ide   `xml:"ide"`
	Emit  Emit  `xml:"emit"`
	Det   Det   `xml:"det"`
	Total Total `xml:"total"`
}
