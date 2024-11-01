# 使用官方的 Python 镜像
FROM python:3.9-slim

# 设置工作目录
WORKDIR /app

# 复制应用程序代码
COPY ./app /app

# 安装依赖
RUN pip install --no-cache-dir -r requirements.txt

# 暴露端口
EXPOSE 5000

# 启动应用
CMD ["python", "train.py"]
