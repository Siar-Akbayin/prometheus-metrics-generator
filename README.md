# prometheus-metrics-generator

This is a simple metrics generator for Prometheus. It generates 3000 different Gauge metrics with 1 dimension to 3000 dimensions.
One sample metric could look like this:
```
sample_gauge_4{dim1="1", dim2="2", dim3="3", dim4="4"} 43534.123 
```
The dimensions follow the pattern `dimX="X"` where X is the dimension number and the value of the metric is random.

The generated metrics are exposed on the `/metrics` endpoint on port 8082.

The purpose of this project is to provide metrics for this Prometheus project that measures the performance of Prometheus
by querying a large number of metrics with different dimension numbers: https://github.com/Siar-Akbayin/prometheus-benchmarking  

It can be tested by running the metrics_generator.go script locally  
or by running:
```
docker build -t metrics-generator .
docker run -p 8082:8082 metrics-generator
``` 
and visiting `http://localhost:8082/metrics` in a browser.