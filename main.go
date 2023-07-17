package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/brymck/irkit-cli/wifi"
)

func getMessages() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://irkita2dc.local/messages", nil)
	if err != nil {
		fmt.Println("Error creating GET request:", err)
		return "", err
	}
	req.Header.Set("X-Requested-With", "curl")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Check if status code is success
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Response status code:", resp.StatusCode)
		return "", errors.New("response status code not ok")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}
	return string(body), nil
}

func setWifi(payload string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://192.168.1.1/wifi", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		fmt.Println("Error creating POST request:", err)
		return err
	}
	req.Header.Set("X-Requested-With", "curl")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return err
	}
	defer resp.Body.Close()

	// Check if status code is success
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Response status code:", resp.StatusCode)
		return errors.New("response status code not ok")
	}

	return nil
}

func main() {
	// Define subcommands
	msgCmd := flag.NewFlagSet("messages", flag.ExitOnError)
	wifiCmd := flag.NewFlagSet("wifi", flag.ExitOnError)
	ssidText := wifiCmd.String("ssid", "", "SSID")
	pwdText := wifiCmd.String("password", "", "Password")
	wep := wifiCmd.Bool("wep", false, "WEP security")
	wpa := wifiCmd.Bool("wpa", false, "WPA/WPA2 security")
	verbose := wifiCmd.Bool("verbose", false, "Verbose output")
	help := wifiCmd.Bool("help", false, "Help")

	// Parse command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Subcommand required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "messages":
		msgCmd.Parse(os.Args[2:])
	case "wifi":
		wifiCmd.Parse(os.Args[2:])
	default:
		fmt.Println("Invalid subcommand")
		os.Exit(1)
	}

	if *help {
		fmt.Println("Usage: irkit-cli [options] <subcommand> [subcommand options]")
		fmt.Println("Subcommands:")
		fmt.Println("  messages")
		fmt.Println("  wifi")
		fmt.Println("Options:")
		wifiCmd.PrintDefaults()
		os.Exit(0)
	}

	// Execute subcommand
	if msgCmd.Parsed() {
		body, err := getMessages()
		if err != nil {
			fmt.Println(body)
		}
	} else if wifiCmd.Parsed() {
		if *ssidText == "" {
			fmt.Println("--ssid flag required for wifi subcommand")
			os.Exit(1)
		}
		if *pwdText == "" {
			fmt.Println("--password flag required for wifi subcommand")
			os.Exit(1)
		}
		security := wifi.SECURITY_NONE
		if *wep {
			security = wifi.SECURITY_WEP
		}
		if *wpa {
			security = wifi.SECURITY_WPA_WPA2
		}
		creds := &wifi.Credentials{
			Security:  security,
			SSID:      *ssidText,
			Password:  *pwdText,
			WifiIsSet: true,
		}
		payload := creds.Serialize()
		if *verbose {
			fmt.Println("Submitting payload to IRKit: ", payload)
		}
		setWifi(payload)
	}
}
