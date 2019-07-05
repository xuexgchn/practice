def solve(max):
	prime = []
	is_prime = [True] * max 
	is_prime[0] = is_prime[1] = False
	for i in range(2, max):
		if is_prime[i]:
			prime.append(i)
			for j in range(2 * i, max, i):
				is_prime[j] = False
	# print prime[0]
	if len(prime) > 10000:
		print prime[10000]

solve(110000)			