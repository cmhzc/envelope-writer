FROM centos:7
ENV KAFKA_HOST kafka-ymqfnev51bodns.cn-beijing.kafka-internal.ivolces.com:9092
ENV KAFKA_GROUPID writer
ENV KAFKA_TOPIC wri
ENV MYSQL_USERNAME group8
ENV MYSQL_PASSWORD Group12345678
ENV MYSQL_HOST rdsmysqlhf6ed4fde2675f947.rds.ivolces.com
ENV MYSQL_PORT 3306
ENV MYSQL_DBNAME envelope_rains
WORKDIR /root
COPY envelope_db_writer ./server
EXPOSE 9090
CMD /root/server
