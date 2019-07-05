def solve(n):
	ans = 1
	for i in range(1, n + 1):
		ans *= i
	return sum([int(x) for x in str(ans)])

print solve(100)