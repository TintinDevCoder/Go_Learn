# Go 操作 Redis 学习指南

## 概述

Go 语言通过 `go-redis` 库提供对 Redis 的完整支持。`go-redis` 是一个高性能、类型安全的 Redis 客户端，支持 Redis 集群、哨兵、管道、事务等特性。

### 核心特性
- 支持 Redis 7.x 的所有命令
- 自动连接池管理
- 管道和事务支持
- 发布/订阅模式
- 集群和哨兵支持
- 上下文支持（超时、取消）
- 高性能，低内存占用

## 安装和配置

### 1. 安装 go-redis
```bash
go get github.com/redis/go-redis/v9
```

### 2. 导入包
```go
import (
    "context"
    "fmt"
    "time"
    
    "github.com/redis/go-redis/v9"
)
```

## 基本使用

### 创建 Redis 客户端
```go
// 创建上下文
var ctx = context.Background()

// 创建单机 Redis 客户端
func createClient() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", // Redis 地址
        Password: "",               // 密码，没有则留空
        DB:       0,                // 数据库编号
    })
    
    // 测试连接
    pong, err := client.Ping(ctx).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Redis连接成功:", pong)
    
    return client
}
```

### 连接池配置
```go
func createClientWithPool() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:         "localhost:6379",
        Password:     "",
        DB:           0,
        PoolSize:     10,     // 连接池大小
        MinIdleConns: 5,      // 最小空闲连接数
        MaxIdleConns: 10,     // 最大空闲连接数
        DialTimeout:  5 * time.Second,  // 连接超时
        ReadTimeout:  3 * time.Second,  // 读超时
        WriteTimeout: 3 * time.Second,  // 写超时
        PoolTimeout:  4 * time.Second,  // 获取连接超时
    })
    
    return client
}
```

## 数据类型操作

### 字符串（String）操作
```go
func stringOperations(client *redis.Client) {
    // 设置值
    err := client.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }
    
    // 设置带过期时间
    err = client.Set(ctx, "key_with_ttl", "value", 10*time.Second).Err()
    
    // 获取值
    val, err := client.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key:", val)
    
    // 获取不存在的键
    val2, err := client.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2:", val2)
    }
    
    // 自增/自减
    client.Incr(ctx, "counter")
    client.IncrBy(ctx, "counter", 10)
    client.Decr(ctx, "counter")
    
    // 批量操作
    client.MSet(ctx, "key1", "value1", "key2", "value2")
    values, err := client.MGet(ctx, "key1", "key2").Result()
    
    // 追加
    client.Append(ctx, "key", " appended")
}
```

### 哈希（Hash）操作
```go
func hashOperations(client *redis.Client) {
    // 设置哈希字段
    client.HSet(ctx, "user:1", "name", "John")
    client.HSet(ctx, "user:1", "age", "30")
    
    // 批量设置
    client.HSet(ctx, "user:2", map[string]string{
        "name": "Alice",
        "age":  "25",
        "email": "alice@example.com",
    })
    
    // 获取字段
    name, err := client.HGet(ctx, "user:1", "name").Result()
    
    // 获取所有字段
    allFields, err := client.HGetAll(ctx, "user:2").Result()
    for field, value := range allFields {
        fmt.Printf("%s: %s\n", field, value)
    }
    
    // 获取多个字段
    values, err := client.HMGet(ctx, "user:2", "name", "email").Result()
    
    // 删除字段
    client.HDel(ctx, "user:2", "email")
    
    // 字段自增
    client.HIncrBy(ctx, "user:1", "age", 1)
    
    // 检查字段是否存在
    exists, err := client.HExists(ctx, "user:1", "name").Result()
}
```

### 列表（List）操作
```go
func listOperations(client *redis.Client) {
    // 从左侧插入
    client.LPush(ctx, "mylist", "world")
    client.LPush(ctx, "mylist", "hello")
    
    // 从右侧插入
    client.RPush(ctx, "mylist", "!")
    
    // 获取列表长度
    length, err := client.LLen(ctx, "mylist").Result()
    
    // 获取元素
    val, err := client.LIndex(ctx, "mylist", 0).Result()
    
    // 获取范围
    items, err := client.LRange(ctx, "mylist", 0, -1).Result()
    
    // 弹出元素
    popped, err := client.LPop(ctx, "mylist").Result()
    popped, err = client.RPop(ctx, "mylist").Result()
    
    // 修剪列表
    client.LTrim(ctx, "mylist", 0, 2)
}
```

