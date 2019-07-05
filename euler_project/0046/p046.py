import math

def sieve(nmax):
	nmax += 1
	is_prime = [True] * nmax
	prime = []
	is_prime[0] = is_prime[1] = False
	for i in range(nmax):
		if is_prime[i]:
			prime.append(i)
			for j in range(i * 2, nmax, i):
				is_prime[j] = False

	return prime,is_prime

def solve(nmax):
	prime, is_prime = sieve(nmax)
	for x in range(7, nmax, 2):
		if not is_prime[x]:
			flag = False
			for i in prime:
				if i >= x:
					break
				tmp = math.sqrt((x - i) / 2)
				if tmp == math.ceil(tmp):
					flag = True
			if not flag:
				return x
	return 0

print solve(10000)