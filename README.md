# ipdb
Get the IP location and info. Ten million QPS.


## Start the service
just run:
``` shell
~: go build
~: ipdb git:(main) ✗ ./ipdb
go http  :10000
```

## Command-line arguments
``` shell
~: ./ipdb -a
Usage of ./ipdb:
  -addr string
    	http listen addr. (default ":10000")
  -ipfile string
    	ip databases file. (default "./ipip.txtx")
  -logdir string
    	log dir. (default "./logs/")
  -loglevel string
    	log level. (default "debug")
```

## API

1. You can take no arguments, and get the IP information of your current machine

```
~:curl http://ip.fuzu.pro
```

The response is JSON, like this:
```
{
  "code": 0,
  "data": {
    "ip": "223.73.123.17",
    "country": "中国",
    "province": "广东",
    "isp": "移动",
    "latitude": 23.12911,
    "longitude": 113.264385
  }
}
```

2. You can also query the specified IP address:

```
curl 'http://ip.fuzu.pro?ip=1.1.1.1'
```

3. Or, Here is a very simple Web page: [http://ip.fuzu.pro](http://ip.fuzu.pro/)
