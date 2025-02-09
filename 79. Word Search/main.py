class Solution:

    def find(self, board, i, j, word, history):

        if len(word) == 0:
            return True

        r1 = False
        r2 = False
        r3 = False
        r4 = False

        if (
            i + 1 < len(board)
            and board[i + 1][j] == word[0]
            and [i + 1, j] not in history
        ):
            r1 = self.find(board, i + 1, j, word[1:], history + [[i + 1, j]])

        if (
            j + 1 < len(board[i])
            and board[i][j + 1] == word[0]
            and [i, j + 1] not in history
        ):
            r2 = self.find(board, i, j + 1, word[1:], history + [[i, j + 1]])

        if i - 1 >= 0 and board[i - 1][j] == word[0] and [i - 1, j] not in history:
            r3 = self.find(board, i - 1, j, word[1:], history + [[i - 1, j]])

        if j - 1 >= 0 and board[i][j - 1] == word[0] and [i, j - 1] not in history:
            r4 = self.find(board, i, j - 1, word[1:], history + [[i, j - 1]])

        return r1 or r2 or r3 or r4

    def exist(self, board: list, word: str) -> bool:

        for i in range(len(board)):
            for j in range(len(board[i])):
                if board[i][j] == word[0]:
                    r = self.find(board, i, j, word[1:], [[i, j]])

                    if r:
                        return True

        return False


board = [["A", "B", "C", "E"], ["S", "F", "E", "S"], ["A", "D", "E", "E"]]
word = "ABCESEEEFS"

s = Solution()

res = s.exist(board, word)
print(res)
