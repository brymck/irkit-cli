irkit-cli
=========

Now that IRKit service has been discontinued, this guide assumes you've managed to flash the firmware to no longer require internet connection ([English][english], [Japanese][japanese]). I'd recommend auto-translating the Japanese page and using that.

Usage
-----

```
git clone https://github.com/brymck/irkit-cli
cd irkit-cli
go build .
```

If you're just setting up Wi-Fi connectivity or have reset your IRKit, you can connect to its network (usually named something like `IRKitABCD` with a default password of `XXXXXXXXXX` if you've flashed the firmware and reset) and set the Wi-Fi details with something like

```
./irkit-cli wifi --ssid yournetwork --password yourpassword --wpa
```

Note that this assumes `192.168.1.1` points to your IRKit.

Now that you've reconnected to your usual Wi-Fi network, you can now configure the name of your IRKit (run `dns-sd -B _irkit._tcp`):

```
./irkit-cli config --name 
```

You can listen for infrared with

```
./irkit-cli messages
```

And then you can send the payload with

```
./irkit-cli messages --content <payload>
```

[english]: https://www.adriancourreges.com/blog/2015/02/01/customizing-irkit-firmware-led-and-offline-mode/
[japanese]: https://zenn.dev/kangaechu/articles/irkit_install_custom_firmware