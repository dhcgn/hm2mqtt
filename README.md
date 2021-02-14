# hm2mqtt

[![CircleCI](https://circleci.com/gh/dhcgn/hm2mqtt.svg?style=svg)](https://circleci.com/gh/dhcgn/hm2mqtt)
![Go](https://github.com/dhcgn/hm2mqtt/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/dhcgn/hm2mqtt)](https://goreportcard.com/report/github.com/dhcgn/hm2mqtt)
[![Maintainability](https://api.codeclimate.com/v1/badges/b5dcdb24ef1e6237d397/maintainability)](https://codeclimate.com/github/dhcgn/hm2mqtt/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/b5dcdb24ef1e6237d397/test_coverage)](https://codeclimate.com/github/dhcgn/hm2mqtt/test_coverage)

:exclamation: Not yet ready for use!

## Status

The prototype can send events and subscribe to commands from a mqtt broker. So I can control all my devices with openHAB. Keep in mind, that a lot comfort feature missing at the moment.

## Introduction

The project enables you to control your Homematic devices with any mqtt compatible system, e.g. openHAB or several apps.

![hm2mqtt Overview HM only](docs/images/hm2mqtt%20-%20Overview%20HM%20only.jpg)

*hm2mqtt* is a native arm application created from Go. Because Go is statically linked all dependencies are included.

With this you can perfectly integrate your Homematic setup with other systems which tend to use mqtt.

### Feature Set

- Route all internal RPC events to a mqtt broker
- Subscribe to commands from a mqtt broker
- Host a website on tte CCU3 for configuration, setup and logging
- Installer for the CCU3 Web Interface for easy setup
- Optional automatic update

### Infrastructure Sample

![hm2mqtt Overview](docs/images/hm2mqtt%20-%20Overview.jpg)

### Trivial openHAP Sample

![Screenshot](https://i.ibb.co/PWphmXK/screenshot.png")

## Getting Started

> This is only for testing purpose, to persist hm2mqtt use the installer.

1. Set up a mqtt Broker e.g. [eclipse-mosquitto](https://registry.hub.docker.com/_/eclipse-mosquitto/)
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

## openHAB

![](https://user-images.githubusercontent.com/6566207/107683288-aba23180-6ca1-11eb-8da5-2bf80df1f850.png)

## MQTT Client

On Windows go to the Store and download [MQTT-Explorer](https://www.microsoft.com/store/productId/9PP8SFM082WD) or [MQTT Box](https://www.microsoft.com/store/productId/9NBLGGH55JZG).

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
- Friendly names mapping for devices  
  /usr/local/etc/config/devies.yaml
- Executable  
  /usr/local/etc/config/addons/www/mqtt/hm2mqtt

## Thanks to XML-API

https://github.com/jens-maus/XML-API/blob/master/update_script
