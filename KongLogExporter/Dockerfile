FROM alpine:latest
RUN mkdir /apps
WORKDIR /apps
ADD ./main /apps/
ADD ./config/json/* /apps/config/json/
RUN chmod 777 /apps/main
#安装tzdata安装包
RUN apk add --no-cache tzdata
#设置时区
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone
ENV ES_URL http://192.168.2.10:9200/kong*/_search
ENV ES_USERNAME admin
ENV ES_PASSWORD rainbow123
EXPOSE 8088
CMD [ "./main" ]