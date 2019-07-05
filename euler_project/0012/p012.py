def solve(nmax):
	ls = [1 for n in range(nmax)]
	for i in range(2, nmax):
		for j in range(i, nmax, i):
			ls[j] += 1
	i = 1
	while True:
		tmp = 0
		# an = n * (n + 1) / 2
		if i % 2 == 0:
			tmp = ls[i / 2] * ls[i + 1]
		else:
			tmp = ls[i] * ls[(i + 1) / 2]
		if tmp > 500:
			return i * (i + 1) / 2
			break
		i += 1
		if i >= nmax:
			break
	return 

print solve(100000)
 