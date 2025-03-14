class Solution:
    def minZeroArray(self, nums, queries) -> int:
        n = len(nums)
        m = len(queries)

        negatives = 0
        for v in nums:
            if v == 0:
                negatives += 1

        if negatives == n:
            return 0

        i = 0
        while i < m:
            start, end, value = queries[i]
            j = start
            while j <= end:
                if nums[j] > 0 and (nums[j] - value) <= 0:
                    negatives += 1
                
                nums[j] -= value
                if negatives == n:
                    return i + 1
                
                j += 1
            i += 1

        return -1


if __name__ == "__main__":
    s = Solution()
    print(s.minZeroArray([2, 0, 2], [[0, 2, 1], [0, 2, 1], [1, 1, 3]]))
    print(s.minZeroArray([4, 3, 2, 1], [[1, 3, 2], [0, 2, 1]]))
    print(s.minZeroArray([10, 3], [[1, 1, 3], [1, 1, 2], [1, 1, 1], [0, 1, 2]]))
