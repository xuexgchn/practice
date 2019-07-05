import math

def Perms(numbers):
	if len(numbers) <= 1:
		return [numbers]
	res = []
	for i in range(len(numbers)): 
		tmp = numbers[:i] + numbers[i + 1:]
		p = Perms(tmp)
		for number in p:
			res.append(numbers[i: i + 1] + number)
	return res

def is_prime(n):
	for i in range(2, int(math.sqrt(n)) + 1):
		if n % i == 0:
			return False
	return True

def solve(nmax):
	ls = range(1, nmax + 1)
	perms = Perms(ls)
	ans = []
	for perm in perms:
		tmp = [str(i) for i in perm]
		tmp = ''.join(tmp)
		if is_prime(int(tmp)):
			ans.append(tmp)
	if ans:
		return max(ans)
	else:
		return 0

print solve(4)
print solve(5)
print solve(6)
print solve(7)
print solve(8)