### 集合（Set）操作
```go
func setOperations(client *redis.Client) {
    // 添加元素
    client.SAdd(ctx, "myset", "a", "b", "c")
    
    // 获取所有元素
    members, err := client.SMembers(ctx, "myset").Result()
    
    // 检查元素是否存在
    isMember, err := client.SIsMember(ctx, "myset", "a").Result()
    
    // 获取集合大小
    size, err := client.SCard(ctx, "myset").Result()
    
    // 移除元素
    client.SRem(ctx, "myset", "a")
    
    // 集合运算
    client.SAdd(ctx, "set1", "a", "b", "c")
    client.SAdd(ctx, "set2", "c", "d", "e")
    
    // 交集
    inter, err := client.SInter(ctx, "set1", "set2").Result()
    
    // 并集
    union, err := client.SUnion(ctx, "set1", "set2").Result()
    
    // 差集
    diff, err := client.SDiff(ctx, "set1", "set2").Result()
}
```

### 有序集合（Sorted Set）操作
```go
func sortedSetOperations(client *redis.Client) {
    // 添加元素（带分数）
    client.ZAdd(ctx, "leaderboard", redis.Z{
        Score:  100,
        Member: "player1",
    })
    
    // 批量添加
    client.ZAdd(ctx, "leaderboard", []redis.Z{
        {Score: 200, Member: "player2"},
        {Score: 150, Member: "player3"},
    }...)
    
    // 获取排名（升序）
    rank, err := client.ZRank(ctx, "leaderboard", "player1").Result()
    
    // 获取排名（降序）
    revRank, err := client.ZRevRank(ctx, "leaderboard", "player1").Result()
    
    // 获取分数
    score, err := client.ZScore(ctx, "leaderboard", "player1").Result()
    
    // 获取范围（按分数）
    items, err := client.ZRangeByScore(ctx, "leaderboard", &redis.ZRangeBy{
        Min: "100",
        Max: "200",
    }).Result()
    
    // 获取范围（按排名）
    items, err = client.ZRange(ctx, "leaderboard", 0, 2).Result()
    
    // 增加分数
    client.ZIncrBy(ctx, "leaderboard", 50, "player1")
}
```

## 高级特性

### 管道（Pipeline）
```go
func pipelineExample(client *redis.Client) {
    // 创建管道
    pipe := client.Pipeline()
    
    // 添加多个命令
    pipe.Set(ctx, "key1", "value1", 0)
    pipe.Set(ctx, "key2", "value2", 0)
    pipe.Get(ctx, "key1")
    pipe.Get(ctx, "key2")
    
    // 执行管道
    cmds, err := pipe.Exec(ctx)
    if err != nil {
        panic(err)
    }
    
    // 处理结果
    for _, cmd := range cmds {
        fmt.Println(cmd.String())
    }
}

// 简化的管道使用
func pipelineSimple(client *redis.Client) {
    var get1, get2 *redis.StringCmd
    
    _, err := client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
        get1 = pipe.Get(ctx, "key1")
        get2 = pipe.Get(ctx, "key2")
        return nil
    })
    
    val1, err := get1.Result()
    val2, err := get2.Result()
}
```

### 事务（Transaction）
```go
func transactionExample(client *redis.Client) {
    // Watch 监控键，如果被修改则事务失败
    err := client.Watch(ctx, func(tx *redis.Tx) error {
        // 获取当前值
        n, err := tx.Get(ctx, "counter").Int()
        if err != nil && err != redis.Nil {
            return err
        }
        
        // 开始事务
        _, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
            pipe.Set(ctx, "counter", n+1, 0)
            return nil
        })
        return err
    }, "counter")
    
    if err != nil {
        panic(err)
    }
}
```

### 发布/订阅（Pub/Sub）
```go
func pubSubExample() {
    client := createClient()
    
    // 订阅
    pubsub := client.Subscribe(ctx, "channel1")
    defer pubsub.Close()
    
    // 接收消息
    ch := pubsub.Channel()
    for msg := range ch {
        fmt.Printf("收到消息: %s from %s\n", msg.Payload, msg.Channel)
    }
}

// 发布消息
func publishMessage(client *redis.Client) {
    err := client.Publish(ctx, "channel1", "hello world").Err()
    if err != nil {
        panic(err)
    }
}
```

