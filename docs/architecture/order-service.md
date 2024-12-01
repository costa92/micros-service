# Order Service 架构文档

## 1. 服务概述

Order Service 是一个基于 Kratos 框架开发的微服务，主要负责订单相关的业务处理。该服务采用了清晰的分层架构，遵循 DDD (Domain-Driven Design) 设计理念。

## 2. 技术栈

- 框架：Kratos v2
- 数据库：MySQL
- 服务发现：支持 Etcd/Consul
- 监控：Prometheus
- 链路追踪：OpenTelemetry
- 依赖注入：Wire

## 3. 项目结构 