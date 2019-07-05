def solve(nmax):
	ans = []
	vs = set([str(x) for x in range(1, 10)])
	for x in range(5, nmax):
		ls = []
		for i in range(1, 10):
			ls = ls + list(str(x * i))
			if len(set(ls)) != len(ls):
				break
			if set(ls) == vs and i > 1:
				ans.append([x, i, ls])
				break
	return ''.join(max(ans)[2])

print solve(1000000)