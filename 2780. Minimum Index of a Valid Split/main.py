from collections import defaultdict


class Solution:

    def minimumIndex(self, nums) -> int:
        m = len(nums)

        mlistval = list()
        mindexed = defaultdict(int)

        dom_count = 1
        dominated = nums[1]
        for n in nums:
            mindexed[n] += 1

            if dom_count < mindexed[n]:
                dominated = n
                dom_count = mindexed[n]
        s = 0
        for i, n in enumerate(nums):
            if n == dominated:
                s += 1
                mlistval.append((s, i + 1))

        for s, i in mlistval:
            if (i / 2) < s and ((m - i) / 2) < (dom_count - s):
                return i - 1

        return -1


if __name__ == "__main__":
    s = Solution()
    print(s.minimumIndex([1, 2, 2, 2]))
    print(s.minimumIndex([2, 1, 3, 1, 1, 1, 7, 1, 2, 1]))
