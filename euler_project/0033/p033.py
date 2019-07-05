def cal():
	ls = []
	for a in range(10, 99):
		for b in range(a + 1, 100):
			ta = [int(x) for x in str(a)]
			tb = [int(x) for x in str(b)]
			if ta[0] == tb[0] and tb[1]:
				if ta[1] * 1.0 / tb[1] == a * 1.0 / b:
					ls.append([a, b])
			elif ta[0] == tb[1] and tb[0]:
				if ta[1] * 1.0 / tb[0] == a * 1.0 / b:
					ls.append([a, b])
			elif ta[1] == tb[0] and tb[1]:
				if ta[0] * 1.0 / tb[1] == a * 1.0 / b:				
					ls.append([a, b])

	return ls
def gcd(a, b):
	if b == 0:
		return a
	return gcd(b, a % b)

def solve():
	ls = cal()
	a = 1
	b = 1
	for elem in ls:
		a *= elem[0]
		b *= elem[1]
	tmp = gcd(a, b)
	return b / tmp

print solve()