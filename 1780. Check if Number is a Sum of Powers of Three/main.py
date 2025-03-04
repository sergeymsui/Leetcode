class Solution:
    def checkPowersOfThree(self, n: int) -> bool:
        powers = [1]

        while powers[-1] <= n:
            powers.append(powers[-1] * 3)

        mused = [True] * len(powers)
        index = len(powers) - 1

        while n > 0 and index >= 0:
            while powers[index] > n:
                index -= 1

            if mused[index]:
                n -= powers[index]
                mused[index] = False

            index -= 1

        return n == 0


if __name__ == "__main__":
    s = Solution()
    print(s.checkPowersOfThree(12))
    print(s.checkPowersOfThree(91))
    print(s.checkPowersOfThree(21))
