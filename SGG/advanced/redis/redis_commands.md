# Redis 数据结构与操作指令指南

## Redis 简介

Redis（Remote Dictionary Server）是一个开源的内存数据结构存储系统，用作数据库、缓存和消息中间件。它支持多种数据结构，如字符串（String）、哈希（Hash）、列表（List）、集合（Set）、有序集合（Sorted Set）等，并提供丰富的操作命令。

### Redis 特性
- **内存存储**：数据存储在内存中，读写速度极快
- **持久化**：支持 RDB 和 AOF 两种持久化方式
- **数据结构丰富**：支持多种数据结构，适应不同场景
- **原子操作**：所有操作都是原子的，支持事务
- **发布/订阅**：支持消息的发布与订阅模式
- **主从复制**：支持数据的主从复制和高可用

## Redis 数据结构

### 1. 字符串（String）
字符串是 Redis 最基本的数据类型，可以存储文本、数字或二进制数据（最大 512MB）。

**常用场景**：
- 缓存数据
- 计数器
- 分布式锁
- 会话存储

### 2. 哈希（Hash）
哈希是字段和值的映射表，适合存储对象。

**常用场景**：
- 存储用户信息（用户名、邮箱、年龄等）
- 存储商品属性
- 存储配置项

### 3. 列表（List）
列表是简单的字符串列表，按照插入顺序排序，可以在两端插入或删除元素。

**常用场景**：
- 消息队列
- 最新消息列表
- 记录用户操作日志

### 4. 集合（Set）
集合是字符串的无序集合，元素唯一，不支持重复元素。

**常用场景**：
- 标签系统
- 好友关系
- 共同关注

### 5. 有序集合（Sorted Set）
有序集合与集合类似，但每个元素关联一个分数（score），元素按分数排序。

**常用场景**：
- 排行榜
- 带权重的队列
- 时间线

## Redis 连接与基本操作

### 连接 Redis
```bash
# 连接本地 Redis（默认端口 6379）
redis-cli

# 连接远程 Redis
redis-cli -h host -p port -a password

# 连接后验证密码
AUTH password

# 选择数据库（0-15，默认0）
SELECT 1
```

### 基本命令
```bash
# 测试连接是否正常
PING

# 查看服务器信息
INFO

# 查看所有键
KEYS *

# 查看键总数
DBSIZE

# 清空当前数据库
FLUSHDB

# 清空所有数据库
FLUSHALL

# 退出
QUIT
```

## 字符串（String）操作命令

### 设置和获取值
```bash
# 设置键值（如果键已存在则覆盖）
SET key value

# 设置键值并指定过期时间（秒）
SET key value EX seconds
SETEX key seconds value

# 设置键值并指定过期时间（毫秒）
SET key value PX milliseconds
PSETEX key milliseconds value

# 只在键不存在时设置
SETNX key value

# 获取键值
GET key

# 获取多个键值
MGET key1 key2 key3

# 获取旧值并设置新值
GETSET key new_value
```

### 数值操作
```bash
# 将键的值增加1
INCR key

# 将键的值增加指定整数
INCRBY key increment

# 将键的值减少1
DECR key

# 将键的值减少指定整数
DECRBY key decrement

# 将键的值增加指定浮点数
INCRBYFLOAT key increment
```

### 字符串操作
```bash
# 获取字符串长度
STRLEN key

# 追加字符串到已有值末尾
APPEND key value

# 获取子字符串（闭区间）
GETRANGE key start end

# 设置子字符串（从指定偏移量开始）
SETRANGE key offset value
```

## 哈希（Hash）操作命令

### 设置和获取字段
```bash
# 设置单个字段
HSET key field value

# 设置多个字段
HMSET key field1 value1 field2 value2

# 只在字段不存在时设置
HSETNX key field value

# 获取单个字段值
HGET key field

# 获取多个字段值
HMGET key field1 field2

# 获取所有字段和值
HGETALL key

# 获取所有字段名
HKEYS key

# 获取所有字段值
HVALS key
```

### 哈希操作
```bash
# 删除一个或多个字段
HDEL key field1 field2

# 检查字段是否存在
HEXISTS key field

# 获取字段数量
HLEN key

# 获取字段值的字符串长度
HSTRLEN key field

# 字段值增加指定整数
HINCRBY key field increment

# 字段值增加指定浮点数
HINCRBYFLOAT key field increment
```

