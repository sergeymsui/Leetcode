class Solution:
    def waysToSplitArray(self, nums: list) -> int:
        res = 0
        totalSum = sum(nums)

        i = 0
        ls = 0
        while i < len(nums) - 1:
            ls += nums[i]
            rs = totalSum - ls

            if ls >= rs:
                res += 1

            i += 1

        return res


def main():
    nums = [0, 1]

    s = Solution()
    print(s.waysToSplitArray(nums))


if __name__ == "__main__":
    main()
