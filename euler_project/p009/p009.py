def solve():
	for a in range(1, 1000):
		for b in range(1, 1000):
			if a * a + b * b == (1000 - a - b) ** 2:
				return a * b * (1000 - a- b)
print solve()