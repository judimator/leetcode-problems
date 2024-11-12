# The problem

You are given a **0-indexed** array of integers `nums` of length `n`. You are initially positioned at `nums[0]`.

Each element `nums[i]` represents the maximum length of a forward jump from index `i`. In other words, if you are at `nums[i]`, you can jump to any `nums[i + j]` where:

- 0 <= j <= nums[i] and
- i + j < n

Return _the minimum number of jumps_ to reach `nums[n - 1]`. The test cases are generated such that you can reach `nums[n - 1]`.

### Example 1:

**Input:** nums = [2,3,1,1,4]  
**Output:** 2
