def Prime(nmax):
	is_prime = [True] * nmax
	prime = []
	is_prime[0] = is_prime[1] = False
	for i in range(nmax):
		if is_prime[i]:
			prime.append(i)
			for j in range(i + i, nmax, i):
				is_prime[j] = False
	return prime, is_prime

def solve(nmax):
	ans = 0
	prime, is_prime = Prime(nmax)
	cnt = 0
	for x in prime:
		s1 = str(x)
		s2 = str(x)
		flag = True
		while s1 and s2:
			if not is_prime[int(s1)] or not is_prime[int(s2)]:
				flag = False
				break
			s1 = s1[1:]
			s2 = s2[:-1]
		if flag and x > 10:
			ans += x
			cnt += 1

	if cnt == 11:
		return ans
	else:
		return cnt

print solve(1000000)
