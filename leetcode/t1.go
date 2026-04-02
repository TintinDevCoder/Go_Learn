package main

import (
	tree "leetcode/TreeNode"
	"math"
	"math/rand"
	"slices"
	"strconv"
	"strings"
)

/*
*	221. 最大正方形 https://leetcode.cn/problems/maximal-square/
 */
func maximalSquare(matrix [][]byte) int {
	len1 := len(matrix)
	len2 := len(matrix[0])
	dp := make([][]int, len1+1)
	for i := range len1 {
		dp[i] = make([]int, len2+1)
	}
	result := 0
	// 初始化
	for i := 0; i < len1; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			result = 1
		}
	}
	for i := 0; i < len2; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			result = 1
		}
	}
	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
			if dp[i][j] > result {
				result = dp[i][j]
			}
		}
	}
	return result * result
}

/*
* 739. 每日温度 https://leetcode.cn/problems/daily-temperatures/
 */
// 从右往左
func dailyTemperatures1(temperatures []int) []int {
	length := len(temperatures)
	result := make([]int, length)
	// 预分配容量以提升性能
	stack := make([]int, 0, length)
	// 从右向左遍历
	for i := length - 1; i >= 0; i-- {
		currentTemp := temperatures[i]
		// 维护单调递减栈（栈底到栈顶）
		// 弹出所有比当前温度小或相等的元素
		for len(stack) > 0 && currentTemp >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		// 如果栈不为空，说明找到了右侧第一个更高的温度
		if len(stack) > 0 {
			result[i] = stack[len(stack)-1] - i
		} else {
			// 栈为空说明右侧没有更高的温度
			result[i] = 0
		}
		// 将当前索引压入栈中
		stack = append(stack, i)
	}
	return result
}

// 从左往右
func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := []int{}
	for i, v := range temperatures {
		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

/*
215. 数组中的第K个最大元素 https://leetcode.cn/problems/kth-largest-element-in-an-array/
*/
// 快排
func findKthLargest1(nums []int, k int) int {
	res := QuickSort2(nums, 0, len(nums)-1, len(nums)-k)
	return res
}

func QuickSort1(nums []int) int {
	for len(nums) > 0 {
		randIndex := rand.Intn(len(nums))
		nums[randIndex], nums[len(nums)-1] = nums[len(nums)-1], nums[randIndex]
		pre := nums[randIndex]
		right := len(nums) - 1
		left := 0
		for i := 0; i <= right; {
			if nums[i] < pre {
				nums[left], nums[i] = nums[i], nums[left]
				left++
				i++
			} else if nums[i] > pre {
				nums[right], nums[i] = nums[i], nums[right]
				right--
			} else {
				i++
			}
		}
		a, b := nums[:left], nums[right+1:]
		if len(a) > len(b) {
			QuickSort1(b)
			nums = a
		} else {
			QuickSort1(a)
			nums = b
		}
	}
	return 1
}
func QuickSort2(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	pivot := nums[l]
	i := l
	j := r
	for i < j {
		for i < j && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		if i >= j {
			break
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	if i == k {
		return nums[i]
	}
	if k < i {
		return QuickSort2(nums, l, i-1, k)
	} else {
		return QuickSort2(nums, i+1, r, k)
	}
}

// 堆排序
func findKthLargest2(nums []int, k int) int {
	buildMaxHeap(nums)
	heapSize := len(nums)
	for i := 0; i < k-1; i++ {
		nums[heapSize-i-1], nums[0] = nums[0], nums[heapSize-i-1]
		maxHeapify(nums, 0, heapSize-i-1)
	}
	return nums[0]
}
func buildMaxHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		maxHeapify(nums, i, len(nums))
	}
}
func maxHeapify(nums []int, i, heapSize int) {
	l, r, n := i*2+1, i*2+2, i
	if l < heapSize && nums[l] > nums[n] {
		n = l
	}
	if r < heapSize && nums[r] > nums[n] {
		n = r
	}
	if n != i {
		nums[i], nums[n] = nums[n], nums[i]
		maxHeapify(nums, n, heapSize)
	}
}

/*
200. 岛屿数量 https://leetcode.cn/problems/number-of-islands/
*/
var numIslandstarget = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func numIslands(grid [][]byte) int {
	len1, len2 := len(grid), len(grid[0])
	result := 0
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if grid[i][j] == '1' {
				result++
				numIslandsDfs(grid, i, j)
			}
		}
	}
	return result
}
func numIslandsDfs(grid [][]byte, x, y int) {
	grid[x][y] = '0'
	for i := 0; i < 4; i++ {
		newx := x + numIslandstarget[i][0]
		newy := y + numIslandstarget[i][1]
		if newx >= 0 && newx < len(grid) && newy >= 0 && newy < len(grid[0]) && grid[newx][newy] == '1' {
			numIslandsDfs(grid, newx, newy)
		}
	}
}

