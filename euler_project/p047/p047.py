def sieve(nmax):
	nmax += 1
	prime = []
	is_prime = [True] * nmax
	is_prime[0] = is_prime[1] = False
	cal_factors = [[]] * nmax

	for i in range(nmax):
		if is_prime[i]:
			prime.append(i)
			for j in range(i * 2, nmax, i):
				is_prime[j] = False
				cal_factors[j] = cal_factors[j] + [i]
	for i in range(nmax):
		cal_factors[i] = set(cal_factors[i])
	return cal_factors

def solve(nmax, cnt):
	cal_factors = sieve(nmax)

	for x in range(4, nmax - cnt):
		flag = True
		for i in range(cnt):
			if len(cal_factors[x + i]) != cnt:
				flag = False
				break
		if flag:
			return x
	return 0

print solve(100, 2)
print solve(1000, 3)
print solve(200000, 4)