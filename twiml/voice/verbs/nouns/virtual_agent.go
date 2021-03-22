package nouns

import "encoding/xml"

type VirtualAgentAttributes struct {
	ConnectorName     string  `xml:"connectorName,attr"`
	Language          *string `xml:"language,attr,omitempty"`
	SentimentAnalysis *bool   `xml:"sentimentAnalysis,attr,omitempty"`
	StatusCallback    *string `xml:"statusCallback,attr,omitempty"`
}

type VirtualAgent struct {
	XMLName xml.Name `xml:"VirtualAgent"`

	VirtualAgentAttributes
}
