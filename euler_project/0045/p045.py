def solve(nmax):
	pn = set([x * (3 * x - 1) / 2 for x in range(nmax)])
	hn = set([x * (2 * x - 1) for x in range(nmax)])

	for x in range(nmax):
		tmp = x * (x + 1) / 2
		if tmp in pn and tmp in hn and tmp > 40755:
			return tmp

	return 0

print solve(100000)