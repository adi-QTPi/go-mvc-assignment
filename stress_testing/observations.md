# Stress testing

## Get /api/item
```bash
ab -C jwt_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOnsidXNlcl9pZCI6IjE4ZTIyMWM0LTdiMmYtMTFmMC04ZjkzLTJmZDI5MTVlMjdhNiIsInVzZXJfbmFtZSI6InNhbnRvb3IiLCJuYW1lIjoiQ3VzdG9tZXJTaHJlZSIsInJvbGUiOiJjdXN0b21lciJ9LCJleHAiOjE3NTU1NTE3Njd9.bQxsWhzNn9wz_4qVeRC3URxcY_3Rj6FbqqsVBi9-FTI -n 100000 -c 1000 localhost:9005/api/item
```
```bash
This is ApacheBench, Version 2.3 <$Revision: 1913912 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            9005

Document Path:          /api/item
Document Length:        14938 bytes

Concurrency Level:      1000
Time taken for tests:   4.330 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1502600000 bytes
HTML transferred:       1493800000 bytes
Requests per second:    23095.88 [#/sec] (mean)
Time per request:       43.298 [ms] (mean)
Time per request:       0.043 [ms] (mean, across all concurrent requests)
Transfer rate:          338904.94 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   19  90.8      5    4006
Processing:     0    6   4.5      6     106
Waiting:        0    6   4.4      5     106
Total:          1   25  91.0     11    4012

Percentage of the requests served within a certain time (ms)
  50%     11
  66%     12
  75%     12
  80%     13
  90%     41
  95%     71
  98%    132
  99%    223
 100%   4012 (longest request)
```

## Get Menu /static/menu
- HTML page served
- enter jwt token manually.
```bash
ab -C jwt_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOnsidXNlcl9pZCI6IjE4ZTIyMWM0LTdiMmYtMTFmMC04ZjkzLTJmZDI5MTVlMjdhNiIsInVzZXJfbmFtZSI6InNhbnRvb3IiLCJuYW1lIjoiQ3VzdG9tZXJTaHJlZSIsInJvbGUiOiJjdXN0b21lciJ9LCJleHAiOjE3NTU1NTE3Njd9.bQxsWhzNn9wz_4qVeRC3URxcY_3Rj6FbqqsVBi9-FTI -n 100000 -c 1000 localhost:9005/static/menu
```
```bash
This is ApacheBench, Version 2.3 <$Revision: 1913912 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            9005

Document Path:          /static/menu
Document Length:        190204 bytes

Concurrency Level:      1000
Time taken for tests:   22.882 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      19051400000 bytes
HTML transferred:       19020400000 bytes
Requests per second:    4370.16 [#/sec] (mean)
Time per request:       228.825 [ms] (mean)
Time per request:       0.229 [ms] (mean, across all concurrent requests)
Transfer rate:          813062.73 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   2.3      0      33
Processing:     2  227 222.1    160    2714
Waiting:        1  225 222.1    158    2713
Total:          2  227 222.1    160    2715

Percentage of the requests served within a certain time (ms)
  50%    160
  66%    247
  75%    315
  80%    364
  90%    515
  95%    666
  98%    872
  99%   1034
 100%   2715 (longest request)
```
<br />
<br />

## POST Order /api/order
### c , n := 1000, 10000
```bash
ab -n 10000 -c 1000 \
  -p payload/demo_order.json \
  -T application/json \
  -H "Cookie: ENTER_COOKIE_TEXT_HERE" \
  http://localhost:9005/api/order
This is ApacheBench, Version 2.3 <$Revision: 1913912 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            9005

Document Path:          /api/order
Document Length:        59 bytes

Concurrency Level:      1000
Time taken for tests:   15.834 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      12980928 bytes
Total body sent:        17360000
HTML transferred:       590000 bytes
Requests per second:    631.56 [#/sec] (mean)
Time per request:       1583.372 [ms] (mean)
Time per request:       1.583 [ms] (mean, across all concurrent requests)
Transfer rate:          800.61 [Kbytes/sec] received
                        1070.70 kb/s sent
                        1871.31 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   8.2      0      62
Processing:    27 1559 1119.6   1308    8096
Waiting:        5 1559 1119.6   1308    8096
Total:         31 1561 1117.1   1309    8096

Percentage of the requests served within a certain time (ms)
  50%   1309
  66%   1799
  75%   2154
  80%   2406
  90%   3088
  95%   3716
  98%   4565
  99%   5057
 100%   8096 (longest request)
```

### c , n := 1000 , 100000
- had to abort midway...
```bash
This is ApacheBench, Version 2.3 <$Revision: 1913912 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
^C

Server Software:        
Server Hostname:        localhost
Server Port:            9001

Document Path:          /api/order
Document Length:        64 bytes

Concurrency Level:      1000
Time taken for tests:   288.812 seconds
Complete requests:      25655
Failed requests:        0
Total transferred:      18548565 bytes
Total body sent:        30386700
HTML transferred:       1641920 bytes
Requests per second:    88.83 [#/sec] (mean)
Time per request:       11257.536 [ms] (mean)
Time per request:       11.258 [ms] (mean, across all concurrent requests)
Transfer rate:          62.72 [Kbytes/sec] received
                        102.75 kb/s sent
                        165.47 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   3.8      0      34
Processing:    43 10798 7167.0   8920   58834
Waiting:       25 10798 7167.0   8920   58834
Total:         43 10799 7166.5   8920   58835

Percentage of the requests served within a certain time (ms)
  50%   8920
  66%  11932
  75%  14232
  80%  15874
  90%  20456
  95%  24975
  98%  30434
  99%  34459
 100%  58835 (longest request)
```

### Placing order using curl with json payload
- `Step 1 :` login usingcurl and store the cookies in a file.
```bash
curl -c cookies.txt -X POST "http://localhost:9001/account/login" \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "user_name=santoor&password=santoor"
```
- `Step 2 :` use curl to access the api endpoint along with the cookies and payload file.
```bash
curl -b cookies.txt -X POST "http://localhost:9001/api/order" \
  -H "Content-Type: application/json" \
  -d @demo_order.json
```