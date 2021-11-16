# Envelope Writer

本项目是[Group8-红包雨](https://github.com/ohroffen/Envelope)的持久化服务，包含以下的功能：
- 以 consumer group 为单位从 kafka 中订阅 `wri` topic 中的红包数据，并插入或更新 Mysql 相应表
