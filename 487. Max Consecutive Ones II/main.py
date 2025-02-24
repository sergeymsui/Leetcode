class Solution:
    def findMaxConsecutiveOnes(self, nums) -> int:

        msum = []
        i = 0
        while i < len(nums):
            j = i + 1

            if nums[i] > 0:
                while j < len(nums) and nums[j] > 0:
                    j += 1
                msum.append(sum(nums[i : j + 1]))
            else:
                msum.append(nums[i])
            i = j

        if len(msum) <= 3:
            if 0 in nums:
                return sum(msum) + 1
            else:
                return sum(msum)

        max_sum = msum[0]
        i = 0
        while i + 3 < len(msum) - 2:
            psum = (
                sum(msum[i : i + 3]) + 1
                if 0 in msum[i : i + 3]
                else sum(msum[i : i + 3])
            )
            max_sum = max(max_sum, psum)
            i += 1

        return max_sum


if __name__ == "__main__":
    s = Solution()
    print(s.findMaxConsecutiveOnes([1, 0, 1, 0, 1, 1, 0, 1, 0]))
    print(s.findMaxConsecutiveOnes([1] * 5000 + [0, 0] + [1] * 4621))
