def solve(nmax):
	nmax += 1
	cal = [0] * nmax
	large = nmax / 2
	for a in range(1, large):
		for b in range(a, large):
			for c in range(b, large):
				if a + b + c > nmax:
					break
				if a ** 2 + b ** 2 == c ** 2:
					cal[a + b + c] += 1
	ans = 0
	cnt = 0
	for i in range(0, nmax):
		if cnt < cal[i]:
			cnt = cal[i]
			ans = i
	return ans

print solve(1000)