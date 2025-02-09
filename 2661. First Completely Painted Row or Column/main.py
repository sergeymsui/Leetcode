from collections import defaultdict


class Solution:
    def firstCompleteIndex(self, arr, mat) -> int:

        m = len(mat)
        n = len(mat[0])

        mdict = dict()
        for i, u in enumerate(mat):
            for j, v in enumerate(mat[i]):
                mdict[mat[i][j]] = [i, j]

        rows = defaultdict(list)
        cols = defaultdict(list)

        for i, u in enumerate(arr):
            r, c = mdict[u]
            rows[r].append(u)
            cols[c].append(u)

            if len(rows[r]) == n or len(cols[c]) == m:
                return i

        return -1


def main():
    arr = [1, 3, 4, 2]
    mat = [[1, 4], [2, 3]]
    s = Solution()
    print(s.firstCompleteIndex(arr, mat))


if __name__ == "__main__":
    main()
