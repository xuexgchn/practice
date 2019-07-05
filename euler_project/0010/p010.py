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
	return sum(prime)

print solve(2000000)			