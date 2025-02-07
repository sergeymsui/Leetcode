class Solution:
    def tupleSameProduct(self, nums) -> int:

        nums.sort()
        result = 0
        n = len(nums)

        for i in range(n):
            for j in range(n - 1, i, -1):
                p = nums[i] * nums[j]

                range_set = set()
                for l in range(i + 1, j):
                    k = nums[l]

                    if p % k == 0:
                        if (p / k) in range_set:
                            result += 8
                        range_set.add(k)

        return result


if __name__ == "__main__":
    s = Solution()
    print(s.tupleSameProduct([1, 2, 4, 5, 10]))
