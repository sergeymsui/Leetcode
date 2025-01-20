class Solution:
    def xorAllNums(self, nums1, nums2) -> int:

        s = None

        for n in nums1:
            for m in nums2:
                if s == None:
                    s = n ^ m
                else:
                    s = s ^ n ^ m

        return s


def main():
    nums1 = [1, 2]
    nums2 = [3, 4]
    s = Solution()
    print(s.xorAllNums(nums1, nums2))


if __name__ == "__main__":
    main()
