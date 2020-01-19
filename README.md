# GoHomeMaticMqttPlugin

[![CircleCI](https://circleci.com/gh/dhcgn/GoHomeMaticMqttPlugin.svg?style=svg)](https://circleci.com/gh/dhcgn/GoHomeMaticMqttPlugin)
[![Go Report Card](https://goreportcard.com/badge/github.com/dhcgn/GoHomeMaticMqttPlugin)](https://goreportcard.com/report/github.com/dhcgn/GoHomeMaticMqttPlugin)
[![Maintainability](https://api.codeclimate.com/v1/badges/b5dcdb24ef1e6237d397/maintainability)](https://codeclimate.com/github/dhcgn/GoHomeMaticMqttPlugin/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/b5dcdb24ef1e6237d397/test_coverage)](https://codeclimate.com/github/dhcgn/GoHomeMaticMqttPlugin/test_coverage)

:exclamation: Not yet ready for use!

## Introduction

With this plugin you can use MQTT to watch and control the HomeMatic CCU3.
And with openHAP you can use **Google Assistent** to control the HomeMatic smart devices.

![Screenshot](https://i.ibb.co/PWphmXK/screenshot.png")

## Getting Started

1. Set up an Broker e.g. [eclipse-mosquitto](https://registry.hub.docker.com/_/eclipse-mosquitto/)
1. Set up [openHAB](https://registry.hub.docker.com/r/openhab/openhab)
1. Copy arm executable to `/tmp/`
1. Set up configuration
1. Run executable in background

## Configuration

```json
{
    "ListenerPort": 8777,
    "InterfaceId": 1,
    "HomematicUrl": "http://localhost:2001/",
    "BrokerUrl": "tcp://192.168.1.100:1883"
}
```

## Running

```bash
nohup /tmp/GoHomeMaticMqtt_linux_arm > /tmp/mqtt.log &
```

```
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/FAULT_REPORTING with value:      0        1.050519ms
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/BATTERY_STATE with value:        2.900000 766.405µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/VALVE_STATE with value:          0        835.728µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/BOOST_STATE with value:          0        862.655µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/ACTUAL_TEMPERATURE with value:  23.700000 843.123µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/SET_TEMPERATURE with value:     18.000000 784.269µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/PARTY_TEMPERATURE with value:    5.000000 856.509µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/PARTY_START_TIME with value:     0        665.051µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/PARTY_START_DAY with value:      1        643.958µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/PARTY_START_MONTH with value:    1        670.78µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/PARTY_START_YEAR with value:     0        653.801µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/PARTY_STOP_TIME with value:      0        667.134µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/PARTY_STOP_DAY with value:       1        656.353µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/PARTY_STOP_MONTH with value:     1        644.113µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/PARTY_STOP_YEAR with value:      0        655.832µs
```
