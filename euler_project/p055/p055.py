def cal_cnt(n):
	cnt = 1
	tmp = int(str(n)) + int(str(n)[::-1])
	while (str(tmp) != str(tmp)[::-1]):
		n = tmp
		tmp = int(str(n)) + int(str(n)[::-1])
		cnt += 1
		if cnt > 50:
			return True
	return False

def solve(nmax):
	cnt = 0
	for x in range(1, nmax + 1):
		if cal_cnt(x):
			cnt += 1
	return cnt

print solve(10000)
