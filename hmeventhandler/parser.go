package hmeventhandler

import (
	"bytes"
	"encoding/xml"
	"log"

	"github.com/dhcgn/hm2mqtt/shared"
	"golang.org/x/net/html/charset"
)

type internalEvent struct {
	MembersInnerXML string `xml:",innerxml"`
}

func parseEventMultiCall(content string) ([]shared.Event, error) {
	reader := bytes.NewReader([]byte(content))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	type xmlStruct struct {
		Methods []internalEvent `xml:"params>param>value>array>data>value"`
	}

	v := xmlStruct{Methods: []internalEvent{}}
	if err := decoder.Decode(&v); err != nil {
		return nil, err
	}

	var events []shared.Event
	for i := range v.Methods {
		serialNumber, what := extractData(v.Methods[i].MembersInnerXML)
		event := shared.Event{
			MethodName:   extractMethodName(v.Methods[i].MembersInnerXML),
			SerialNumber: serialNumber,
			Type:         what,
			DataValue:    extractDataValue(v.Methods[i].MembersInnerXML),
		}
		events = append(events, event)
	}

	return events, nil
}

func extractDataValue(innerXML string) (dataValue string) {
	reader := bytes.NewReader([]byte(innerXML))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	type XMLData struct {
		ValueInt4    string `xml:"member>value>array>data>value>i4"`
		ValueDouble  string `xml:"member>value>array>data>value>double"`
		ValueBoolean string `xml:"member>value>array>data>value>boolean"`
	}

	xmlData := XMLData{}

	if err := decoder.Decode(&xmlData); err != nil {
		log.Fatalf("unable to parse XML '%s'", err)
	}

	if xmlData.ValueDouble != "" {
		return xmlData.ValueDouble
	}

	if xmlData.ValueInt4 != "" {
		return xmlData.ValueInt4
	}

	if xmlData.ValueBoolean != "" {
		return xmlData.ValueBoolean
	}

	return "unknown"
}

func extractData(innerXML string) (serialNumber, what string) {
	reader := bytes.NewReader([]byte(innerXML))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	type DummyData struct {
		Value []string `xml:"member>value>array>data>value"`
	}

	dMethodName := DummyData{}

	if err := decoder.Decode(&dMethodName); err != nil {
		log.Fatalf("unable to parse XML '%s'", err)
	}

	serialNumber = dMethodName.Value[1]
	what = dMethodName.Value[2]

	return serialNumber, what
}

func extractMethodName(innerXML string) string {
	reader := bytes.NewReader([]byte(innerXML))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	type DummyMethodName struct {
		Value []string `xml:"member>value"`
	}

	dMethodName := DummyMethodName{}

	if err := decoder.Decode(&dMethodName); err != nil {
		log.Fatalf("unable to parse XML '%s'", err)
	}

	s := dMethodName.Value[0]
	return s
}
