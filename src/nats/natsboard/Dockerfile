FROM node:20.14.0

LABEL authors="zhouzhenyong"

# 安装natsboard全局
RUN npm install -g natsboard

# 运行gnatsd服务，并将其放到后台
# 注意：这里使用8222端口，因为8222是natsboard的默认端口
#RUN ["gnatsd", "-m", "8222"]

# 启动natsboard，连接到gnatsd服务
# 注意：这里不需要使用CMD，因为natsboard已经在上一步的RUN指令中启动了
# 如果你希望natsboard在gnatsd启动后立即启动，你可以使用类似下面的CMD指令
CMD ["natsboard", "--nats-mon-url", "http://localhost:8222"]
