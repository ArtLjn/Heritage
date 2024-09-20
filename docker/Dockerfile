FROM registry.cn-hangzhou.aliyuncs.com/ubuntu_extend/mysql_rabbitmq_redis:latest
WORKDIR /root/
COPY deploy ./deploy
RUN chmod +x ./deploy/* && \
    chmod +x ./deploy/script/* && \
    chmod +x ./deploy/back/*
EXPOSE 8080
EXPOSE 8081
EXPOSE 3306
EXPOSE 6379
EXPOSE 15672

# Update package sources and install required packages
RUN  service mysql start && \
    mysql -u root -p123456 -e "CREATE DATABASE IF NOT EXISTS heritage CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;" && \
    mysql -u root -p123456 heritage < ./deploy/sql/heritage.sql && \
    sed -i 's/^bind-address\s*=\s*127.0.0.1/bind-address = 0.0.0.0/' /etc/mysql/mysql.conf.d/mysqld.cnf

RUN cp -f ./deploy/nginx.conf /etc/nginx/nginx.conf
WORKDIR /root/deploy/back

CMD ../script/start_all.sh && ./back