/*
198. 打家劫舍 https://leetcode.cn/problems/house-robber/
*/
func rob(nums []int) int {
	len := len(nums)
	dp := make([][2]int, len) //0: 偷  1: 不偷
	dp[0][0] = nums[0]
	for i := 1; i < len; i++ {
		dp[i][0] = dp[i-1][1] + nums[i]
		dp[i][1] = max(dp[i-1][0], dp[i-1][1])
	}
	return max(dp[len-1][0], dp[len-1][1])
}

/*
169. 多数元素 https://leetcode.cn/problems/majority-element/
*/
func majorityElement(nums []int) int {
	result := nums[0]
	num := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != result {
			num--
		} else {
			num++
		}
		if num == 0 {
			result = nums[i]
			num = 1
		}
	}
	return result
}

/*
238. 除了自身以外数组的乘积 https://leetcode.cn/problems/product-of-array-except-self/
*/
func productExceptSelf(nums []int) []int {
	len := len(nums)
	result := make([]int, len)
	for i := 0; i < len; i++ {
		result[i] = 1
	}
	left, right := 0, len-1
	lp, rp := 1, 1
	for left < len && right >= 0 {
		result[left] *= lp
		result[right] *= rp
		lp *= nums[left]
		left++
		rp *= nums[right]
		right--
	}
	return result
}

/*
152. 乘积最大子数组 https://leetcode.cn/problems/maximum-product-subarray/
*/
func maxProduct(nums []int) int {
	if len(nums) == 1 && nums[0] < 0 {
		return nums[0]
	}
	now := 1
	leftMax, rightMax := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			now = 1
		} else {
			now *= nums[i]
			leftMax = max(now, leftMax)
		}
	}
	now = 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			now = 1
		} else {
			now *= nums[i]
			rightMax = max(now, rightMax)
		}
	}
	return max(leftMax, rightMax)
}

/*
139. 单词拆分 https://leetcode.cn/problems/word-break/
*/
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
160. 相交链表
*/
// 双指针
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	len1 := 0
	len2 := 0
	HA := headA
	HB := headB
	for HA != nil {
		len1++
		HA = HA.Next
	}
	for HB != nil {
		len2++
		HB = HB.Next
	}
	len3 := len1 - len2
	if len1 >= len2 {
		HA = headA
		HB = headB
	} else {
		len3 = len2 - len1
		HA = headB
		HB = headA
	}
	for len3 > 0 {
		HA = HA.Next
		len3--
	}
	for HA != nil && HB != nil {
		if HA == HB {
			return HA
		}
		HA = HA.Next
		HB = HB.Next
	}
	return nil
}

// 哈希表法
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for h := headA; h != nil; h = h.Next {
		vis[h] = true
	}
	for h := headB; h != nil; h = h.Next {
		if vis[h] {
			return h
		}
	}
	return nil
}

/*
236. 二叉树的最近公共祖先
*/
// 递归
func lowestCommonAncestor1(root, p, q *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)
	if root == p || root == q {
		return root
	}
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

