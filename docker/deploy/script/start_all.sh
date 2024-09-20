#!/bin/bash

function Run() {
  rabbitmq-server -detached
  # 等待 RabbitMQ 启动
  sleep 3
  # 确保 RabbitMQ 应用程序正在运行
  echo "确保 RabbitMQ 应用程序正在运行..."
  rabbitmqctl start_app

  echo "Erlang 和 RabbitMQ 安装完成，RabbitMQ 已启动并启用管理插件。"

  # 添加管理员用户
  echo "添加管理员用户 root..."
  rabbitmqctl add_user root 123456
  rabbitmqctl set_user_tags root administrator

  # 创建虚拟主机
  echo "创建虚拟主机 heritage..."
  rabbitmqctl add_vhost heritage

  # 为 root 用户设置虚拟主机权限
  echo "为 root 用户设置虚拟主机 heritage 的权限..."
  rabbitmqctl set_permissions -p heritage root ".*" ".*" ".*"

  echo "RabbitMQ 安装和配置完成。"
  redis-server /etc/redis/redis.conf
  service mysql restart
  nginx
}

Run
