# 长安链DockerGo智能合约编写GPT助手
本示例程序调用了OpenAI的API，实现了长安链DockerGo智能合约的编写助手功能。用户可以通过输入自然语言描述，获取智能合约的代码模板。

## 使用方法
1. 申请Azure的OpenAI API服务，获取API地址、密钥等信息。
2. 设置环境变量`OPENAI_API_KEY`为OpenAI API的密钥，`OPENAI_API_BASE`为API的URL地址。
3. 在Go环境中，运行`go run main.go`启动服务。访问`http://localhost:8080`即可使用。

## 运行效果
![效果截图](./cm-contract-gpt-demo.png)