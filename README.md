Mercury
====

What is Mercury
----
Mercury is a room based chat server which is aiming to help deveplors building a chat service in a fast way. Mercury is an independent service from your app which make you can take more concentration to your app. We are also planning to make Mercury to a distributed chat server for a more scalability.

How Mercury Work
----
The mercury is designed into two part, **Http Rest API** server and **Websocket connection** server.

### Rest API
| API | Method | Description |
| ---- | :----: | :----: |
| /api/token | GET  | Get token of a user for a websocket connection |
| /api/room/add  | POST | add members into a chat room |

### Websocket connection
Connect to the server using the token.
```
ws://<ip>:<port>/ws/connect?token=xxxx
```
Merucy server receive a json format message. The client should send a json data to the server in the websocket connection.
#### send message to a chat room
```
{
  "type": 1,
  "rid": "<room id>",
  "text" "<message>",
}
```

#### get history message
```
{
  "type": 2,
  "msgid": "<start message id>",
  "offset" "<number of message before msgid>",
}
```


Server Configuration
---
### command-line flags

#### Server
Basic server configuration

| config | default value |
| ----- | :----: |
| --server.address | 127.0.0.1  |
| --server.port  | 6010 |

#### Log

Logging in mercury is using the interface provided by [go-kit](https://github.com/go-kit/kit/tree/master/log)

| config | available value |
| ---- | :----: |
| --log.format | json(default) \| logfmt   |
| --log.level  | info(default) \| warn \| error \| debug  |

#### Storage
You can choose from multiple back end storage to store the data in mercury.

MySQL

| config | default value |
| ---- | :----: |
| --mysql.host | ""  |
| --mysql.port | ""  |
| --mysql.user | ""   |
| --mysql.password |  ""  |


Simple Demo
----
```
~$ go get github.com/leeif/mercury
~$ cd $GOPATH/src/github.com/leeif/mercury
~$ dep ensure
```
Start mercury server in 127.0.0.1
```
~$ go run ./ --log.level=debug
```

Test client
```
// Terminal 1
~$ go run ./test --member=1

// Terminal 2
~$ go run ./test --member=2
```

Send Message
```
// Terminal 1
send> Hello World 1

// Termianl 2
send> Hello World 2
```