// 哈希
func lowestCommonAncestor2(root, p, q *tree.TreeNode) *tree.TreeNode {
	parent := map[int]*tree.TreeNode{}
	vis := map[int]bool{}
	dfs := func(*tree.TreeNode) {}
	dfs = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parent[node.Left.Val] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			parent[node.Right.Val] = node
			dfs(node.Right)
		}
	}
	dfs(root)
	for p != nil {
		vis[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if vis[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

/*
234. 回文链表
*/
// 反转链表
func isPalindrome(head *ListNode) bool {
	H := &ListNode{0, head}
	slow := H
	fast := H
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	left := head
	right := reverseListNode(slow)
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
func reverseListNode(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	slow.Next = nil
	for fast != nil {
		temp := fast.Next
		fast.Next = slow
		slow = fast
		fast = temp
	}
	return slow
}

/*
226. 翻转二叉树
*/
func invertTree(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	left := root.Left
	root.Left = root.Right
	root.Right = left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

/*
207. 课程表
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	target := make(map[int][]int)
	visit := make(map[int]int)
	for _, v := range prerequisites {
		target[v[0]] = append(target[v[0]], v[1])
	}
	for _, v := range prerequisites {
		if !canFinishDfs(v[0], target, visit) {
			return false
		}
	}
	return true
}
func canFinishDfs(index int, target map[int][]int, visit map[int]int) bool {
	if visit[index] == 2 {
		return false
	} else if visit[index] == 1 {
		return true
	}
	visit[index] = 2
	for _, v := range target[index] {
		if !canFinishDfs(v, target, visit) {
			return false
		}
	}
	visit[index] = 1
	return true
}

/*
206. 反转链表
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head.Next
	slow.Next = nil
	for fast != nil {
		temp := fast.Next
		fast.Next = slow
		slow = fast
		fast = temp
	}
	return slow
}

/*
148. 排序链表
*/
func sortList(head *ListNode) *ListNode {
	return sortListT(head, nil)
}
func sortListT(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return nil
	}
	if left.Next == right {
		left.Next = nil
		return left
	}
	slow := left
	fast := left
	for fast != right {
		slow = slow.Next
		fast = fast.Next
		if fast != right {
			fast = fast.Next
		}
	}
	return mergeList(sortListT(left, slow), sortListT(slow, right))
}
func mergeList(list1 *ListNode, list2 *ListNode) *ListNode {
	H := &ListNode{Val: 0, Next: nil}
	list3 := H
	for list1 != nil || list2 != nil {
		if list2 == nil || (list1 != nil && list1.Val < list2.Val) {
			list3.Next = list1
			list1 = list1.Next
		} else {
			list3.Next = list2
			list2 = list2.Next
		}
		list3 = list3.Next
	}
	return H.Next
}

// 142. 环形链表 II
func detectCycle(head *ListNode) *ListNode {
	var slow, fast *ListNode = head, head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		if slow == fast {
			temp := head
			for fast != temp {
				fast = fast.Next
				temp = temp.Next
			}
			return fast
		}
	}
	return nil
}

// 141. 环形链表
// 双指针法
func hasCycle1(head *ListNode) bool {
	var slow, fast *ListNode = head, head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 哈希表法
func hasCycle2(head *ListNode) bool {
	hash := make(map[*ListNode]struct{})
	for ; head != nil; head = head.Next {
		if _, ok := hash[head]; ok {
			return true
		}
		hash[head] = struct{}{}
	}
	return false
}

// 136. 只出现一次的数字
func singleNumber(nums []int) int {
	num := 0
	for _, n := range nums {
		num ^= n
	}
	return num
}

// 647. 回文子串
// dp
func countSubstrings1(s string) int {
	length := len(s)
	dp := make([][]bool, length)
	for i := range length {
		dp[i] = make([]bool, length)
	}
	result := 0
	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			if s[i] == s[j] {
				if j-i <= 1 || dp[i+1][j-1] {
					result++
					dp[i][j] = true
				}
			}
		}
	}
	return result
}

// 遍历
func countSubstrings2(s string) int {
	length := len(s)
	result := 0
	for i := 0; i < length; i++ {
		var left, right int = i, i
		for left >= 0 && right < length && s[left] == s[right] {
			result++
			left--
			right++
		}
		left, right = i, i+1
		for left >= 0 && right < length && s[left] == s[right] {
			result++
			left--
			right++
		}
	}
	return result
}

// 128. 最长连续序列
// 哈希表法
func longestConsecutive1(nums []int) int {
	// 1. 使用空结构体节省内存，构建哈希集合
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	longest := 0
	for num := range set {
		// 2. 核心优化：只有当 num-1 不存在时，才说明 num 是一个序列的起点
		// 这保证了每个数字最多被访问两次，时间复杂度为 O(n)
		if _, hasPre := set[num-1]; !hasPre {
			currentNum := num
			currentStreak := 1

			// 3. 向后枚举连续序列
			for {
				if _, hasNext := set[currentNum+1]; hasNext {
					currentNum++
					currentStreak++
				} else {
					break
				}
			}

			if currentStreak > longest {
				longest = currentStreak
			}
		}
	}
	return longest
}

// 带备忘录的递归（记忆化搜索）
func longestConsecutive2(nums []int) int {
	hash := make(map[int]int)
	hashV := make(map[int]int)
	for _, num := range nums {
		hash[num] = 1
	}
	result := 0
	for _, num := range nums {
		n := longestConsecutiveDg(num, hash, hashV)
		result = max(result, n)
	}
	return result
}
func longestConsecutiveDg(num int, hash map[int]int, hashV map[int]int) int {
	if hash[num] == 2 {
		return hashV[num]
	} else if hash[num] == 1 {
		n := longestConsecutiveDg(num-1, hash, hashV) + 1
		hashV[num] = n
		hash[num] = 2
		return n
	} else {
		return 0
	}
}

// 124. 二叉树中的最大路径和
func maxPathSum(root *tree.TreeNode) int {
	result := math.MinInt
	var getGain func(*tree.TreeNode) int
	getGain = func(node *tree.TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(getGain(node.Left), 0)
		right := max(getGain(node.Right), 0)
		n := node.Val + left + right
		result = max(result, n)
		return max(left, right) + node.Val
	}
	getGain(root)
	return result
}

// 322. 零钱兑换
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt - 1
	}
	for _, coin := range coins {
		for a := coin; a <= amount; a++ {
			dp[a] = min(dp[a], dp[a-coin]+1)
		}
	}
	if dp[amount] == math.MaxInt-1 {
		return -1
	}
	return dp[amount]
}

