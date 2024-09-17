FROM ubuntu:20.04
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
RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list && \
    sed -i 's@//ports.ubuntu.com@//mirrors.ustc.edu.cn@g' /etc/apt/sources.list && \
    apt-get update && \
    apt-get install -y curl unzip wget && \
    apt-get install -y redis-server && \
    sed -i 's/^bind 127.0.0.1 ::1/bind 0.0.0.0/' /etc/redis/redis.conf && \
    sed -i 's/^protected-mode yes/protected-mode no/' /etc/redis/redis.conf && \
    redis-server /etc/redis/redis.conf && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y mysql-server mysql-client && \
    service mysql start && \
    sleep 5 && \
    mysql -u root -e "INSTALL PLUGIN validate_password SONAME 'validate_password.so'; \
    SET GLOBAL validate_password_policy=0; \
    SET GLOBAL validate_password_mixed_case_count=0; \
    SET GLOBAL validate_password_number_count=0; \
    SET GLOBAL validate_password_special_char_count=0; \
    SET GLOBAL validate_password_length=6; \
    SHOW VARIABLES LIKE 'validate_password%'; \
    ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '123456'; \
    UPDATE mysql.user SET Host='%' WHERE User='root' AND Host='localhost'; \
    FLUSH PRIVILEGES;" && \
    mysql -u root -p123456 -e "CREATE DATABASE IF NOT EXISTS heritage CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;" && \
    mysql -u root -p123456 heritage < ./deploy/sql/heritage.sql && \
    sed -i 's/^bind-address\s*=\s*127.0.0.1/bind-address = 0.0.0.0/' /etc/mysql/mysql.conf.d/mysqld.cnf && \
    service mysql restart && \
    apt install -y nginx && \
    ./deploy/script/install_mq.sh


RUN cp -f ./deploy/nginx.conf /etc/nginx/nginx.conf
WORKDIR /root/deploy/back

CMD ../script/start_all.sh && ./back