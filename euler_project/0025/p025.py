def solve(n):
	fibs = [1, 1]
	while True:
		tmp = fibs[-1] + fibs[-2]
		fibs.append(tmp)
		if len(str(tmp)) == n:
			return len(fibs)

print solve(1000)
