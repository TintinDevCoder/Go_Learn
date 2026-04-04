package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

// 自定义错误类型
var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
)

// 自定义错误结构体
type ValidationError struct {
	Field   string
	Message string
	Value   interface{}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s (value: %v)",
		e.Field, e.Message, e.Value)
}

func (e *ValidationError) Is(target error) bool {
	_, ok := target.(*ValidationError)
	return ok
}

func main() {
	fmt.Println("=== Go 错误处理示例 ===")

	// 1. 基本错误处理
	basicErrorHandling()

	// 2. 错误包装和检查
	errorWrapping()

	// 3. 自定义错误
	customErrors()

	// 4. 错误类型断言
	errorTypeAssertion()

	// 5. 延迟和清理
	deferredCleanup()

	// 6. 错误处理模式
	errorPatterns()

	fmt.Println("\n=== 错误处理示例完成 ===")
}

// 1. 基本错误处理
func basicErrorHandling() {
	fmt.Println("\n--- 基本错误处理 ---")

	// 函数返回错误
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("除法结果: %v\n", result)
	}

	// 除数为零
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("除法结果: %v\n", result)
	}

	// 字符串转整数
	num, err := strconv.Atoi("123")
	if err != nil {
		fmt.Printf("转换错误: %v\n", err)
	} else {
		fmt.Printf("转换结果: %d\n", num)
	}

	// 无效转换
	num, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Printf("转换错误: %v\n", err)
	} else {
		fmt.Printf("转换结果: %d\n", num)
	}
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// 2. 错误包装和检查
func errorWrapping() {
	fmt.Println("\n--- 错误包装和检查 ---")

	// 模拟多层调用
	err := processData("invalid data")
	if err != nil {
		fmt.Printf("处理数据错误: %v\n", err)

		// 检查特定错误
		if errors.Is(err, ErrInvalidInput) {
			fmt.Println("错误类型: 无效输入")
		}

		// 解包错误链
		fmt.Println("错误链:")
		for unwrapped := err; unwrapped != nil; unwrapped = errors.Unwrap(unwrapped) {
			fmt.Printf("  - %v\n", unwrapped)
		}
	}
}

func processData(data string) error {
	if err := validateData(data); err != nil {
		// 包装错误，添加上下文
		return fmt.Errorf("processData failed: %w", err)
	}
	return nil
}

func validateData(data string) error {
	if data == "" {
		return fmt.Errorf("validateData: %w", ErrInvalidInput)
	}
	if len(data) < 3 {
		return fmt.Errorf("validateData: data too short: %w", ErrInvalidInput)
	}
	return nil
}

// 3. 自定义错误
func customErrors() {
	fmt.Println("\n--- 自定义错误 ---")

	user := map[string]interface{}{
		"name":  "",
		"email": "invalid-email",
		"age":   150,
	}

	// 验证用户数据
	if err := validateUser(user); err != nil {
		fmt.Printf("用户验证失败: %v\n", err)

		// 检查是否为ValidationError
		var valErr *ValidationError
		if errors.As(err, &valErr) {
			fmt.Printf("验证错误详情: 字段=%s, 消息=%s, 值=%v\n",
				valErr.Field, valErr.Message, valErr.Value)
		}
	}
}

func validateUser(user map[string]interface{}) error {
	// 检查姓名
	if name, ok := user["name"].(string); !ok || name == "" {
		return &ValidationError{
			Field:   "name",
			Message: "姓名不能为空",
			Value:   user["name"],
		}
	}

	// 检查邮箱
	if email, ok := user["email"].(string); ok {
		if !isValidEmail(email) {
			return &ValidationError{
				Field:   "email",
				Message: "邮箱格式无效",
				Value:   email,
			}
		}
	}

	// 检查年龄
	if age, ok := user["age"].(int); ok {
		if age < 0 || age > 120 {
			return &ValidationError{
				Field:   "age",
				Message: "年龄必须在0-120之间",
				Value:   age,
			}
		}
	}

	return nil
}

