# 监控指标指南

## 1. 概述

本服务使用 OpenTelemetry 和 Prometheus 实现监控指标收集。主要包括请求计数、延迟统计等关键指标。

## 2. 可用指标

### 2.1 基础指标

- `req_order_count`: 订单请求计数
  - 标签：`order_id`
  - 类型：Counter
  - 描述：记录订单相关请求的总数

### 2.2 性能指标

- 请求延迟直方图
- 请求总数统计

## 3. 指标使用

### 3.1 在代码中添加指标 

