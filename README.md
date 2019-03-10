# go-tcp-reset
TCP reset attack implementation written in Go for Linux systems. 

This attack was discovered many years ago and it's public domain, however it's responsibility of the person who uses this script to know that interfering with other people connections without their consent is unethical, this is just for research and learning purposes.

The TCP/IP protocol has a special reset packet that when is sent in the middle of any connection, knowing the sequence number it terminates immediately the connection. Based on window size it's possible to make a guessing brute force attack for sequence numbers so the reset packet is treated as a legit packet between the connection nodes.

This script, based on Sam Heybey's TCP-Reset python implementation (link below) makes a TCP-Reset attack based on the real traffic, packets are sniffed from a network interface, treated and sent back with the proper modifications to be used as reset. This of course in a local environment without any special treatment will only monitor and terminate your own connections. I don't encourage using this script for messing with other people's network activity so I won't help anyone interested into weaponizing this script.

Any doubts please contact me: kirebyte@gmail.com

## Dependencies

* pcap

* Go

  ```bash
  sudo apt-get install libpcap-dev
  ```

## Building

### Make sure the script fits your network interface, this is inside main.go file

```go
	// Open connection
	handle, err := pcap.OpenLive(
		"eth0", // network device, CHANGE THIS!
		int32(65535),
		false,
		time.Millisecond,
```

Get all the go dependencies and build.

```bash
go get -d -v
```

## Usage

```bash
sudo ./tcp-reset -S <IP-ADDRESS>
```

## Acknowledgments

- Sam Heybey's TCP-Reset python implementation <https://github.com/sheybey/hoyahacks-tcp-reset>