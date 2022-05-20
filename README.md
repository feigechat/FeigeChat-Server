# FeigeChat-Server

a microservice architecture im(chat) server with golang, suport docker & k8s.

开源飞鸽IM服务端，使用微服务分布式架构，基于go语言，同时支持k8s和docker compose部署。

特点：
- `简化设计`：除了WebSocket推送消息（s2c）使用protobuf协议，其他全部使用HTTP API接口，包括发消息（c2s），方便研究学习和交流。
- `写扩散消息存储`：优先使用写扩散模式（发一个消息，假设群有500成员，存储500份），降低消息同步复杂度，故限制只支持500人以下群聊。
- `微服务分布式架构`：按照理论百万级在线，1W条消息/秒设计和优化架构。
- `k8s动态扩容`：实验特性，会定期在拥有3个节点的k8s集群上部署和测试。

技术栈：
- go 1.18
- kafka
- grpc
- protobuf
- consul
- mongodb：离线消息存储，为方便使用 `timeline` 同步消息给客户端
- mysql：持久化消息，用于消息漫游
- redis

部署方式：
- centos 7
- docker compose
- k8s

## Architecture

模块架构:  
![模块架构](https://raw.githubusercontent.com/xmcy0011/CoffeeChat/master/images/structure-v2.png)

单聊模块交互图:
![单聊](https://raw.githubusercontent.com/xmcy0011/CoffeeChat/master/images/seq-c2c.png)

See More [architecture](https://github.com/xmcy0011/CoffeeChat/blob/master/docs/02-%E6%9E%B6%E6%9E%84%E5%92%8C%E5%8D%8F%E8%AE%AE%E8%AE%BE%E8%AE%A1.md)

## Quick Start

> PS：请切换到**master**分支，编译和运行！

1. 启动Server（要求安装docker desktop >= 4.0.1）:
```bash
$ git clone https://github.com/xmcy0011/CoffeeChat.git
$ cd CoffeeChat/server
# 从代码编译docker镜像，安装mysql,redis等依赖，并自动初始化mysql数据
$ docker-compose.yml up -d
```
2. 编译客户端。推荐iOS客户端（模拟器选择iphone 11），请移步：[client](https://github.com/xmcy0011/CoffeeChat/blob/master/client/cc_flutter_app/README.md)
3. iOS模拟器和app启动后，点击“注册”，更改服务器IP地址为本机IP（不需要输入端口），注册成功后，登录即可。
4. 内置了2个机器人（思知和微信）和3个好友，可以测试聊天功能。

更多细节，请移步：
- [client](https://github.com/xmcy0011/CoffeeChat/blob/master/client/cc_flutter_app/README.md)
- [server](https://github.com/xmcy0011/CoffeeChat/blob/master/server/src/README.md)

### Document

1. [产品介绍](https://github.com/xmcy0011/CoffeeChat/blob/master/docs/01-%E4%BA%A7%E5%93%81%E4%BB%8B%E7%BB%8D.md)
2. [架构和协议设计](https://github.com/xmcy0011/CoffeeChat/blob/master/docs/02-%E6%9E%B6%E6%9E%84%E5%92%8C%E5%8D%8F%E8%AE%AE%E8%AE%BE%E8%AE%A1.md)
3. [消息分表存储](https://github.com/xmcy0011/CoffeeChat/blob/master/docs/03-%E6%B6%88%E6%81%AF%E5%88%86%E8%A1%A8%E5%AD%98%E5%82%A8.md)
4. [IM 消息 ID 生成原理和常见技术难点](https://github.com/xmcy0011/CoffeeChat/blob/master/docs/04_IM%e5%b8%b8%e8%a7%81%e6%8a%80%e6%9c%af%e9%9a%be%e7%82%b9.md)
5. [进度计划](https://github.com/xmcy0011/CoffeeChat/blob/master/docs/05-%E8%BF%9B%E5%BA%A6%E8%AE%A1%E5%88%92.md)
6. [MQ在IM中的实践和选型](https://github.com/xmcy0011/CoffeeChat/blob/master/docs/06_MQ%e5%9c%a8IM%e4%b8%ad%e7%9a%84%e5%ae%9e%e8%b7%b5.md)

更多文章请移步：
- [CoffeeChat-GitBook](https://xmcy0011.github.io/CoffeeChat-GitBook/)

### Thinks

- [mattermost](https://github.com/mattermost/mattermost-server)：主要学习其go工程实践方面的一些技巧，目前还处在研究阶段。
- [Open-IM-Server](https://github.com/OpenIMSDK/Open-IM-Server)：通过分析它的架构和代码，理解了收件箱机制和im 微服务(go)的划分实践。
- [goim](https://github.com/Terry-Mao/goim)：学习到百万级架构下kafka是如何应用在聊天室场景的。
- [Terry-Ye/im](https://github.com/Terry-Ye/im)：结合goim，理解了所谓的job含义，看懂了goim的架构。
- [gim](https://github.com/alberliu/gim)：一个简单的写扩散项目，可以更深入理解写扩散的架构和原理。

更多开源im，请移步：[史上最全开源IM盘点](https://blog.csdn.net/xmcy001122/article/details/110679978)

## Contact

email：xmcy0011@sina.com  
微信交流：xuyc1992（请备注：im）  

喜欢的话，关注下公众号吧😊  
《Go和分布式IM》👈👈  
![qrcode](./images/qrcode.png)

## LICENSE

FeigeChat is provided under the [mit license](https://github.com/xmcy0011/CoffeeChat/blob/master/LICENSE).