### 键操作
```go
func keyOperations(client *redis.Client) {
    // 检查键是否存在
    exists, err := client.Exists(ctx, "key1", "key2").Result()
    
    // 删除键
    deleted, err := client.Del(ctx, "key1", "key2").Result()
    
    // 设置过期时间
    client.Expire(ctx, "key", 10*time.Second)
    client.ExpireAt(ctx, "key", time.Now().Add(10*time.Second))
    
    // 获取剩余生存时间
    ttl, err := client.TTL(ctx, "key").Result()
    
    // 移除过期时间
    client.Persist(ctx, "key")
    
    // 重命名键
    client.Rename(ctx, "oldkey", "newkey")
    
    // 扫描键
    var cursor uint64
    var keys []string
    for {
        var scannedKeys []string
        scannedKeys, cursor, err = client.Scan(ctx, cursor, "pattern:*", 10).Result()
        if err != nil {
            panic(err)
        }
        keys = append(keys, scannedKeys...)
        if cursor == 0 {
            break
        }
    }
}
```

## 错误处理

### 常见错误类型
```go
func errorHandling(client *redis.Client) {
    // 检查键不存在错误
    val, err := client.Get(ctx, "nonexistent").Result()
    if err == redis.Nil {
        fmt.Println("键不存在")
    } else if err != nil {
        panic(err)
    }
    
    // 检查网络错误
    _, err = client.Ping(ctx).Result()
    if err != nil {
        if strings.Contains(err.Error(), "connection refused") {
            fmt.Println("无法连接到Redis")
        } else {
            panic(err)
        }
    }
    
    // 检查类型错误
    // 如果对哈希键使用GET命令，会返回类型错误
    client.HSet(ctx, "myhash", "field", "value")
    _, err = client.Get(ctx, "myhash").Result()
    if err != nil && strings.Contains(err.Error(), "WRONGTYPE") {
        fmt.Println("类型错误: 对哈希键使用了GET命令")
    }
}
```

## 性能优化

### 1. 使用连接池
合理配置连接池参数，避免频繁创建连接。

### 2. 使用管道
批量操作使用管道减少网络往返。

### 3. 合理使用数据类型
根据场景选择合适的数据类型。

### 4. 避免大键
单个键的值不宜过大。

### 5. 使用本地缓存
对频繁读取的数据使用本地缓存。

### 6. 监控性能
使用客户端统计信息监控性能。

```go
func monitorPerformance(client *redis.Client) {
    // 获取连接池统计
    stats := client.PoolStats()
    fmt.Printf("总连接数: %d\n", stats.TotalConns)
    fmt.Printf("空闲连接数: %d\n", stats.IdleConns)
    fmt.Printf("命中次数: %d\n", stats.Hits)
    fmt.Printf("等待超时: %d\n", stats.Timeouts)
}
```

## 最佳实践

### 1. 键命名规范
```go
// 使用冒号分隔层级
func keyNaming() {
    // 用户相关
    userKey := "user:1001:profile"
    userSessionKey := "user:1001:session"
    
    // 商品相关
    productKey := "product:2001:info"
    productInventoryKey := "product:2001:inventory"
    
    // 缓存相关
    cacheKey := "cache:homepage:data"
}
```

### 2. 过期时间设置
```go
func expireSettings(client *redis.Client) {
    // 设置合理的过期时间
    client.Set(ctx, "cache:data", "value", 5*time.Minute)
    
    // 对不同的数据类型设置不同的过期策略
    client.Set(ctx, "session:token", "data", 24*time.Hour)      // 会话数据
    client.Set(ctx, "cache:api:response", "data", 1*time.Minute) // API缓存
    client.Set(ctx, "rate:limit:ip", "count", 1*time.Second)     // 限流
}
```

### 3. 序列化策略
```go
import "encoding/json"

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func serializationExample(client *redis.Client) {
    user := User{ID: 1, Name: "John"}
    
    // 序列化为JSON存储
    data, err := json.Marshal(user)
    if err != nil {
        panic(err)
    }
    
    client.Set(ctx, "user:1", data, 0)
    
    // 反序列化
    data, err = client.Get(ctx, "user:1").Bytes()
    if err != nil {
        panic(err)
    }
    
    var fetchedUser User
    err = json.Unmarshal(data, &fetchedUser)
    if err != nil {
        panic(err)
    }
}
```

### 4. 分布式锁实现
```go
func acquireLock(client *redis.Client, lockKey string, timeout time.Duration) (bool, error) {
    // 使用SETNX实现分布式锁
    return client.SetNX(ctx, lockKey, "locked", timeout).Result()
}

func releaseLock(client *redis.Client, lockKey string) error {
    return client.Del(ctx, lockKey).Err()
}
```

## 实际应用示例

