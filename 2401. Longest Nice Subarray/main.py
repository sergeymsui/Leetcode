class Solution:

    def longestNiceSubarray(self, nums) -> int:
        n = len(nums)

        if n == 1:
            return 1

        max_arr = 1
        i = 0
        while i < n:
            j = i + 1
            mprx = nums[i]
            while j < n:

                if mprx & nums[j] == 0:
                    mprx = mprx | nums[j]
                    max_arr = max(max_arr, j - i + 1)
                else:
                    break
                j += 1
            i += 1

        return max_arr


if __name__ == "__main__":
    s = Solution()
    print(s.longestNiceSubarray([135745088, 609245787, 16, 2048, 2097152]))
