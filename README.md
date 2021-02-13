# hm2mqtt

[![CircleCI](https://circleci.com/gh/dhcgn/hm2mqtt.svg?style=svg)](https://circleci.com/gh/dhcgn/hm2mqtt)
[![Go Report Card](https://goreportcard.com/badge/github.com/dhcgn/hm2mqtt)](https://goreportcard.com/report/github.com/dhcgn/hm2mqtt)
[![Maintainability](https://api.codeclimate.com/v1/badges/b5dcdb24ef1e6237d397/maintainability)](https://codeclimate.com/github/dhcgn/hm2mqtt/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/b5dcdb24ef1e6237d397/test_coverage)](https://codeclimate.com/github/dhcgn/hm2mqtt/test_coverage)

:exclamation: Not yet ready for use!

## Status

The prototyp can send events and subsribe to commands. So I can control all devices with openHAB. But a lot comfort feature missing at the moment. 

## Introduction

The project enables you to control your HomeMatic devices with any mqtt compatible system, e.g. openHAB or serveral apps.

![hm2mqtt Overview HM only](docs/images/hm2mqtt%20-%20Overview%20HM%20only.jpg)

*hm2mqtt* is a nativ arm application created from Go. Because Go is statically linked all dependecies are included.

With this you can perfectly integrate your homematik setup with other systems which tend to use mqtt.

### Feature Set

- Route all internal RPC events to a mqtt broker
- Subsribe to commands from a mqtt broker
- Host a website on te CCU3 for configuration
- Installer for the CCU3 Web Interface for easy setup
- Optional automatic update

### Infrastruture Sample

![hm2mqtt Overview](docs/images/hm2mqtt%20-%20Overview.jpg)

### Trival openHAP Sample

![Screenshot](https://i.ibb.co/PWphmXK/screenshot.png")

## Getting Started

1. Set up an Broker e.g. [eclipse-mosquitto](https://registry.hub.docker.com/_/eclipse-mosquitto/)
1. Set up [openHAB docker](https://registry.hub.docker.com/r/openhab/openhab)
1. Copy arm executable to `/tmp/`
1. Set up configuration
1. Run executable in background

## Parameter

```plain
 Usage of hm2mqtt:
  -config string
        Overrides the path to the config file
```

## Configuration

```json
{
    "ListenerPort": 8777,
    "InterfaceId": 1,
    "HomematicUrl": "http://localhost:2001/",
    "BrokerUrl": "tcp://192.168.1.100:1883"
}
```

## openHAP

![](https://user-images.githubusercontent.com/6566207/107683288-aba23180-6ca1-11eb-8da5-2bf80df1f850.png)

## MQTT Client

![image](https://user-images.githubusercontent.com/6566207/106018455-d1381400-60c1-11eb-8201-16bcfb69bdab.png)


## Running

> Add to `rc script` soon

/usr/local/etc/config/rc.d/mqtt

or

```bash
nohup /tmp/GoHomeMaticMqtt_linux_arm -config config.json > /tmp/mqtt.log &
```

```
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/FAULT_REPORTING with value:      0        1.050519ms
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/BATTERY_STATE with value:        2.900000 766.405µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/VALVE_STATE with value:          0        835.728µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/BOOST_STATE with value:          0        862.655µs
2020/01/13 22:22:12 OK:    topic: hm/KEQ0000000:4/ACTUAL_TEMPERATURE with value:  23.700000 843.123µs
```

## Files

- Config  
  /usr/local/etc/config/mqtt.json
- Executable  
  /usr/local/etc/config/addons/www/mqtt/GoHomeMaticMqtt_linux_arm

## Thanks to XML-API

https://github.com/jens-maus/XML-API/blob/master/update_script
