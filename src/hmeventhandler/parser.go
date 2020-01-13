package hmeventhandler

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"log"
)

type Event struct {
	MembersInnerXml string `xml:",innerxml"`
	MethodName      string
	SerialNumber    string
	Type            string
	DataValue       string
}

func parseEventMultiCall(content string) []Event {

	reader := bytes.NewReader([]byte(content))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	type People struct {
		Methods []Event `xml:"params>param>value>array>data>value"`
	}

	v := People{Methods: []Event{}}
	if err := decoder.Decode(&v); err != nil {
		log.Fatalf("unable to parse XML '%s'", err)
	}

	fmt.Println("len: ", len(v.Methods))

	for i, _ := range v.Methods {
		// fmt.Println(i, value)

		v.Methods[i].MethodName = extractMethodName(v.Methods[i].MembersInnerXml)
		serialNumber, what := extractData(v.Methods[i].MembersInnerXml)
		v.Methods[i].SerialNumber = serialNumber
		v.Methods[i].Type = what

		v.Methods[i].DataValue = extractDataValue(v.Methods[i].MembersInnerXml)

		// fmt.Println(v.Methods[i].SerialNumber, v.Methods[i].Type, v.Methods[i].DataValue)
	}

	return v.Methods
}

func extractDataValue(innerXml string) (dataValue string) {
	reader := bytes.NewReader([]byte(innerXml))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	type DummyData struct {
		ValueInt4    string `xml:"member>value>array>data>value>i4"`
		ValueDouble  string `xml:"member>value>array>data>value>double"`
		ValueBoolean string `xml:"member>value>array>data>value>boolean"`
	}

	dDummyData := DummyData{}

	if err := decoder.Decode(&dDummyData); err != nil {
		log.Fatalf("unable to parse XML '%s'", err)
	}

	if dDummyData.ValueDouble != "" {
		return dDummyData.ValueDouble
	}

	if dDummyData.ValueInt4 != "" {
		return dDummyData.ValueInt4
	}

	if dDummyData.ValueBoolean != "" {
		return dDummyData.ValueBoolean
	}

	return "unknown"
}

func extractData(innerXml string) (serialNumber, what string) {
	reader := bytes.NewReader([]byte(innerXml))
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

func extractMethodName(innerXml string) string {
	reader := bytes.NewReader([]byte(innerXml))
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
