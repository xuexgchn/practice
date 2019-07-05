def solve(nmax):
	ans = []
	for i in range(1, nmax):
		if (i % 2 == 0):
			ans.insert(0, i * i + 1)
		else:
			ans.insert(0, i ** 2)
		ans.append(ans[-1] + 2 * i)
	ans.append(nmax * nmax)
	return sum(ans)

print solve(1001)