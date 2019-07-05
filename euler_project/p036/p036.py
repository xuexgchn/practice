def solve(nmax):
	sum = 0
	for x in range(1, nmax + 1):
		if str(x) == str(x)[::-1]:
			tmp = bin(x)[2:]
			if tmp == tmp[::-1]:
				sum += x
	return sum

print solve(1000000)
		 	
