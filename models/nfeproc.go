package models

type NfeProc struct {
	Xmlns  string `xml:"xmlns,attr"`
	Versao string `xml:"versao,attr"`
	NFe    NFe    `xml:"NFe"`
}
