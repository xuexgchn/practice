def dfs(target, coins):
	if target == 0 or len(coins) == 1:
		return 1
	else:
		coins = sorted(coins)
		large = coins[-1]
		cnt = target / large + 1
		ans = 0
		for i in range(cnt):
			ans += dfs(target - large * i, coins[:-1])
		return ans

print dfs(200, [1, 2, 5, 10, 20, 50, 100, 200])

# how to solve by using dp??