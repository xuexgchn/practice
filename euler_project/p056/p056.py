def cal(a, b):
	tmp = a ** b
	ans = [int(i) for i in str(tmp)]

	return sum(ans)

def solve(nmax):
	nmax += 1
	ans = 0
	for a in range(1, nmax):
		for b in range(1, nmax):
			tmp = cal(a, b)
			if tmp > ans:
				ans = tmp

	return ans

print solve(100)

