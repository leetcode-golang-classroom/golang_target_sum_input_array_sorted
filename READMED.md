# golang_target_sum_input_array_sorted

Given a **1-indexed** array of integers `numbers` that is already ***sorted in non-decreasing order***, find two numbers such that they add up to a specific `target` number. Let these two numbers be `numbers[index1]` and `numbers[index2]` where `1 <= index1 < index2 <= numbers.length`.

Return *the indices of the two numbers,* `index1` *and* `index2`*, **added by one** as an integer array* `[index1, index2]` *of length 2.*

The tests are generated such that there is **exactly one solution**. You **may not** use the same element twice.

Your solution must use only constant extra space.

## Examples

**Example 1:**

```
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

```

**Example 2:**

```
Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

```

**Example 3:**

```
Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

```

**Constraints:**

- `2 <= numbers.length <= $3 * 10^4$`
- `1000 <= numbers[i] <= 1000`
- `numbers` is sorted in **non-decreasing order**.
- `1000 <= target <= 1000`
- The tests are generated such that there is **exactly one solution**.

## 解析

給定一個由小到大排序過的整數陣列 numbers. 給定一個整數 target

要求實作一個演算法找出 numbers 中兩個數相加合的 target

並且要求空間複雜度必須要是 O(1)

因為是排序過的整數陣列

所以可以透過兩個 pointer , lp =0, rp = len(numbers) -1

當 lp < rp 時做以下判斷

當 numbers[lp] + numbers[rp] == target  則找到這兩個index 回傳 [lp + 1 , rp +1 ] // 因為題目要求是 1-indexed 的座標 而 golang 與 java 都是 0-indexed

當 numbers[lp] + numbers[rp] > target 

因為左界已經到最小值，所以要讓兩數相加變小只能把右界向左移 更新 rp -=1 

當 numbers[lp] + numbers[rp] < target  

因為左界已經到最大值，所以要讓兩數相加變大只能把左界向右移 更新 lp +=1

當跑到最後都找不到則回傳 nil

詳細作法如下圖

![](https://i.imgur.com/knDWRGn.png)

 
## 程式碼
```go
package sol

func twoSum(numbers []int, target int) []int {
	lp, rp := 0, len(numbers)-1
	for lp < rp {
		if numbers[lp]+numbers[rp] == target {
			return []int{lp + 1, rp + 1}
		}
    if numbers[lp]+numbers[rp] > target {
      rp--
    }
		if numbers[lp]+numbers[rp] < target {
			lp++
		}
	}
	return nil
}

```

## 困難點

1. 要想出左右指標移動的條件

## Solve Point

- [x]  初始化 lp = 0, rp = len(numbers) - 1
- [x]  當 lp < rp 時 做以下運算
- [x]  當 numbers[lp] + numbers[rp] == target 則回傳 [lp+1, rp+1]
- [x]  當 numbers[lp] + numbers[rp] > target 則更新 rp -= 1
- [x]  當 numbers[lp] + numbers[rp] < target 則更新 lp += 1
- [x]  當跑到 lp = rp  時 , 回傳 nil 代表沒有找到符合的 pair