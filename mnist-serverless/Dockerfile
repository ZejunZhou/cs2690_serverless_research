# 使用官方的 Python 3.8 基础镜像 (Use the official Python 3.8 base image)
FROM python:3.8-slim

# 设置工作目录 (Set the working directory)
WORKDIR /function

# 复制当前目录内容到容器内的 /function
# Copy the current directory contents into /function in the container
COPY . /function/

# 升级 pip (Upgrade pip)
RUN pip install --no-cache-dir --upgrade pip

# 安装依赖 (Install dependencies)
RUN pip install --no-cache-dir -r requirements.txt

# 设置入口命令 (Set the entrypoint command)
ENTRYPOINT ["python", "func.py"]
