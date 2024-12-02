(base) zhouzejun@zhouzejuns-MacBook-Pro-2 ~ % curl http://localhost:8080/r/benchmark/user
{"correct":false}
(base) zhouzejun@zhouzejuns-MacBook-Pro-2 ~ % curl http://localhost:8080/r/benchmark/usershared
{"correct":false}
(base) zhouzejun@zhouzejuns-MacBook-Pro-2 ~ % curl http://localhost:8080/r/benchmark/user
{"correct":false}
(base) zhouzejun@zhouzejuns-MacBook-Pro-2 ~ % wrk -t2 -c3 -d30s http://localhost:8080/r/benchmark/user     
Running 30s test @ http://localhost:8080/r/benchmark/user
  2 threads and 3 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.10s    77.03ms   1.33s    68.52%
    Req/Sec     0.04      0.19     1.00     96.30%
  54 requests in 30.06s, 7.17KB read
Requests/sec:      1.80
Transfer/sec:     244.34B
(base) zhouzejun@zhouzejuns-MacBook-Pro-2 ~ % wrk -t2 -c3 -d30s http://localhost:8080/r/benchmark/usershared
Running 30s test @ http://localhost:8080/r/benchmark/usershared
  2 threads and 3 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.12s    81.22ms   1.34s    67.92%
    Req/Sec     0.02      0.14     1.00     98.11%
  53 requests in 30.07s, 9.12KB read
Requests/sec:      1.76
Transfer/sec:     310.64B
(base) zhouzejun@zhouzejuns-MacBook-Pro-2 ~ % wrk -t2 -c3 -d30s http://localhost:8080/r/benchmark/user
Running 30s test @ http://localhost:8080/r/benchmark/user
  2 threads and 3 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.11s    82.83ms   1.29s    72.22%
    Req/Sec     0.07      0.26     1.00     92.59%
  54 requests in 30.06s, 7.17KB read
Requests/sec:      1.80
Transfer/sec:     244.31B
(base) zhouzejun@zhouzejuns-MacBook-Pro-2 ~ % wrk -t2 -c3 -d30s http://localhost:8080/r/benchmark/usershared
Running 30s test @ http://localhost:8080/r/benchmark/usershared
  2 threads and 3 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.10s    84.39ms   1.30s    66.67%
    Req/Sec     0.02      0.14     1.00     98.15%
  54 requests in 30.06s, 9.19KB read
Requests/sec:      1.80
Transfer/sec:     313.16B