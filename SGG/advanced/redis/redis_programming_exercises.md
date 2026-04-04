# Redis 编程练习题

基于《Go 操作 Redis 学习指南》，以下是针对每个部分的编程练习题。

## 一、基本使用

### 1. 连接管理
1. **编写一个函数 `TestConnection()`**，创建 Redis 客户端并测试连接，如果连接成功打印 "连接成功"，否则打印错误信息。

2. **编写函数 `CreateCustomClient(addr, password string, db int, poolSize int) *redis.Client`**，根据参数创建自定义配置的 Redis 客户端。

## 二、数据类型操作

### 2.1 字符串操作
3. **实现一个计数器功能**，包含以下方法：
   - `Increment(key string, value int64)`: 按指定值增加计数器
   - `Decrement(key string, value int64)`: 按指定值减少计数器
   - `GetValue(key string)`: 获取计数器当前值
   - `Reset(key string)`: 重置计数器为0

4. **编写函数 `BatchSetAndGet(keys []string, values []string)`**，批量设置键值对并验证设置是否正确。

### 2.2 哈希操作
5. **实现用户信息管理系统**，使用哈希存储用户信息（ID、姓名、年龄、邮箱），提供以下功能：
   - `AddUser(id, name, age, email string)`: 添加用户
   - `GetUser(id string)`: 获取用户信息
   - `UpdateUser(id, field, value string)`: 更新用户字段
   - `DeleteUser(id string)`: 删除用户

6. **编写函数 `UpdateUserAge(userID string, increment int)`**，安全地更新用户的年龄（使用管道确保原子性）。

### 2.3 列表操作
7. **实现简单的消息队列**，支持以下操作：
   - `Enqueue(queueName, message string)`: 从右侧入队
   - `Dequeue(queueName string)`: 从左侧出队
   - `QueueLength(queueName string)`: 获取队列长度
   - `Peek(queueName string, n int)`: 查看前n个消息（不出队）

8. **编写函数 `GetLastNElements(listKey string, n int)`**，获取列表的最后n个元素。

### 2.4 集合操作
9. **实现标签系统**，支持以下功能：
   - `AddTag(articleID, tag string)`: 给文章添加标签
   - `RemoveTag(articleID, tag string)`: 移除文章标签
   - `GetArticleTags(articleID string)`: 获取文章的所有标签
   - `GetArticlesByTag(tag string)`: 获取具有特定标签的文章

10. **编写函数 `FindCommonTags(articleIDs []string)`**，找出多篇文章的共同标签。

### 2.5 有序集合操作
11. **实现游戏排行榜系统**，支持以下功能：
    - `AddScore(player string, score float64)`: 添加/更新玩家分数
    - `GetRank(player string)`: 获取玩家排名（从高到低）
    - `GetTopN(n int)`: 获取前N名玩家
    - `GetPlayersByScoreRange(min, max float64)`: 获取分数在指定范围内的玩家

12. **编写函数 `UpdateScore(player string, delta float64)`**，增加玩家的分数（使用ZINCRBY）。

## 三、高级特性

### 3.1 管道
13. **使用管道优化批量数据插入**，编写函数 `BatchInsert(records map[string]string)`，插入100个键值对。

14. **实现函数 `BatchIncrement(counterKeys []string)`**，使用管道批量增加多个计数器的值。

### 3.2 事务
15. **实现安全的计数器操作**，使用 `WATCH` 确保在事务执行期间计数器没有被其他客户端修改：
    ```go
    func SafeIncrement(key string, increment int) error
    ```

16. **编写转账函数 `Transfer(from, to string, amount int)`**，使用事务确保转账的原子性。

### 3.3 发布/订阅
17. **实现简单聊天室系统**：
    - `PublishMessage(channel, message string)`: 发布消息
    - `SubscribeChannel(channel string)`: 订阅频道并打印收到的消息
    - 实现多个goroutine同时订阅同一频道

18. **编写事件通知系统**，当用户完成操作时发布事件：
    - `PublishEvent(eventType, data string)`: 发布事件
    - `SubscribeEvents(eventTypes []string)`: 订阅多个事件类型

### 3.4 键操作
19. **编写函数 `FindKeysByPattern(pattern string)`**，使用 `SCAN` 命令查找匹配模式的键。

20. **实现缓存清理函数**，删除所有过期的缓存键。

## 四、错误处理

21. **编写健壮的 `GetValue` 函数**，正确处理以下情况：
    - 键不存在
    - 网络错误
    - 类型错误（如对哈希键使用GET命令）
    - 返回适当的错误信息

22. **实现重试机制**，当 Redis 操作失败时自动重试指定次数：
    ```go
    func RetryOperation(operation func() error, maxRetries int) error
    ```

## 五、性能优化

