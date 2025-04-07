from functools import lru_cache


class Solution:
    def canPartition(self, nums) -> bool:
        msum = sum(nums)

        @lru_cache(maxsize=None)
        def bft(nums, summa=0):

            if summa == 0:
                return True

            if len(nums) == 0 or summa < 0:
                return False

            lf, rf = False, False
            if len(nums):
                lf = bft(nums[1:], summa - nums[0])
                rf = bft(nums[1:], summa)

            return lf or rf

        if msum % 2 > 0:
            return False

        return bft(tuple(nums), msum // 2)


if __name__ == "__main__":
    s = Solution()
    print(s.canPartition(tuple([1, 2, 3, 4, 5, 6, 7])))
    print(s.canPartition(tuple([1, 2, 3, 5])))
    print(s.canPartition(tuple([9, 1, 2, 4, 10])))
