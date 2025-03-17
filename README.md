# Go-Pro-BLE

### What is this?

This is a Go implementation of the following OpenGoPro Python scripts:
- [connect_ble.py](https://github.com/gopro/OpenGoPro/blob/main/demos/python/tutorial/tutorial_modules/tutorial_1_connect_ble/ble_connect.py)
- [enable_wifi_api.py](https://github.com/gopro/OpenGoPro/blob/main/demos/python/tutorial/tutorial_modules/tutorial_6_connect_wifi/enable_wifi_ap.py)
<!-- - [connect_as_sta.py] (https://github.com/gopro/OpenGoPro/blob/main/demos/python/tutorial/tutorial_modules/tutorial_6_connect_wifi/connect_as_sta.py) -->

The Python implementation uses [Bleak](https://github.com/hbldh/bleak).

This implementation uses [TinyGo](https://tinygo.org/).

![tg](images/tg.png)

### Why?

To connect to your GoPro via BlueTooth in order to and enable it's WIFI AP ( This can only be done programmatically ). 

Once enabled, you can use something like https://github.com/chrisjoyce911/goprowifi to connect to that access point and begin working with your GoPro programmatically.

### Considerations

This was manually tested on a GoPro Hero 12 Black and OSX.

### Instructions

1. On your computer, make sure Bluetooth is enabled.

2. On your GoPro, navigate to your settings and click "Pair Device"

![settings](images/settings.png)
![searching](images/searching.png)

3. On your computer, run the script

        go run examples/enable-wap/main.go

4. Use the credentials in the output to connect to your GoPro

![credentials](images/credentials.png)

![ssid](images/ssid.png)
![connected](images/connected.png)

### How it works

![init](images/init.png)







