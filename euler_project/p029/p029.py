def solve(a, b):
	ans = []
	for i in range(2, a + 1):
		for j in range(b - 1):
			if j == 0:
				ans.append(i * i)
			else:
				ans.append(ans[-1] * i)
	return len(set(ans))

print solve(100, 100)