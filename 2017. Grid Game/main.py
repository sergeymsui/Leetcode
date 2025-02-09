class Solution:

    def maxSumPath(self, grid):
        path_len = len(grid[0])

        max_sum, max_path = 0, 0

        for i in range(path_len):
            s = sum(grid[0][: i + 1]) + sum(grid[1][i:])

            if max_sum < s:
                max_sum = s
                max_path = i

        return max_sum, max_path

    def gridGame(self, grid) -> int:
        _, max_path = self.maxSumPath(grid)

        grid[0][: max_path + 1] = [0] * (max_path + 1)
        grid[1][max_path:] = [0] * (len(grid[0]) - max_path)

        max_sum, _ = self.maxSumPath(grid)

        return max_sum


def main():
    grid = [[20, 3, 20, 17, 2, 12, 15, 17, 4, 15], [20, 10, 13, 14, 15, 5, 2, 3, 14, 3]]
    s = Solution()
    print(s.gridGame(grid))


if __name__ == "__main__":
    main()
