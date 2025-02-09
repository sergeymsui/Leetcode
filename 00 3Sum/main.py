from collections import defaultdict

nums = [-1, 0, 1, 2, -1, -4]


class Solution:
    def threeSum(self, nums):
        seen = {}
        res = set()
        dups = set()

        i = 0
        while i < len(nums):

            if nums[i] not in dups:

                dups.add(nums[i])

                j = i + 1
                while j < len(nums):
                    d = 0 - nums[i] - nums[j]
                    if d in seen and seen[d] == i:
                        m_tray = [nums[i], nums[j], d]
                        m_tray.sort()
                        res.add(tuple(m_tray))

                    seen[nums[j]] = i
                    j += 1

            i += 1

        return [list(x) for x in res]


s = Solution()

print(s.threeSum(nums))
