package model

import "encoding/xml"

//Envelope struct
type Envelope struct {
	XMLName xml.Name `xml:"Envelope" json:"-"`
	Cube    CubeMain `xml:"Cube"`
	Xmlns   string   `xml:"xmlns,attr"`
}

//CubeMain struct
type CubeMain struct {
	XMLName xml.Name `xml:"Cube" json:"-"`
	Cube    []Cube   `xml:"Cube"`
}

//Cube struct
type Cube struct {
	XMLName xml.Name   `xml:"Cube" json:"-"`
	Time    string     `xml:"time,attr" json:"time"`
	Cube    []CubeItem `xml:"Cube"`
}

//CubeItem struct
type CubeItem struct {
	XMLName  xml.Name `xml:"Cube" json:"-"`
	Currency string   `xml:"currency,attr" json:"currency,omitempty"`
	Rate     string   `xml:"rate,attr" json:"rate,omitempty"`
}
