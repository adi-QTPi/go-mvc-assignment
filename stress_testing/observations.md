# Stress testing

## Get /api/item
```bash
ab -C jwt_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOnsidXNlcl9pZCI6IjE4ZTIyMWM0LTdiMmYtMTFmMC04ZjkzLTJmZDI5MTVlMjdhNiIsInVzZXJfbmFtZSI6InNhbnRvb3IiLCJuYW1lIjoiQ3VzdG9tZXJTaHJlZSIsInJvbGUiOiJjdXN0b21lciJ9LCJleHAiOjE3NTU1NDkxMzl9.7LX02C6yCj8yFqcquMVjBo8q7pG095RWhVQiI5tAdT8 -c 1000 -n 100000 localhost:9005/api/item 
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
Time taken for tests:   4.211 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1502600000 bytes
HTML transferred:       1493800000 bytes
Requests per second:    23748.03 [#/sec] (mean)
Time per request:       42.109 [ms] (mean)
Time per request:       0.042 [ms] (mean, across all concurrent requests)
Transfer rate:          348474.56 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   25 116.2      5    3940
Processing:     1    6   3.2      6      62
Waiting:        1    6   3.2      5      62
Total:          2   31 116.3     11    3948

Percentage of the requests served within a certain time (ms)
  50%     11
  66%     12
  75%     13
  80%     18
  90%     42
  95%     74
  98%    220
  99%    344
 100%   3948 (longest request)
```

## Get Menu /static/menu
- HTML page served
- enter jwt token manually.
```bash
ab -C jwt_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOnsidXNlcl9pZCI6IjMyODU0MjBlLTc4MWEtMTFmMC1iOTA0LTZlNTk5ZDE1YjRhMiIsInVzZXJfbmFtZSI6InNhbnRvb3IiLCJuYW1lIjoiYWRpdHlhIiwicm9sZSI6ImN1c3RvbWVyIn0sImV4cCI6MTc1NTM2ODI5N30.lHVqdiy8JmCapDvfN5jdvq6H7CsGru6SYm_Kov5_DGQ -c 1000 -n 100000 localhost:9001/static/menu
```
- Result 

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
Server Port:            9001

Document Path:          /static/menu
Document Length:        187139 bytes

Concurrency Level:      1000
Time taken for tests:   30.573 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      18744900000 bytes
HTML transferred:       18713900000 bytes
Requests per second:    3270.91 [#/sec] (mean)
Time per request:       305.725 [ms] (mean)
Time per request:       0.306 [ms] (mean, across all concurrent requests)
Transfer rate:          598758.36 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   2.7      0      62
Processing:     7  304 208.1    257    2504
Waiting:        6  302 208.1    255    2501
Total:          7  304 208.1    257    2505

Percentage of the requests served within a certain time (ms)
  50%    257
  66%    342
  75%    406
  80%    451
  90%    581
  95%    706
  98%    867
  99%    984
 100%   2505 (longest request)
```
<br />
<br />

## POST Order /api/order
### c , n := 1000, 10000
```bash
ab -n 10000 -c 1000 \
   -p payload/demo_order.json \
   -T application/json \
   -H "Cookie: jwt_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOnsidXNlcl9pZCI6IjMyODU0MjBlLTc4MWEtMTFmMC1iOTA0LTZlNTk5ZDE1YjRhMiIsInVzZXJfbmFtZSI6InNhbnRvb3IiLCJuYW1lIjoiYWRpdHlhIiwicm9sZSI6ImN1c3RvbWVyIn0sImV4cCI6MTc1NTM3MjQ1Mn0.Cplu8SzvEiP-mKpWzOkNpmFtn0-cMy0fdCBUQnUa7JE; foodopia-session=MTc1NTI4NTc1OHxEWDhFQVFMX2dBQUJFQUVRQUFBd180QUFBUVp6ZEhKcGJtY01DQUFHWDJac1lYTm9EbHRkYVc1MFpYSm1ZV05sSUh0OV80RUNBUUxfZ2dBQkVBQUFXXy1DTlFBQkN5cDFkR2xzTGxCdmNIVndfNE1EQVFFRlVHOXdkWEFCXzRRQUFRSUJBMDF6WndFTUFBRUhTWE5GY25KdmNnRUNBQUFBSXYtRUh3RWNVM1ZqWTJWemMyWjFiR3g1SUhCc1lXTmxaQ0JQY21SbGNpQWpPUUE9fNut9PcyPTRpEcVqf8YLiIHOdfinCQtJDDsWqxXJOLq7" \
   http://localhost:9001/api/order
```
### results
```bash
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
Server Port:            9001

Document Path:          /api/order
Document Length:        64 bytes

Concurrency Level:      1000
Time taken for tests:   20.145 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      7230000 bytes
Total body sent:        11400000
HTML transferred:       640000 bytes
Requests per second:    496.41 [#/sec] (mean)
Time per request:       2014.483 [ms] (mean)
Time per request:       2.014 [ms] (mean, across all concurrent requests)
Transfer rate:          350.49 [Kbytes/sec] received
                        552.64 kb/s sent
                        903.13 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    3   9.9      0      73
Processing:    47 1981 1401.3   1625   11567
Waiting:       22 1981 1401.3   1625   11567
Total:         47 1984 1398.8   1627   11568

Percentage of the requests served within a certain time (ms)
  50%   1627
  66%   2245
  75%   2699
  80%   3013
  90%   3871
  95%   4715
  98%   5814
  99%   6505
 100%  11568 (longest request)
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