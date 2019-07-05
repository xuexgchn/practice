def sieve(nmax):
	nmax += 1
	is_prime = [True] * nmax
	is_prime[0] = is_prime[1] = False
	prime = []

	for i in range(nmax):
		if is_prime[i]:
			prime.append(i)
			for j in range(i * 2, nmax, i):
				is_prime[j] = False
	return is_prime

def solve(nmax):
	cnt = 3
	ans = []
	is_prime = sieve(nmax)
	for x in range(1000, nmax - cnt):
		if is_prime[x]:
			check = list(str(x))
			check.sort()
			flag = True
			t = x
			for i in range(1, cnt): 
				t += 3330
				if t > nmax:
					flag = False
					break
				if is_prime[t]:
					tmp = list(str(t))
					tmp.sort()
					if tmp != check:
						flag = False
						break
				else:
					flag = False
					break
			if flag:
				ans.append(x)
	res = []
	if ans:
		for x in ans:
			tmp = ""
			for i in range(cnt):
				tmp += str(x)
				x += 3330
			res.append(tmp)
	
	return res
print solve(10000)