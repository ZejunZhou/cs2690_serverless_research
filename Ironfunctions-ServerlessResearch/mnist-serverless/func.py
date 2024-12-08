# func.py
import sys
import json
import os
import tensorflow as tf
from tensorflow.keras.datasets import mnist
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import Dense, Flatten

def main():
    # 从 stdin 读取请求体 (Read the request body from stdin)
    req = sys.stdin.read()
    
    # 解析请求数据 (Parse the request data)
    try:
        data = json.loads(req)
        epochs = data.get('epochs', 2)
        batch_size = data.get('batch_size', 16)
    except json.JSONDecodeError:
        # 如果解析失败，使用默认参数 (using default parameters here)
        epochs = 2
        batch_size = 16

    # 加载数据 (loading data)
    (x_train, y_train), (x_test, y_test) = mnist.load_data()

    # 数据归一化 (training data)
    x_train, x_test = x_train / 255.0, x_test / 255.0

    # 构建模型 (building model)
    model = Sequential([
        Flatten(input_shape=(28, 28)),
        Dense(128, activation='relu'),
        Dense(10, activation='softmax')
    ])

    # 编译模型 (compile model)
    model.compile(optimizer='adam',
                  loss='sparse_categorical_crossentropy',
                  metrics=['accuracy'])

    # 训练模型 (train model)
    model.fit(x_train, y_train, epochs=epochs, batch_size=batch_size, verbose=0)

    # 评估模型 (evaluate model)
    test_loss, test_acc = model.evaluate(x_test, y_test, verbose=0)

    # 可选：保存模型（如果需要保存模型，请确保有合适的存储路径）
    # Optional: save model (ensure a suitable storage path if saving the model is needed)
    # model.save('model.h5')

    # 返回训练结果 (return training results)
    response = {
        'epochs': epochs,
        'batch_size': batch_size,
        'test_accuracy': test_acc,
        'test_loss': test_loss
    }

    print(json.dumps(response))

if __name__ == '__main__':
    main()