### 哈希扫描
```bash
# 迭代哈希表中的键值对
HSCAN key cursor [MATCH pattern] [COUNT count]
```

## 列表（List）操作命令

### 插入元素
```bash
# 在列表左侧插入一个或多个元素
LPUSH key element1 element2

# 在列表右侧插入一个或多个元素
RPUSH key element1 element2

# 在列表左侧插入元素（仅当列表存在时）
LPUSHX key element

# 在列表右侧插入元素（仅当列表存在时）
RPUSHX key element

# 在指定元素前/后插入元素
LINSERT key BEFORE|AFTER pivot element
```

### 获取元素
```bash
# 通过索引获取元素（0表示第一个，-1表示最后一个）
LINDEX key index

# 获取列表指定范围内的元素
LRANGE key start stop

# 获取列表长度
LLEN key
```

### 删除元素
```bash
# 移除并返回列表左侧的第一个元素
LPOP key

# 移除并返回列表右侧的第一个元素
RPOP key

# 移除列表中指定数量的匹配元素
LREM key count element

# 修剪列表，只保留指定范围内的元素
LTRIM key start stop
```

### 阻塞操作
```bash
# 阻塞式移除并返回列表左侧的第一个元素
BLPOP key1 key2 timeout

# 阻塞式移除并返回列表右侧的第一个元素
BRPOP key1 key2 timeout

# 从源列表移除最后一个元素并添加到目标列表
RPOPLPUSH source destination

# 阻塞版本的RPOPLPUSH
BRPOPLPUSH source destination timeout
```

## 集合（Set）操作命令

### 添加和删除元素
```bash
# 添加一个或多个元素
SADD key member1 member2

# 移除一个或多个元素
SREM key member1 member2

# 随机移除并返回一个元素
SPOP key [count]

# 随机返回一个或多个元素（不移除）
SRANDMEMBER key [count]
```

### 集合操作
```bash
# 获取所有元素
SMEMBERS key

# 判断元素是否在集合中
SISMEMBER key member

# 获取集合元素数量
SCARD key

# 移动元素到另一个集合
SMOVE source destination member
```

### 集合运算
```bash
# 返回多个集合的交集
SINTER key1 key2 key3

# 计算多个集合的交集并存储到新集合
SINTERSTORE destination key1 key2

# 返回多个集合的并集
SUNION key1 key2 key3

# 计算多个集合的并集并存储到新集合
SUNIONSTORE destination key1 key2

# 返回第一个集合与其他集合的差集
SDIFF key1 key2 key3

# 计算差集并存储到新集合
SDIFFSTORE destination key1 key2
```

### 集合扫描
```bash
# 迭代集合中的元素
SSCAN key cursor [MATCH pattern] [COUNT count]
```

## 有序集合（Sorted Set）操作命令

### 添加和删除元素
```bash
# 添加一个或多个元素（带分数）
ZADD key score1 member1 score2 member2

# 移除一个或多个元素
ZREM key member1 member2

# 移除指定排名范围内的元素
ZREMRANGEBYRANK key start stop

# 移除指定分数范围内的元素
ZREMRANGEBYSCORE key min max
```

### 获取元素
```bash
# 通过索引获取元素（按分数升序）
ZRANGE key start stop [WITHSCORES]

# 通过索引获取元素（按分数降序）
ZREVRANGE key start stop [WITHSCORES]

# 通过分数范围获取元素
ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]

# 获取元素的分数
ZSCORE key member

# 获取元素的排名（升序，0表示第一名）
ZRANK key member

# 获取元素的排名（降序）
ZREVRANK key member
```

### 有序集合操作
```bash
# 获取有序集合元素数量
ZCARD key

# 统计指定分数范围内的元素数量
ZCOUNT key min max

# 增加元素的分数
ZINCRBY key increment member
```

### 有序集合运算
```bash
# 计算多个有序集合的交集并存储到新集合
ZINTERSTORE destination numkeys key1 key2 [WEIGHTS weight1 weight2] [AGGREGATE SUM|MIN|MAX]

# 计算多个有序集合的并集并存储到新集合
ZUNIONSTORE destination numkeys key1 key2 [WEIGHTS weight1 weight2] [AGGREGATE SUM|MIN|MAX]
```

