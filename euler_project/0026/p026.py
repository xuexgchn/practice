def cal(d):
	for t in range(1, d):
		if 10 ** t % d == 1:
			return t, d
	return 0

def solve(nmax):
	ans = [cal(x) for x in range(2, nmax + 1)]
	res = max(ans)[1]

	return res

print solve(1000)
