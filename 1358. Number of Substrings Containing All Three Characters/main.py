from collections import defaultdict


class Solution:

    def isValid(self, mdict):
        s = 0
        for w in self.mset:
            if mdict[w] > 0:
                s += 1
        return s == len(self.mset)

    def numberOfSubstrings(self, s: str) -> int:
        n = len(s)
        self.mset = set(["a", "b", "c"])
        mdict = defaultdict(int)
        res = 0

        i = 0
        j = 0
        while i < n:
            if self.isValid(mdict):
                res += (n - i) + 1

                while self.isValid(mdict):
                    mdict[s[j]] -= 1
                    j += 1

            mdict[s[i]] += 1
            i += 1

        while self.isValid(mdict):
            res += 1
            mdict[s[j]] -= 1
            j += 1

        return res


if __name__ == "__main__":
    s = Solution()
    print("abcabc = ", s.numberOfSubstrings("abcabc"))
    print("aaacb = ", s.numberOfSubstrings("aaacb"))
