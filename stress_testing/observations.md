# Stress testing

## Get /api/item
```bash
ab -C jwt_token=TWT_TOKEN -n 100000 -c 1000 localhost:9005/api/item
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
Document Length:        14743 bytes

Concurrency Level:      1000
Time taken for tests:   4.428 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1483100000 bytes
HTML transferred:       1474300000 bytes
Requests per second:    22581.61 [#/sec] (mean)
Time per request:       44.284 [ms] (mean)
Time per request:       0.044 [ms] (mean, across all concurrent requests)
Transfer rate:          327058.40 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   22 114.9      5    3939
Processing:     3    6   3.6      6      57
Waiting:        2    6   3.4      5      57
Total:          7   28 115.1     11    3944

Percentage of the requests served within a certain time (ms)
  50%     11
  66%     11
  75%     11
  80%     12
  90%     41
  95%     70
  98%    134
  99%    341
 100%   3944 (longest request)
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
Time taken for tests:   18.598 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      7300000 bytes
Total body sent:        11610000
HTML transferred:       590000 bytes
Requests per second:    537.69 [#/sec] (mean)
Time per request:       1859.811 [ms] (mean)
Time per request:       1.860 [ms] (mean, across all concurrent requests)
Transfer rate:          383.31 [Kbytes/sec] received
                        609.63 kb/s sent
                        992.94 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   6.7      0      33
Processing:    12 1786 1712.3   1288   15942
Waiting:       12 1786 1712.3   1288   15942
Total:         14 1788 1710.6   1289   15942

Percentage of the requests served within a certain time (ms)
  50%   1289
  66%   1991
  75%   2512
  80%   2899
  90%   4090
  95%   5126
  98%   6670
  99%   7823
 100%  15942 (longest request)
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