// 494. 目标和
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	sum -= target
	if sum < 0 || sum%2 != 0 {
		return 0
	}
	sum /= 2
	dp := make([]int, sum+1)
	dp[0] = 1
	for _, num := range nums {
		for t := sum; t >= num; t-- {
			dp[t] += dp[t-num]
		}
	}
	return dp[sum]
}

// 461. 汉明距离
func hammingDistance(x int, y int) int {
	x = x ^ y
	result := 0
	for x != 0 {
		if x%2 == 1 {
			result++
		}
		x /= 2
	}
	return result
}

// 448. 找到所有数组中消失的数字
func findDisappearedNumbers1(nums []int) []int {
	leng := len(nums)
	for _, num := range nums {
		if nums[(num-1)%leng] <= leng {
			nums[(num-1)%leng] += leng
		}
	}
	result := make([]int, 0)
	for i, num := range nums {
		if num <= leng {
			result = append(result, i+1)
		}
	}
	return result
}
func findDisappearedNumbers2(nums []int) []int {
	leng := len(nums)
	for i := 0; i < leng; {
		n1 := nums[i] - 1
		if nums[n1] != n1+1 {
			nums[i], nums[n1] = nums[n1], nums[i]
		} else {
			i++
		}
	}
	result := make([]int, 0)
	for i, num := range nums {
		if i != num-1 {
			result = append(result, i+1)
		}
	}
	return result
}

// 438. 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	plen := len(p)
	slen := len(s)
	if plen > slen {
		return nil
	}
	phash := [26]int{}
	shash := [26]int{}
	for _, v := range p {
		phash[int(v-'a')]++
	}
	for i := 0; i < plen; i++ {
		shash[int(s[i]-'a')]++
	}
	result := []int{}
	if feq(shash, phash) {
		result = append(result, 0)
	}
	for i := 1; i <= slen-plen; i++ {
		shash[int(s[i-1]-'a')]--
		shash[int(s[i+plen-1]-'a')]++
		if feq(shash, phash) {
			result = append(result, i)
		}
	}
	return result
}
func feq(shash [26]int, phash [26]int) bool {
	for i := 0; i < 26; i++ {
		if shash[i] != phash[i] {
			return false
		}
	}
	return true
}

// 437. 路径总和 III
// dfs
func pathSum1(root *tree.TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	// 定义：从当前指定节点开始，连续向下延伸满足 targetSum 的路径数量
	var rootSum func(*tree.TreeNode, int) int
	rootSum = func(node *tree.TreeNode, target int) int {
		if node == nil {
			return 0
		}
		count := 0
		if node.Val == target {
			count++
		}
		count += rootSum(node.Left, target-node.Val)
		count += rootSum(node.Right, target-node.Val)
		return count
	}

	return rootSum(root, targetSum) + pathSum1(root.Left, targetSum) + pathSum1(root.Right, targetSum)
}

