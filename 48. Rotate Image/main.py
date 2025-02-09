class Solution:
    def transpose(self, matrix: list):
        n = len(matrix)
        i = 0
        while i < n:
            j = i
            while j < n:
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
                j += 1
            i += 1
        pass

    def reverse(self, matrix: list):
        n = len(matrix)

        i = 0
        while i < n:
            l = 0
            k = n - 1
            while l < k:
                matrix[i][l], matrix[i][k] = matrix[i][k], matrix[i][l]
                l += 1
                k -= 1
            i += 1

    def rotate(self, matrix: list):
        self.transpose(matrix)
        self.reverse(matrix)


def main():
    matrix = [[5, 1, 9, 11], [2, 4, 8, 10], [13, 3, 6, 7], [15, 14, 12, 16]]

    s = Solution()
    s.rotate(matrix)

    for line in matrix:
        print(*line)


if __name__ == "__main__":
    main()
