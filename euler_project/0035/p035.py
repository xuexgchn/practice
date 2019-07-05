def Prime(nmax):
	prime = []
	is_prime = [True] * nmax 
	is_prime[0] = is_prime[1] = False
	for i in range(2, nmax):
		if is_prime[i]:
			prime.append(i)
			for j in range(2 * i, nmax, i):
				is_prime[j] = False
	# print prime[0]
	return prime, is_prime


def solve(nmax):
	prime, is_prime = Prime(nmax)
	cnt = 0
	for num in prime:
		flag = True
		s = str(num)
		length = len(s)
		for i in range(length):
			tmp = int(s[i:] + s[:i])
			if not is_prime[tmp]:
				flag = False
				break
		if flag:
			cnt += 1

	return cnt


print solve(1000000)