class Solution:
    def findDifferentBinaryString(self, nums):
        n = len(nums)
        nset = set(nums)
        bstr = [0] * n
        rstring = "".join(map(str, bstr))

        for i in range(n):
            right_border = n - i - 1

            for right_index in range(right_border, -1, -1):
                dcopy = [*bstr]
                dcopy[right_index] = 1
                djoin = "".join(map(str, dcopy))

                if djoin not in nset:
                    return djoin

            bstr[right_border] = 1

        return rstring


if __name__ == "__main__":
    s = Solution()
    print(s.findDifferentBinaryString(["1"]))
