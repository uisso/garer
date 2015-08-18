Garer IOT Project
==========
 This project was made to open/close the my home garage gate from Raspberry Pi. (DY gate remote control)

 I play with docker, html, golang and gobot.

 You be patient, this first project in golang :P

Playing with:
 * **Hypriot**: http://blog.hypriot.com/
 * **GOBOT** nice golang framework to iot: http://gobot.io/

* Orientations:
    * https://github.com/hybridgroup/gobot/tree/master/platforms/raspi
    * plug GPIO 7, pin 26

* Install pi-blaster (important)

  https://github.com/sarfata/pi-blaster


### Go Build Environment with Docker

 * https://github.com/hypriot/rpi-golang

```sh
docker run -v /root/gopath:/gopath --name goal -t -i hypriot/rpi-golang bash
```

### Usage

 * browser it: http://localhost:8080/garer

### Hardware

<img src="https://github.com/favicon.ico" width="48">
![alt text](https://github.com/uisso/garer/raw/master/doc/hardware.png "Rasp + relay + gate remote control")
