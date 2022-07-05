# ohnamil - 오늘 남은 일정

카카오워크의 내 남은 일정을 PNG파일이나, E-paper에 표시

![ohilnam-epaper](_img/ohilnam.jpg)

## Install requirements

    GOPRIVATE=github.kakaoenterprise.in/IoTEngine go get

## Deploy binary and service script

    ./deploy {IP_OF_RPI}

## install to RPi

Enable SPI from;

    rasp-config

Install ther service

    sudo ln -s /home/pi/ohnamil.service /lib/systemd/system/
    sudo systemctl enable ohnamil.service

Increase spi buffer size:

    sudo vi /etc/modprobe.d/local.conf
    ...
    options spidev bufsiz=65536