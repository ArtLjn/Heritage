FROM registry.cn-hangzhou.aliyuncs.com/ubuntu_extend/mysql_rabbitmq_redis:latest

WORKDIR /root/

COPY deploy ./deploy
RUN chmod +x ./deploy/* && \
    chmod +x ./deploy/script/* && \
    chmod +x ./deploy/back/*

# 暴露服务端口
EXPOSE 8080 8081 3306 6379 15672

# 设置环境变量
ENV CONTRACT_ADDRESS="0x34d59ea8a51d613226137ea5f1f689c875fb4d31"
ENV CONTRACT_USER="0x29e2db8ad37fb85a2466db8517d380dc98c2d51c"
ENV IP="localhost"
ENV MYSQL_PASSWORD="123456"
ENV ADMIN_NAME="admin"
ENV ADMIN_PASSWORD="admin"

# 初始化数据库并导入 SQL
RUN service mysql start && \
    mysql -u root -p${MYSQL_PASSWORD} -e "CREATE DATABASE IF NOT EXISTS heritage CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;" && \
    mysql -u root -p${MYSQL_PASSWORD} heritage < ./deploy/sql/heritage.sql && \
    sed -i 's/^bind-address\s*=\s*127.0.0.1/bind-address = 0.0.0.0/' /etc/mysql/mysql.conf.d/mysqld.cnf

# 拷贝 nginx 配置
RUN cp -f ./deploy/nginx.conf /etc/nginx/nginx.conf

# 设置后端工作目录
WORKDIR /root/deploy/back

CMD ../script/start_all.sh && ./back


