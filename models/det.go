package models

type Det struct {
	NItem string `xml:"nItem,attr"`
	Prod  []Prod `xml:"prod"`
}
