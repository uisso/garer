Garer IOT Project
==========

[![Go 1.5](https://img.shields.io/badge/go-1.5-blue.svg)]() or [![Go 1.6](https://img.shields.io/badge/go-1.6-green.svg)]()


 This project was made to open/close the my garage door by Raspberry Pi. (DIY gate remote control)

 I play with docker, html, golang and gobot.

 You can be patient, this first project in golang :P

## Requirements (playing with)

- [Golang](http://golang.org/)
- [Docker](http://docker.com) (for ops only, soon)
- [GOBOT](http://gobot.io) nice golang framework to iot
- [Hypriot](http://blog.hypriot.com) (for ops only, soon)


## Orientations:
 * https://github.com/hybridgroup/gobot/tree/master/platforms/raspi
    * plug GPIO 7, pin 26

 * Install pi-blaster (important)
    * https://github.com/sarfata/pi-blaster

### Go Build Environment

 * Build

```sh
make build
```

### Docker Build

 * Soon


### Usage
 * put your binary `garer` into your raspberry
 * execute `./garer`
 * browser it: http://`rasp_host`:8080/garer

### Hardware

<img src="https://raw.githubusercontent.com/uisso/garer/master/doc/hardware.jpg" width="50%">