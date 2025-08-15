# Stress testing

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
```bash
ab -n 1000 -c 10 \
   -p payload/demo_order.json \
   -T application/json \
   -H "Cookie: jwt_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOnsidXNlcl9pZCI6IjMyODU0MjBlLTc4MWEtMTFmMC1iOTA0LTZlNTk5ZDE1YjRhMiIsInVzZXJfbmFtZSI6InNhbnRvb3IiLCJuYW1lIjoiYWRpdHlhIiwicm9sZSI6ImN1c3RvbWVyIn0sImV4cCI6MTc1NTM3MjQ1Mn0.Cplu8SzvEiP-mKpWzOkNpmFtn0-cMy0fdCBUQnUa7JE; foodopia-session=MTc1NTI4NTc1OHxEWDhFQVFMX2dBQUJFQUVRQUFBd180QUFBUVp6ZEhKcGJtY01DQUFHWDJac1lYTm9EbHRkYVc1MFpYSm1ZV05sSUh0OV80RUNBUUxfZ2dBQkVBQUFXXy1DTlFBQkN5cDFkR2xzTGxCdmNIVndfNE1EQVFFRlVHOXdkWEFCXzRRQUFRSUJBMDF6WndFTUFBRUhTWE5GY25KdmNnRUNBQUFBSXYtRUh3RWNVM1ZqWTJWemMyWjFiR3g1SUhCc1lXTmxaQ0JQY21SbGNpQWpPUUE9fNut9PcyPTRpEcVqf8YLiIHOdfinCQtJDDsWqxXJOLq7" \
   http://localhost:9001/api/order
```
- result
```bash
This is ApacheBench, Version 2.3 <$Revision: 1913912 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            9001

Document Path:          /api/order
Document Length:        61 bytes

Concurrency Level:      10
Time taken for tests:   1.055 seconds
Complete requests:      1000
Failed requests:        912
   (Connect: 0, Receive: 0, Length: 912, Exceptions: 0)
Total transferred:      716972 bytes
Total body sent:        1140000
HTML transferred:       61924 bytes
Requests per second:    948.14 [#/sec] (mean)
Time per request:       10.547 [ms] (mean)
Time per request:       1.055 [ms] (mean, across all concurrent requests)
Transfer rate:          663.86 [Kbytes/sec] received
                        1055.55 kb/s sent
                        1719.40 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     4   10  10.6      9     114
Waiting:        4   10  10.6      9     114
Total:          4   10  10.6      9     114

Percentage of the requests served within a certain time (ms)
  50%      9
  66%      9
  75%     10
  80%     10
  90%     11
  95%     13
  98%     20
  99%    111
 100%    114 (longest request)
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