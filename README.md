# 思路

为了实现基于多线程的端口扫描程序，您可以按照以下步骤进行设计：

1. 设计用户界面 您可以使用 Tkinter 或 PyQt 等 GUI 库来创建一个用户友好的界面，包括输入 IP 地址、选择 TCP 或 UDP 等信息的输入框和按钮，以及扫描结果的展示框等。
2. 确定端口范围 根据 TCP 和 UDP 的标准端口范围，确定要扫描的端口号范围。
3. 实现扫描函数 使用 socket 库的 connect_ex()方法来检测端口是否开放，通过多线程技术实现同时对多个端口的扫描。对于每个 IP 地址，创建多个线程同时进行扫描。
4. 展示扫描结果 在 GUI 界面中展示扫描结果，包括每个 IP 地址的开放端口列表。