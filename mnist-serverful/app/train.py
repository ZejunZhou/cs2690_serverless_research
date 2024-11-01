from flask import Flask, request, jsonify
import tensorflow as tf
from tensorflow.keras.datasets import mnist
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import Dense, Flatten
import os

app = Flask(__name__)

@app.route('/train', methods=['GET', 'POST'])
def train_model():
    if request.method == 'POST':
        data = request.get_json()
        epochs = data.get('epochs', 2)
        batch_size = data.get('batch_size', 16)
    else:
        epochs = int(request.args.get('epochs', 2))
        batch_size = int(request.args.get('batch_size', 16))

    # 加载数据
    (x_train, y_train), (x_test, y_test) = mnist.load_data()

    # 数据归一化
    x_train, x_test = x_train / 255.0, x_test / 255.0

    # 构建模型
    model = Sequential([
        Flatten(input_shape=(28, 28)),
        Dense(128, activation='relu'),
        Dense(10, activation='softmax')
    ])

    # 编译模型
    model.compile(optimizer='adam',
                  loss='sparse_categorical_crossentropy',
                  metrics=['accuracy'])

    # 训练模型
    history = model.fit(x_train, y_train, epochs=epochs, batch_size=batch_size, verbose=0)

    # 评估模型
    test_loss, test_acc = model.evaluate(x_test, y_test, verbose=0)

    # 可选：保存模型
    # model.save('model.h5')

    # 返回训练结果
    return jsonify({
        'epochs': epochs,
        'batch_size': batch_size,
        'test_accuracy': test_acc,
        'test_loss': test_loss
    })

if __name__ == '__main__':
    # 获取环境变量中的端口号，默认为 5000
    port = int(os.environ.get('PORT', 5000))
    app.run(host='0.0.0.0', port=port)