### 1. 缓存实现
```go
type Cache struct {
    client *redis.Client
    prefix string
}

func NewCache(client *redis.Client, prefix string) *Cache {
    return &Cache{
        client: client,
        prefix: prefix,
    }
}

func (c *Cache) Get(key string) (string, error) {
    return c.client.Get(ctx, c.prefix+key).Result()
}

func (c *Cache) Set(key, value string, expiration time.Duration) error {
    return c.client.Set(ctx, c.prefix+key, value, expiration).Err()
}

func (c *Cache) Delete(key string) error {
    return c.client.Del(ctx, c.prefix+key).Err()
}

func (c *Cache) Clear() error {
    // 注意：KEYS命令在生产环境慎用，可能阻塞Redis
    keys, err := c.client.Keys(ctx, c.prefix+"*").Result()
    if err != nil {
        return err
    }
    
    if len(keys) > 0 {
        return c.client.Del(ctx, keys...).Err()
    }
    
    return nil
}
```

### 2. 限流器
```go
type RateLimiter struct {
    client *redis.Client
    prefix string
}

func NewRateLimiter(client *redis.Client, prefix string) *RateLimiter {
    return &RateLimiter{
        client: client,
        prefix: prefix,
    }
}

func (r *RateLimiter) Allow(key string, limit int, window time.Duration) (bool, error) {
    key = r.prefix + key
    
    // 使用INCR和EXPIRE实现滑动窗口限流
    count, err := r.client.Incr(ctx, key).Result()
    if err != nil {
        return false, err
    }
    
    if count == 1 {
        // 第一次设置过期时间
        r.client.Expire(ctx, key, window)
    }
    
    return count <= int64(limit), nil
}
```

### 3. 会话存储
```go
type SessionStore struct {
    client *redis.Client
}

func NewSessionStore(client *redis.Client) *SessionStore {
    return &SessionStore{client: client}
}

func (s *SessionStore) Set(sessionID string, data map[string]string, expiration time.Duration) error {
    key := "session:" + sessionID
    return s.client.HSet(ctx, key, data).Err()
}

func (s *SessionStore) Get(sessionID string) (map[string]string, error) {
    key := "session:" + sessionID
    return s.client.HGetAll(ctx, key).Result()
}

func (s *SessionStore) Delete(sessionID string) error {
    key := "session:" + sessionID
    return s.client.Del(ctx, key).Err()
}
```

## 测试

### 单元测试
```go
import "testing"

func TestRedisOperations(t *testing.T) {
    // 使用测试Redis实例
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
    defer client.Close()
    
    // 清空测试数据
    client.FlushDB(ctx)
    
    // 测试字符串操作
    err := client.Set(ctx, "test_key", "test_value", 0).Err()
    if err != nil {
        t.Fatal(err)
    }
    
    val, err := client.Get(ctx, "test_key").Result()
    if err != nil {
        t.Fatal(err)
    }
    
    if val != "test_value" {
        t.Errorf("期望 'test_value', 得到 '%s'", val)
    }
}
```

## 常见问题

### 1. 连接超时
检查Redis服务是否运行，网络是否通畅。

### 2. 内存不足
监控内存使用，设置合理的maxmemory策略。

### 3. 性能下降
使用管道优化批量操作，避免大键。

### 4. 数据不一致
合理使用事务和WATCH命令。

### 5. 连接泄漏
确保及时关闭连接和发布订阅。

## 扩展学习

### 1. Redis集群
```go
func createClusterClient() *redis.ClusterClient {
    client := redis.NewClusterClient(&redis.ClusterOptions{
        Addrs: []string{
            "localhost:7000",
            "localhost:7001",
            "localhost:7002",
        },
        Password: "",
    })
    
    return client
}
```

### 2. Redis哨兵
```go
func createSentinelClient() *redis.Client {
    client := redis.NewFailoverClient(&redis.FailoverOptions{
        MasterName:    "mymaster",
        SentinelAddrs: []string{"localhost:26379", "localhost:26380"},
        Password:      "",
        DB:            0,
    })
    
    return client
}
```

### 3. 性能监控
使用Redis的INFO命令监控服务状态。

## 总结

Go语言的`go-redis`库提供了强大而灵活的Redis操作能力。通过合理使用连接池、管道、事务等特性，可以构建高性能的Redis应用。

建议在实际开发中：
1. 封装Redis操作，提供统一的接口
2. 合理配置连接池参数
3. 使用管道优化批量操作
4. 设置合理的过期时间
5. 监控Redis性能指标
6. 编写单元测试确保功能正确

随着对Redis的深入理解，可以进一步学习Redis集群、哨兵、流等高级特性。