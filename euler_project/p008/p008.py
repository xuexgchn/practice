def solve():
	f = open("in", "r")
	s = f.read()
	ls = s.split('\n')
	s = ""
	for i in ls:
		s += i
	ans = 0
	for i in range(len(s) - 13):
		tmp = 1
		for j in range(13):
			tmp *= int(s[i + j])
		if tmp >= ans:
			ans = tmp
	return ans

print solve()