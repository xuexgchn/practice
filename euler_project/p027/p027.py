def sieve(nmax):
	is_prime = [True] * nmax
	prime = []
	is_prime[0] = is_prime[1] = False
	for i in range(nmax):
		if is_prime[i]:
			prime.append(i)
			for j in range(i + i, nmax, i):
				is_prime[j] = False

	return prime, is_prime

def solve():
	ans = (0, 0, 0)
	prime, is_prime = sieve(100000)

	for a in range(-999, 1000):
		for b in range(max(2, 1 - a), 1000):
		# tips :
		# b >= 2; a + b + 1 >= 2
			n, cnt = 0, 0
			while True:
				v = n * n + a * n + b
				if is_prime[v]:
					cnt += 1
				else:
					break
				n += 1
			if cnt > ans[2]:
				ans = (a, b, cnt)

	return ans[0] * ans[1]

print solve()