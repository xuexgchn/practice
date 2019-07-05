def solve(nmax):
	pn = [x * (3 * x - 1) / 2 for x in range(1, nmax)]
	st = set(pn)
	ans = []
	for i in range(nmax - 1):
		for j in range(i + 1, nmax - 1):
			if (pn[i] + pn[j]) < pn[nmax - 2] and (pn[i] + pn[j]) in st and (pn[j] - pn[i]) in st:
				# print pn[i], pn[j]
				ans.append(pn[j] - pn[i])

	if ans:
		return min(ans)
	else:
		return 0

print solve(3000)