func isValidEmail(email string) bool {
	// 简单的邮箱验证
	return len(email) > 3 && contains(email, "@") && contains(email, ".")
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// 4. 错误类型断言
func errorTypeAssertion() {
	fmt.Println("\n--- 错误类型断言 ---")

	// 模拟文件操作
	err := openAndReadFile("nonexistent.txt")
	if err != nil {
		fmt.Printf("文件操作错误: %v\n", err)

		// 检查是否为路径错误
		if pathErr, ok := err.(*os.PathError); ok {
			fmt.Printf("路径错误: 操作=%s, 路径=%s, 错误=%v\n",
				pathErr.Op, pathErr.Path, pathErr.Err)
		}

		// 使用errors.As
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			fmt.Printf("使用errors.As: 操作=%s, 路径=%s\n", pathErr.Op, pathErr.Path)
		}
	}
}

func openAndReadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 尝试读取
	data := make([]byte, 100)
	_, err = file.Read(data)
	if err != nil && err != io.EOF {
		return fmt.Errorf("读取文件失败: %w", err)
	}

	return nil
}

// 5. 延迟和清理
func deferredCleanup() {
	fmt.Println("\n--- 延迟和清理 ---")

	// 模拟资源管理
	err := withResources()
	if err != nil {
		fmt.Printf("资源操作错误: %v\n", err)
	}
}

func withResources() (err error) {
	// 模拟打开文件
	fmt.Println("打开文件...")

	// 使用延迟函数进行清理
	defer func() {
		fmt.Println("执行清理操作...")
		if r := recover(); r != nil {
			fmt.Printf("从panic恢复: %v\n", r)
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()

	// 模拟可能失败的操作
	simulateOperation()

	// 模拟panic
	// panic("模拟panic")

	return nil
}

func simulateOperation() error {
	// 模拟可能失败的操作
	return nil
}

// 6. 错误处理模式
func errorPatterns() {
	fmt.Println("\n--- 错误处理模式 ---")

	// 模式1: 错误聚合
	errors := collectErrors()
	if len(errors) > 0 {
		fmt.Printf("收集到 %d 个错误:\n", len(errors))
		for i, err := range errors {
			fmt.Printf("  错误 %d: %v\n", i+1, err)
		}
	}

	// 模式2: 重试机制
	maxRetries := 3
	for i := 0; i <= maxRetries; i++ {
		err := unreliableOperation()
		if err == nil {
			fmt.Println("操作成功!")
			break
		}
		fmt.Printf("尝试 %d 失败: %v\n", i+1, err)
		if i == maxRetries {
			fmt.Println("达到最大重试次数，放弃")
		}
	}

	// 模式3: 错误转换
	err := operationWithContext()
	if err != nil {
		fmt.Printf("操作失败: %v\n", err)
		// 转换为用户友好的错误
		userErr := toUserFriendlyError(err)
		fmt.Printf("用户友好错误: %v\n", userErr)
	}
}

func collectErrors() []error {
	var errs []error

	// 模拟多个可能失败的操作
	if err := operation1(); err != nil {
		errs = append(errs, err)
	}
	if err := operation2(); err != nil {
		errs = append(errs, err)
	}
	if err := operation3(); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func operation1() error {
	return nil // 模拟成功
}

func operation2() error {
	return errors.New("operation2 failed")
}

func operation3() error {
	return errors.New("operation3 failed")
}

func unreliableOperation() error {
	// 模拟不可靠的操作，有一定失败概率
	return errors.New("模拟失败")
}

func operationWithContext() error {
	// 模拟带上下文的操作
	return fmt.Errorf("数据库连接失败: %w", errors.New("connection timeout"))
}

func toUserFriendlyError(err error) error {
	// 将技术错误转换为用户友好的错误
	if err == nil {
		return nil
	}

	// 根据错误类型返回不同的用户友好消息
	if errors.Is(err, ErrNotFound) {
		return errors.New("请求的内容不存在")
	}

	if errors.Is(err, ErrUnauthorized) {
		return errors.New("请先登录")
	}

	// 默认错误消息
	return errors.New("操作失败，请稍后重试")
}
