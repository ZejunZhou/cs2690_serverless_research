## Get start

Steps to Set Up and Run the Benchmark
1. Clone the Repositories
Clone the benchmark repository:
```
git clone https://github.com/ZejunZhou/Serverless_Comparison_Benchmark.git
```
Clone the serverless functions repository:
```
git clone https://github.com/ZejunZhou/Ironfunctions-ServerlessResearch.git
```
2. Run the Benchmark
Navigate to the benchmark repository's hotelReservation/ directory:

```
cd Serverless_Comparison_Benchmark/hotelReservation/
```
Run the benchmark setup script:

```
./run.sh
```

3. Set Up and Run Serverless Functions
Return to the root directory where you cloned the repositories:

```
cd ../../Ironfunctions-ServerlessResearch/
```

### Build the Docker image:

```
docker build -f Dockerfile.2690 -t 2690 .
```

### Run the Docker container to start serverless functions:

```
docker run -it --name functions -v ${PWD}/data:/app/data -v /var/run/docker.sock:/var/run/docker.sock -p 8080:8080 2690
```

### Register the serverless functions by running the script:

```
./run.sh
```

4. Run the Benchmark Tests
Return to the root directory where you cloned the repositories:

```
cd ../Serverless_Comparison_Benchmark/
```

Run the benchmark directly (without serverless functions):

```
wrk -t2 -c2 -d30s -s mixed-workload_type_1.lua http://localhost:9977
```

Run the benchmark with serverless functions:

```
wrk -t2 -c2 -d30s -s mixed-workload_type_2.lua http://localhost:8080
```