### 有序集合扫描
```bash
# 迭代有序集合中的元素
ZSCAN key cursor [MATCH pattern] [COUNT count]
```

## 键（Key）操作命令

### 键管理
```bash
# 判断键是否存在
EXISTS key

# 删除键
DEL key

# 设置键的过期时间（秒）
EXPIRE key seconds

# 设置键的过期时间（毫秒）
PEXPIRE key milliseconds

# 设置键在指定时间戳过期（秒）
EXPIREAT key timestamp

# 设置键在指定时间戳过期（毫秒）
PEXPIREAT key milliseconds-timestamp

# 移除键的过期时间
PERSIST key

# 获取键的剩余生存时间（秒）
TTL key

# 获取键的剩余生存时间（毫秒）
PTTL key

# 重命名键
RENAME key newkey

# 只在newkey不存在时重命名
RENAMENX key newkey

# 随机返回一个键
RANDOMKEY

# 将键移动到另一个数据库
MOVE key db
```

### 键扫描
```bash
# 迭代数据库中的键
SCAN cursor [MATCH pattern] [COUNT count]
```

## 事务与流水线

### 事务操作
```bash
# 开始事务
MULTI

# 执行事务中的所有命令
EXEC

# 取消事务
DISCARD

# 监视一个或多个键，如果在事务执行前被修改，则事务中断
WATCH key1 key2

# 取消所有WATCH命令
UNWATCH
```

### 流水线（Pipeline）
Redis流水线允许客户端一次性发送多个命令，减少网络往返时间，提高性能。

## 发布/订阅

### 发布订阅命令
```bash
# 向频道发布消息
PUBLISH channel message

# 订阅一个或多个频道
SUBSCRIBE channel1 channel2

# 取消订阅一个或多个频道
UNSUBSCRIBE [channel1 channel2]

# 按模式订阅频道
PSUBSCRIBE pattern1 pattern2

# 取消按模式订阅
PUNSUBSCRIBE [pattern1 pattern2]
```

## 持久化与备份

### 持久化相关命令
```bash
# 保存数据到磁盘（同步，阻塞）
SAVE

# 后台保存数据到磁盘（异步）
BGSAVE

# 查看最后一次SAVE操作的状态
LASTSAVE

# 开启AOF持久化
BGREWRITEAOF
```

## Redis 配置

### 配置相关命令
```bash
# 获取配置参数
CONFIG GET parameter

# 设置配置参数
CONFIG SET parameter value

# 重置统计信息
CONFIG RESETSTAT

# 重写配置文件
CONFIG REWRITE
```

## 实用技巧与最佳实践

### 1. 键命名规范
- 使用冒号分隔层级，如 `user:1001:profile`
- 保持一致性，便于管理和查找

### 2. 合理设置过期时间
- 对缓存数据设置合理的过期时间，避免内存泄漏
- 使用 `EXPIRE` 或 `SETEX` 命令

### 3. 使用批量操作
- 使用 `MGET`、`MSET`、`HMSET` 等批量命令减少网络开销

### 4. 监控内存使用
- 定期使用 `INFO memory` 查看内存使用情况
- 设置 `maxmemory` 策略避免内存溢出

### 5. 连接池管理
- 使用连接池复用连接，减少连接建立开销

### 6. 避免大键
- 单个键的值不宜过大（建议小于1MB）
- 大键会影响网络传输和内存分配

## 常见问题与解决方案

### 1. 内存使用过高
- 使用 `INFO memory` 分析内存使用
- 设置合适的 `maxmemory` 和淘汰策略
- 对大数据进行分片存储

### 2. 性能下降
- 使用 `SLOWLOG` 查看慢查询
- 优化命令使用，避免 `KEYS *` 等阻塞命令
- 使用管道和批量操作

### 3. 数据丢失风险
- 配置合适的持久化策略（RDB + AOF）
- 定期备份数据
- 使用主从复制提高可用性

## 总结

Redis 是一个功能强大的内存数据结构存储系统，掌握其数据结构和操作命令对于开发高性能应用至关重要。本指南涵盖了 Redis 的主要数据结构和常用命令，可作为日常开发的参考手册。

建议在实际使用中结合具体场景选择合适的数据结构和命令，并遵循最佳实践，以充分发挥 Redis 的性能优势。