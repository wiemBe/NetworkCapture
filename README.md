#Network Capture

Network capture is simple program for capturing the interface using go.

##Required

if you don't know your interface name use interface.go program and select which interface do you want to use.

##Usage
First you need to build interface.go and sniff.go then paste your interface name to sniff.go  

```bash
go run -tags interface interface.go
go run -tags sniff sniff.go

