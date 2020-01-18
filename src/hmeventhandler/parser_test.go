package hmeventhandler

import (
	"reflect"
	"testing"
)

func getTestData() string {
	xml :=  `<?xml version="1.0" encoding="iso-8859-1"?>
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
	return xml
}

func Test_parseEventMultiCall(t *testing.T) {

	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want []Event
	}{
		{name: "multiple events", args: struct{ content string }{content: getTestData()}, want: []Event{
			{MethodName:"event" , SerialNumber: "NEQ0000000:4", Type: "FAULT_REPORTING", DataValue: "0"},
			{MethodName:"event" , SerialNumber: "NEQ0000000:4", Type: "BATTERY_STATE", DataValue: "2.500000"},
			{MethodName:"event" , SerialNumber: "JEQ000000:0", Type: "UNREACH", DataValue: "1"},
		} },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseEventMultiCall(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseEventMultiCall() = %v, want %v", got, tt.want)
			}
		})
	}
}