# 使用官方的 Python 3.8 基础镜像
FROM python:3.8-slim

# 设置工作目录
WORKDIR /function

# 复制当前目录内容到容器内的 /function
COPY . /function/

# 升级 pip
RUN pip install --no-cache-dir --upgrade pip

# 安装依赖
RUN pip install --no-cache-dir -r requirements.txt

# 设置入口命令
ENTRYPOINT ["python", "func.py"]
