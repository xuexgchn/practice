def solve(nmax, cnt):
	for x in range(1, nmax):
		real = sorted(str(x))
		flag = True
		for i in range(1, cnt + 1):
			t = x * i
			tmp = sorted(str(t))
			if tmp != real:
				flag = False
				break
		if flag:
			return x
	return 0

print solve(1000000, 6)
