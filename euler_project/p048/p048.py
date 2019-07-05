def solve(nmax):
	ans = [x ** x for x in range(1, nmax + 1)]
	return str(sum(ans))[-10:]

# print solve(10)
print solve(1000)