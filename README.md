# ohnamil - 오늘 남은 일정

Google Calendar 의 내 남은 일정을 PNG파일이나, E-paper에 표시

![ohilnam-epaper](_img/ohilnam.jpg)

## Install requirements

    GOPRIVATE=github.kakaoenterprise.in/IoTEngine go get

## Deploy binary and service script

    ./deploy {IP_OF_RPI}

## install to RPi

Enable SPI from;

    sudo raspi-config

Install ther service

    sudo ln -s $(pwd)/ohnamil.service /lib/systemd/system/
    sudo systemctl enable ohnamil.service