23. **实现性能监控函数**，定期收集并输出 Redis 连接池统计信息：
    - 总连接数
    - 空闲连接数
    - 命中次数
    - 等待超时次数

24. **实现缓存包装器**，添加缓存命中率统计功能。

## 六、最佳实践

### 6.1 键命名规范
25. **设计电商系统键命名方案**，实现以下键的生成函数：
    - 用户信息：`user:{id}:profile`
    - 商品信息：`product:{id}:info`
    - 购物车：`cart:{user_id}`
    - 订单：`order:{order_id}`

### 6.2 过期时间设置
26. **为不同类型的缓存数据设置合理的过期时间**：
    - 用户会话数据：24小时
    - API响应缓存：5分钟
    - 热门商品信息：1小时
    - 限流计数器：1秒

### 6.3 序列化策略
27. **实现通用缓存包装器**，支持任意类型的序列化和反序列化（使用JSON）：
    ```go
    type Cache struct {
        client *redis.Client
        prefix string
    }
    
    func (c *Cache) Set(key string, value interface{}, expiration time.Duration) error
    func (c *Cache) Get(key string, result interface{}) error
    ```

### 6.4 分布式锁
28. **实现完整的分布式锁**：
    ```go
    type DistributedLock struct {
        client *redis.Client
        key    string
        value  string
    }
    
    func (l *DistributedLock) Acquire(timeout time.Duration) (bool, error)
    func (l *DistributedLock) Release() error
    func (l *DistributedLock) Renew(extension time.Duration) error
    ```

29. **编写使用分布式锁的示例**，确保资源独占访问。

## 七、实际应用

### 7.1 缓存实现
30. **扩展基础缓存结构**，添加以下功能：
    - 缓存命中率统计
    - 缓存预热
    - 缓存穿透防护（缓存空值）

31. **实现多级缓存系统**（Redis + 本地内存缓存）。

### 7.2 限流器
32. **实现滑动窗口限流器**：
    ```go
    type RateLimiter struct {
        client *redis.Client
        prefix string
    }
    
    func (r *RateLimiter) Allow(key string, limit int, window time.Duration) (bool, error)
    ```

33. **实现令牌桶算法限流器**。

### 7.3 会话存储
34. **扩展会话存储**，添加会话过期时间管理：
    ```go
    type SessionStore struct {
        client *redis.Client
    }
    
    func (s *SessionStore) Set(sessionID string, data map[string]string, expiration time.Duration) error
    func (s *SessionStore) Get(sessionID string) (map[string]string, error)
    func (s *SessionStore) IsValid(sessionID string) (bool, error)
    ```

35. **实现分布式会话管理系统**，支持会话共享。

## 八、测试

36. **为字符串操作编写单元测试**，覆盖以下情况：
    - 正常设置和获取
    - 键不存在
    - 设置过期时间
    - 批量操作

37. **编写性能测试**，比较使用管道和不使用管道的性能差异。

38. **编写集成测试**，测试完整的缓存系统功能。

## 九、综合项目

### 项目一：电商系统缓存设计
39. **设计并实现电商系统的完整缓存方案**，包括：
    - 商品信息缓存（带过期时间和预热）
    - 用户会话管理（分布式会话）
    - 购物车实现（使用哈希）
    - 订单缓存（使用有序集合按时间排序）
    - 库存缓存（使用事务确保一致性）
    - 限流和防刷机制

### 项目二：社交网络功能
40. **使用 Redis 实现社交网络功能**：
    - 用户关注/粉丝系统（使用集合）
    - 动态消息流（Timeline，使用列表和有序集合）
    - 点赞和评论计数（使用哈希和管道）
    - 在线用户状态（使用有序集合记录最后活跃时间）
    - 私信系统（使用列表和发布/订阅）

### 项目三：实时数据分析
41. **构建实时数据分析系统**：
    - 实时计数器（PV/UV，使用HyperLogLog）
    - 热门内容排行榜（使用有序集合）
    - 用户行为分析（使用位图记录用户行为）
    - 实时推荐系统（基于用户行为相似度）

## 十、挑战题

42. **实现 Redis 事务回滚机制**（虽然 Redis 不支持真正的回滚，但可以通过补偿操作实现）。

43. **实现分布式ID生成器**，使用 Redis 生成全局唯一的递增ID。

44. **实现延时队列**，使用有序集合实现任务延时执行。

45. **实现分布式信号量**，控制对共享资源的并发访问数量。

46. **实现布隆过滤器**，使用 Redis 的位图实现高效的元素存在性检查。

## 学习建议

1. 从简单题目开始，逐步增加难度
2. 每个题目都实际编写代码并运行测试
3. 思考不同实现方案的优缺点
4. 参考 `redis_guide.go` 中的示例代码
5. 在实际环境中测试性能

每个题目都应包含完整的函数签名和简要说明，具体的实现需要你自己完成。祝你学习顺利！