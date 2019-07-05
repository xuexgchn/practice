def solve(n):
	return sum([int(x) for x in str(2 ** n)])

print solve(15)
print solve(1000)