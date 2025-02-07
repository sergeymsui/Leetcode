class Solution:
    def longestMonotonicSubarray(self, nums) -> int:

        left_side = 0
        right_side = 0

        i = 0
        while i < len(nums):

            j = i + 1
            while j < len(nums) and nums[j - 1] < nums[j]:
                j += 1

            if right_side - left_side < j - i:
                left_side = i
                right_side = j

            j = i + 1
            while j < len(nums) and nums[j - 1] > nums[j]:
                j += 1

            if right_side - left_side < j - i:
                left_side = i
                right_side = j

            i += 1

        return right_side - left_side


if __name__ == "__main__":
    s = Solution()
    print(s.longestMonotonicSubarray([1, 4, 3, 3, 2]))
