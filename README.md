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
 $ go get github.com/tarm/serial
 ```




## Arguments
#### udpserver.go

| Number | Name | Description |
| ------ | ---- | ----------- |
| `1` | `ip` | `ip of the server`|
| `2` | `port`|`port of the udp socket on server`|


#### client.go

| Name | Description |
| ---- | ----------- |
| `1` | `ip`|`ip of the server`|
| `2` | `port`|`port of the udp socket on server`|
| `3` | `name`|`the name of the serial port`|
| `4` | `baud` | `the baud rate of the serial port `|


## Builiding and Running
If you want to compile the source code into the platform executable then you can simply run these commands
```
$ go build udpserver.go
$ go build client.go
$ ./udpserver 127.0.0.1 8080 
$ ./client 127.0.0.1 8080 ETC 1222
```

Or you can also run it like this

```
$ go run udpserver.go 127.0.0.1 8080
$ go run client.go 127.0.0.1 8080 ETC 1222
```

*If you don't want to set the values in the argument then you can leave it to default values by not giving it in the arguments*


#### Description
The Udp Server and Client will run forever unless there is an error. The client will open a serial port and then it will open the udp socket and after reading from the serial port it will send the json data to the server and it will print the json to the console