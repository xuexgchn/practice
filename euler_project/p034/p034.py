def factorial(n):
	fac = [1]
	for x in range(1, n + 1):
		fac.append(fac[-1] * x)
	return fac 

def solve(nmax):
	fac = factorial(nmax)
	cnt = 0

	for x in range(3, nmax):
		res = sum([fac[int(i)] for i in str(x)])
		if res == x:
			cnt += x

	return cnt

print solve(50000)