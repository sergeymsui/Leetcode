def max_flowers(n, m, field):

    dp = [[0] * m for _ in range(n)]

    dp[0][0] = int(field[0][0])

    for j in range(1, m):
        dp[0][j] = dp[0][j - 1] + int(field[0][j])

    for i in range(1, n):
        dp[i][0] = dp[i - 1][0] + int(field[i][0])

    for i in range(1, n):
        for j in range(1, m):
            dp[i][j] = max(dp[i - 1][j], dp[i][j - 1]) + int(field[i][j])

    return dp[n - 1][m - 1]


n, m = map(int, input().split())
field = [input().strip() for _ in range(n)]
field.reverse()


print(max_flowers(n, m, field))