// 前缀和
func pathSum2(root *tree.TreeNode, targetSum int) int {
	preSum := make(map[int]int)
	curSum := 0
	res := 0
	preSum[0] = 1
	var rootSum func(*tree.TreeNode)
	rootSum = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		curSum += node.Val
		res += preSum[curSum-targetSum]
		preSum[curSum]++
		rootSum(node.Left)
		rootSum(node.Right)
		preSum[curSum]--
		curSum -= node.Val
	}
	rootSum(root)
	return res
}

// 416. 分割等和子集
func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := range nums {
		for j := sum; j >= nums[i]; j-- {
			if dp[j-nums[i]] {
				dp[j] = true
			}
		}
		if dp[sum] {
			return true
		}
	}
	return false
}

// 406. 根据身高重建队列
func reconstructQueue(people [][]int) [][]int {
	slices.SortFunc(people, func(a, b []int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return b[0] - a[0]
	})
	result := make([][]int, 0, len(people))
	for _, p := range people {
		result = slices.Insert(result, p[1], p)
	}
	return result
}

// 399. 除法求值
type UnionFind struct {
	Parent []int
	Weight []float64
}

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{Parent: make([]int, n), Weight: make([]float64, n)}
	for i := range u.Parent {
		u.Parent[i] = i
		u.Weight[i] = 1.0
	}
	return u
}
func (u *UnionFind) find(x int) int {
	if x != u.Parent[x] {
		parentx := u.Parent[x]
		u.Parent[x] = u.find(u.Parent[x])
		u.Weight[x] = u.Weight[x] * u.Weight[parentx]
	}
	return u.Parent[x]
}
func (u *UnionFind) union(x int, y int, value float64) {
	findx := u.find(x)
	findy := u.find(y)
	if findx == findy {
		return
	}
	u.Parent[findx] = findy
	u.Weight[findx] = u.Weight[y] * value / u.Weight[x]
}
func (u *UnionFind) connectedWeight(x int, y int) float64 {
	findx := u.find(x)
	findy := u.find(y)
	if findx == findy {
		return u.Weight[x] / u.Weight[y]
	} else {
		return -1.0
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	leng := len(equations)
	unionFind := NewUnionFind(leng * 2)
	index := 0
	mapd := make(map[string]int)
	for i, e := range equations {
		s1 := e[0]
		s2 := e[1]
		if _, o1 := mapd[s1]; !o1 {
			mapd[s1] = index
			index++
		}
		if _, o2 := mapd[s2]; !o2 {
			mapd[s2] = index
			index++
		}
		unionFind.union(mapd[s1], mapd[s2], values[i])
	}
	qlen := len(queries)
	result := make([]float64, qlen)
	for i, q := range queries {
		s1 := q[0]
		s2 := q[1]
		key1, o1 := mapd[s1]
		key2, o2 := mapd[s2]
		if !o1 || !o2 {
			result[i] = -1
			continue
		}
		result[i] = unionFind.connectedWeight(key1, key2)
	}
	return result
}

// 394. 字符串解码
func decodeString(s string) string {
	slen := len(s)
	stackN := make([]int, 0, slen)
	stackS := make([]string, 0, slen)
	for i := 0; i < slen; {
		if s[i] >= '0' && s[i] <= '9' {
			j := i
			for s[j] >= '0' && s[j] <= '9' {
				j++
			}
			num, _ := strconv.Atoi(s[i:j])
			stackN = append(stackN, num)
			i = j
		} else if s[i] == ']' {
			j := len(stackS) - 1
			sub := ""
			for stackS[j] != "[" {
				j--
			}
			for i := j + 1; i < len(stackS); i++ {
				sub += stackS[i]
			}
			num := stackN[len(stackN)-1]
			r := ""
			for i := 0; i < num; i++ {
				r += sub
			}
			stackN = stackN[:len(stackN)-1]
			stackS = stackS[:j]
			stackS = append(stackS, r)
			i++
		} else {
			stackS = append(stackS, string(s[i]))
			i++
		}
	}
	return strings.Join(stackS, "")
}

// 347. 前 K 个高频元素
func topKFrequent1(nums []int, k int) []int {
	key := make([]int, 0, len(nums))
	value := make([]int, 0, len(nums))
	target := make(map[int]int)
	i := 0
	for _, n := range nums {
		if index, ok := target[n]; ok {
			value[index]++
		} else {
			key = append(key, n)
			value = append(value, 1)
			target[n] = i
			i++
		}
	}
	result := make([]int, k)
	for i = 0; i < k; i++ {
		for j := 0; j < len(value)-i-1; j++ {
			if value[j] > value[j+1] {
				value[j], value[j+1] = value[j+1], value[j]
				key[j], key[j+1] = key[j+1], key[j]
			}
		}
		result[i] = key[len(value)-i-1]
	}
	return result
}
func topKFrequent2(nums []int, k int) []int {
	num := make(map[int]int)
	maxt := 0
	for _, n := range nums {
		num[n]++
		maxt = max(maxt, num[n])
	}
	target := make(map[int][]int)
	for key, value := range num {
		target[value] = append(target[value], key)
	}
	result := make([]int, 0, k)
	for i := maxt; i > 0 && len(result) < k; i-- {
		result = append(result, target[i]...)
	}
	return result
}

// 338. 比特位计数
func countBits(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i/2] + 1
		}
	}
	return dp
}

