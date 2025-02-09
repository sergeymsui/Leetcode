class Solution:
    def check(self, nums) -> bool:

        i = 0
        j = i + 1
        while j < len(nums):
            if nums[i] > nums[j]:
                break
            i += 1
            j += 1

        nums = nums[j:] + nums[:j]

        f = True
        i = 0
        j = i + 1
        while j < len(nums):
            f = f and (nums[i] <= nums[j])
            i += 1
            j += 1

        return f


if __name__ == "__main__":
    s = Solution()
    print(s.check([3, 4, 5, 1, 2]))
