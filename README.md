# rpi-websrv

Simple small web server, which can serve pages based on ERB style templates (gerb)

## Install:
```
go get github.com/karlseguin/gerb
go install github.com/ci2rpi/rpi-websrv
```

## Config file:
All configuration parameters can be configured using a single json file as below:
The default port is 7777

```json
{
    "Port": 8888,
    "ContentDirectory": "web"
}
```

## Usage:
### start server
```
$GOPATH/bin/rpi-websrv config.json
```

### sporadic failing health check
```
wget -O - http://pi:7777/health
```

### page changing bg color based on hostname 
```
wget -q -O - http://pi:7777/web/colors
```

### page printing just the host name
```
wget -q -O - http://pi:7777/web/hostname
```
