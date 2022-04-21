
<h2 align="center">
<img src="https://monibuca.com/img/logo.089ef700.png"></h2>
## Stargazers over time

# Introduction

🧩 Monibuca is a Modularized, Extensible framework for building Streaming Server. 
- Customize the server by combining function plug-ins. 
- It's easy to develop plug-ins to implement business logic. 
- Reduce enterprise development cost and improve development efficiency

# Quick start

## Go has not been installed
```
bash <(curl -s -S -L https://monibuca.com/go.sh)
```
## Go is already installed

1. git clone https://github.com/lyq183/monibuca.git
2. go build && ./monibuca
3. open your browser http://localhost:8080
4. use ffmpeg or OBS to push video streaming to rtmp://localhost/live/user1

# Ecosystem

go to 
[https://plugins.monibuca.com](https://plugins.monibuca.com).
to submit your own plugin

| Project | Description  |
|---------| -------------|
|[plugin-rtmp]|rtmp protocol support.push rtmp stream to monibuca.play stream from monibuca.
|[plugin-rtsp]|rtsp protocol support.pull/push rtsp stream to monibuca
|[plugin-hls]|pull hls stream to monibuca
|[plugin-ts]|used by plugin-hls. read ts file to publish
|[plugin-hdl]|http-flv protocol support. pull http-flv stream from monibuca
|[plugin-gateway]|a console and dashboard to display information and status of monibuca ,also can display UI of other plugins 
|[plugin-record]|record multimedia stream to flv files
|[plugin-cluster]|cascade transmission of multimedia by cluster network
|[plugin-jesscia]|play multimedia stream through websocket protocol
|[plugin-logrotate]|split log files by date or size
|[plugin-rtp]|used by plugin-webrtc and plugin-rtsp
|[plugin-webrtc]|webrtc protocol support. push webrtc stream to monibuca or pull webrtc stream from monibuca
|[plugin-gb28181]|gb28181 protocol support.

[plugin-rtmp]: https://github.com/Monibuca/plugin-rtmp
[plugin-rtsp]: https://github.com/Monibuca/plugin-rtsp
[plugin-hls]:https://github.com/Monibuca/hlspplugin
[plugin-ts]:https://github.com/Monibuca/tspplugin
[plugin-hdl]:https://github.com/Monibuca/plugin-hdl
[plugin-gateway]:https://github.com/Monibuca/plugin-gateway
[plugin-record]:https://github.com/Monibuca/plugin-record
[plugin-cluster]:https://github.com/Monibuca/plugin-cluster
[plugin-jesscia]:https://github.com/Monibuca/plugin-jesscia
[plugin-logrotate]:https://github.com/Monibuca/plugin-logrotate
[plugin-rtp]:https://github.com/Monibuca/plugin-rtp
[plugin-webrtc]:https://github.com/Monibuca/plugin-webrtc
[plugin-gb28181]:https://github.com/Monibuca/plugin-gb28181
# Protocol Functions
| Protocol | Pusher（push）-->Monibuca  |Source-->Monibuca（pull）|Monibuca-->Player（pull）|Monibuca（push）-->Other Server
|---------| -------------|-------------| -------------|-------------|
|rtmp|✔||✔|
|rtsp|✔|✔|✔|✔
|http-flv||✔|✔|
|hls||✔|✔|
|ws-flv|||✔|
|webrtc|✔||✔

# Build & Test with docker

> development and testing only: IP and udp ports need to be exposed carefully in production.
```shell
docker build . -f dockerfile -t m7s:3.0
docker run --name m7s -p 1935:1935 -p 8081:8081 -p 8082:8082 -p 554:554 m7s:3.0
```

# Documentation

中文文档：
[http://docs.monibuca.com](http://docs.monibuca.com).


# Q&A

## Q: There are so many streaming server projects in the world，why need to create Monibuca?

A: Monibuca is different from other streaming servers,that it was created for facilitate secondary development.

## Q: Why use golang?

A: Golang is a greate programming language. It is very suited to build streaming server since streaming server is a kind of IO intensive system. Goroutine is good at doing these jobs. Another important reason of using Golang is that people read the source code or doing secondary development easier.

## Q: What does "Monibuca" mean?

A: No special meaning. Just from monica —— a girl name. 
