# go-read-serial-port-write-udp



## Files
```
    udpserver.go     This createst the udp server and prints to stdout when receives a data
    client.go        This opens the serial port and forward data to the sudo socket 
    .gitignore
    README.md
```

### Requirements 
- Go

Before building the source code please run this command

 ```
 go get github.com/tarm/serial
 ```




## Arguments
To see the arguments you can type `go run filename.go`
#### udpserver.go

| Name | Description |
| ---- | ----------- |
| `ip` | `ip of the server`|
| `port`|`port of the udp socket on server`|

``` 
$ go run udpserver.go --help

  -ip string
        ip of the server (default "127.0.0.1")
  -port int
        port of the udp socket on server (default 1234)

$ go run udpserver.go -port 8080 -ip 127.0.0.1
```


#### client.go

| Name | Description |
| ---- | ----------- |
| `baud` | `the baud rate of the serial port `|
| `name`|`the name of the serial port`|
| `port`|`port of the udp socket on server`|
| `ip`|`ip of the server`|

``` 
$ go run client.go -ip 127.0.0.1 -port 8080 -name ETC -baud 1222
```


*If you don't want to set the values in the argument then you can leave it to default by not giving it in the arguments flags*


#### Description
The Udp Server and Client will run forever unless there is an error. The client will open a serial port and then it will open the udp socket and after reading from the serial port it will send the json data to the server and it will print the json to the console