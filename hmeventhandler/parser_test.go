package hmeventhandler

import (
	"github.com/dhcgn/gohomematicmqttplugin/shared"
	"reflect"
	"testing"
)

func getTestData(testCase string) string {
	xmlIso := `<?xml version="1.0" encoding="iso-8859-1"?>
<methodCall>
    <methodName>system.multicall</methodName>
    <params>
        <param>
            <value>
                <array>
                    <data>
                        <value>
                            <struct>
                                <member>
                                    <name>methodName</name>
                                    <value>event</value>
                                </member>
                                <member>
                                    <name>params</name>
                                    <value>
                                        <array>
                                            <data>
                                                <value>1</value>
                                                <value>NEQ0000000:4</value>
                                                <value>FAULT_REPORTING</value>
                                                <value>
                                                    <i4>0</i4>
                                                </value>
                                            </data>
                                        </array>
                                    </value>
                                </member>
                            </struct>
                        </value>
                        <value>
                            <struct>
                                <member>
                                    <name>methodName</name>
                                    <value>event</value>
                                </member>
                                <member>
                                    <name>params</name>
                                    <value>
                                        <array>
                                            <data>
                                                <value>1</value>
                                                <value>NEQ0000000:4</value>
                                                <value>BATTERY_STATE</value>
                                                <value>
                                                    <double>2.500000</double>
                                                </value>
                                            </data>
                                        </array>
                                    </value>
                                </member>
                            </struct>
                        </value>
                        <value>
                            <struct>
                                <member>
                                    <name>methodName</name>
                                    <value>event</value>
                                </member>
                                <member>
                                    <name>params</name>
                                    <value>
                                        <array>
                                            <data>
                                                <value>1</value>
                                                <value>JEQ000000:0</value>
                                                <value>UNREACH</value>
                                                <value>
                                                    <boolean>1</boolean>
                                                </value>
                                            </data>
                                        </array>
                                    </value>
                                </member>
                            </struct>
                        </value>
                    </data>
                </array>
            </value>
        </param>
    </params>
</methodCall>`

	xml := `<?xml version="1.0" ?>
<methodCall>
    <methodName>system.multicall</methodName>
    <params>
        <param>
            <value>
                <array>
                    <data>
                        <value>
                            <struct>
                                <member>
                                    <name>methodName</name>
                                    <value>event</value>
                                </member>
                                <member>
                                    <name>params</name>
                                    <value>
                                        <array>
                                            <data>
                                                <value>1</value>
                                                <value>NEQ0000000:4</value>
                                                <value>FAULT_REPORTING</value>
                                                <value>
                                                    <i4>0</i4>
                                                </value>
                                            </data>
                                        </array>
                                    </value>
                                </member>
                            </struct>
                        </value>
                        <value>
                            <struct>
                                <member>
                                    <name>methodName</name>
                                    <value>event</value>
                                </member>
                                <member>
                                    <name>params</name>
                                    <value>
                                        <array>
                                            <data>
                                                <value>1</value>
                                                <value>NEQ0000000:4</value>
                                                <value>BATTERY_STATE</value>
                                                <value>
                                                    <double>2.500000</double>
                                                </value>
                                            </data>
                                        </array>
                                    </value>
                                </member>
                            </struct>
                        </value>
                        <value>
                            <struct>
                                <member>
                                    <name>methodName</name>
                                    <value>event</value>
                                </member>
                                <member>
                                    <name>params</name>
                                    <value>
                                        <array>
                                            <data>
                                                <value>1</value>
                                                <value>JEQ000000:0</value>
                                                <value>UNREACH</value>
                                                <value>
                                                    <boolean>1</boolean>
                                                </value>
                                            </data>
                                        </array>
                                    </value>
                                </member>
                            </struct>
                        </value>
                    </data>
                </array>
            </value>
        </param>
    </params>
</methodCall>`

	xmlDefect := `<?xml `

	switch testCase {
	case "iso-8859-1":
		return xmlIso
	case "standard":
		return xml
	case "defect":
		return xmlDefect
	}

	return ""
}

func Test_parseEventMultiCall(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    []shared.Event
		wantErr bool
	}{
		{name: "multiple events", args: struct{ content string }{content: getTestData("standard")}, want: []shared.Event{
			{MethodName: "event", SerialNumber: "NEQ0000000:4", Type: "FAULT_REPORTING", DataValue: "0"},
			{MethodName: "event", SerialNumber: "NEQ0000000:4", Type: "BATTERY_STATE", DataValue: "2.500000"},
			{MethodName: "event", SerialNumber: "JEQ000000:0", Type: "UNREACH", DataValue: "1"},
		}},
		{name: "multiple events iso-8859-1", args: struct{ content string }{content: getTestData("iso-8859-1")}, want: []shared.Event{
			{MethodName: "event", SerialNumber: "NEQ0000000:4", Type: "FAULT_REPORTING", DataValue: "0"},
			{MethodName: "event", SerialNumber: "NEQ0000000:4", Type: "BATTERY_STATE", DataValue: "2.500000"},
			{MethodName: "event", SerialNumber: "JEQ000000:0", Type: "UNREACH", DataValue: "1"},
		}},
		{name: "multiple events iso-8859-1", args: struct{ content string }{content: getTestData("defect")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseEventMultiCall(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseEventMultiCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseEventMultiCall() got = %v, want %v", got, tt.want)
			}
		})
	}
}
