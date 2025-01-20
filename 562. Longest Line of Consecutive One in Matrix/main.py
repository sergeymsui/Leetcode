!!!

class Solution:
    def longestLine(self, mat) -> int:
        h = len(mat)
        w = len(mat[0])

        horizontal = [0] * w
        vertical = [0] * h

        for i in range(0, h):
            for j in range(0, w):


def main():
    mat = [[0, 1, 1, 0], [0, 1, 1, 0], [0, 0, 0, 1]]
    s = Solution()
    print(s.longestLine(mat))


if __name__ == "__main__":
    main()