// 337. 打家劫舍 III
func rob3(root *tree.TreeNode) int {
	var robdfs func(*tree.TreeNode) [2]int
	robdfs = func(node *tree.TreeNode) [2]int { // 0偷 1不偷
		if node == nil {
			return [2]int{0, 0}
		}
		left := robdfs(node.Left)
		right := robdfs(node.Right)
		n := [2]int{}
		n[0] = left[1] + right[1] + node.Val
		n[1] = max(left[0], left[1]) + max(right[0], right[1])
		return n
	}
	result := robdfs(root)
	return max(result[0], result[1])
}

// 121. 买卖股票的最佳时机
func maxProfit(prices []int) int {
	leng := len(prices)
	maxn := prices[leng-1]
	result := 0
	for i := leng - 2; i >= 0; i-- {
		maxn = max(maxn, prices[i])
		result = max(result, maxn-prices[i])
	}
	return result
}

// 309. 买卖股票的最佳时机含冷冻期
func maxProfit2(prices []int) int {
	leng := len(prices)
	dp := [3]int{} //0买入  1不持有 在冷冻期  2不持有 不在冷冻期
	dp[0] = -prices[0]
	for i := 1; i < leng; i++ {
		temp := dp[1]
		dp[1] = dp[0] + prices[i]
		dp[0] = max(dp[0], dp[2]-prices[i])
		dp[2] = max(temp, dp[2])
	}
	return max(dp[1], dp[2])
}

// 301. 删除无效的括号
func removeInvalidParentheses(s string) []string {
	leftremove := 0
	rightremove := 0
	m := ""
	for i := range s {
		if s[i] == '(' {
			leftremove++
		} else if s[i] == ')' {
			if leftremove == 0 {
				rightremove++
			} else {
				leftremove--
			}
		}
	}
	result := make(map[string]struct{}, 0)
	var dfs func(string, int, int, int)
	dfs = func(s string, index int, left int, right int) {
		if left == 0 && right == 0 {
			if checkInvalidParentheses(s) {
				result[s] = struct{}{}
			}
			return
		}
		for i := index; i < len(s); i++ {
			if left+right > len(s)-i {
				return
			}
			if s[i] == '(' && left > 0 {
				dfs(s[:i]+s[i+1:], i, left-1, right)
			} else if s[i] == ')' && right > 0 {
				dfs(s[:i]+s[i+1:], i, left, right-1)
			}
		}
	}
	dfs(s, 0, leftremove, rightremove)
	r := make([]string, len(result))
	if len(result) > 0 {
		i := 0
		for key := range result {
			r[i] = key
			i++
		}
	} else {
		r = append(r, m)
	}

	return r
}
func checkInvalidParentheses(s string) bool {
	left := 0
	for _, v := range s {
		if v == '(' {
			left++
		} else if v == ')' {
			if left <= 0 {
				return false
			}
			left--
		}
	}
	return left == 0
}

// 300. 最长递增子序列
func lengthOfLIS(nums []int) int {
	leng := len(nums)
	dp := make([]int, leng)
	dp[0] = 1
	result := 1
	for i := 1; i < leng; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}
	return result
}

// 287. 寻找重复数
func findDuplicate1(nums []int) int {
	m := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return n
		}
		m[n] = struct{}{}
	}
	return -1
}
func findDuplicate2(nums []int) int {
	fast, slow := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
