def sieve(nmax):
	nmax += 1
	is_prime = [True] * nmax
	is_prime[0], is_prime[1] = False, False
	prime = []

	for i in range(nmax):
		if is_prime[i]:
			prime.append(i)
			for j in range(i * 2, nmax, i):
				is_prime[j] = False
	return prime, is_prime

def solve(nmax, test):
	prime, is_prime = sieve(nmax)
	largest = 0
	left, right = 0, 0
	for i in range(test):
		tmp = prime[i:]
		total, j = 0, 0
		for x in tmp:
			total += x
			if total > nmax:
				break
			j += 1
			if total in prime:
				if j > right - left:
					left = i
					right = i + j
	return sum(prime[left:right])

print solve(100, 10)
print solve(1000, 10)
print solve(1000000, 10)