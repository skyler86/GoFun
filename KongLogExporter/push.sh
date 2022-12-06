#!/bin/bash
image="kong_log_exporter"
now=$(date +%Y%m%d%H%M)
docker build -t harborsy.lenovo.com.cn/public-service/$image:$now /apps/kongLogExporter/
docker push harborsy.lenovo.com.cn/public-service/$image:$now
docker rmi harborsy.lenovo.com.cn/public-service/$image:$now