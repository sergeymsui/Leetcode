from collections import defaultdict


class Solution:
    def distinctNumbers(self, nums, k):
        n = len(nums)
        result = list()

        unique = 0
        window_size = 0
        mdict = defaultdict(int)

        for i in range(n):
            if nums[i] not in mdict or mdict[nums[i]] == 0:
                unique += 1

            mdict[nums[i]] += 1
            window_size += 1

            if window_size > k:
                if mdict[nums[i - k]] == 1:
                    unique -= 1
                mdict[nums[i - k]] -= 1
                window_size -= 1

            if window_size == k:
                result.append(unique)

        return result


if __name__ == "__main__":
    s = Solution()
    print(s.distinctNumbers([1, 2, 3, 2, 2, 1, 3], 3))
