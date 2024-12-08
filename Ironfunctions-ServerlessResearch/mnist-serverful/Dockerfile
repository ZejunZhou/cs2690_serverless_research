# 使用官方的 Python 镜像 (Use the official Python image)
FROM python:3.9-slim

# 设置工作目录 (Set the working directory)
WORKDIR /app

# 复制应用程序代码 (Copy application code)
COPY ./app /app

# 安装依赖 (Install dependencies)
RUN pip install --no-cache-dir -r requirements.txt

# 暴露端口 (Expose port num)
EXPOSE 5000

# 启动应用 (start app)
CMD ["python", "train.py"]
