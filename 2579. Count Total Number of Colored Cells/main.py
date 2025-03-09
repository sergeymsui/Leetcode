class Solution:
    def coloredCells(self, n: int) -> int:
        if n == 1:
            return 1
        return int((4*(n-1)) * n / 2) + 1

if __name__ == "__main__":
    s = Solution() 
    print(s.coloredCells(5))