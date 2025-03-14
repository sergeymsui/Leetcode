class Solution:

    def numberOfAlternatingGroups(self, colors, k) -> int:
        n = len(colors)
        uvalid = list()
        s = 0

        i = 0
        while i + 1 < n:
            uvalid.append((colors[i] == colors[i + 1]))
            i += 1

        uvalid.append((colors[-1] == colors[0]))
        uvalid += uvalid[:k]

        badpos = -1

        for i in range(k):
            if uvalid[i]:
                badpos = i

        for i in range(n):
            if not i < badpos:
                s += 1

            if uvalid[i + k]:
                badpos = i + k

        return s


if __name__ == "__main__":
    s = Solution()
    print(s.numberOfAlternatingGroups([0, 1, 0, 1, 0], 3))
    print(s.numberOfAlternatingGroups([0, 1, 0, 0, 1, 0, 1], 6))
    print(s.numberOfAlternatingGroups([0, 1, 1], 3))

    print(s.numberOfAlternatingGroups([0, 1, 0, 1, 1], 5))
