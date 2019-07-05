import math

def is_prime(x):
	for i in range(2, x):
		if x % i == 0:
			return False
	return True

def solve(x):
	t = math.sqrt(x) + 1
	for i in range(int(t), 1, -1):
		if x % i == 0:
			if is_prime(i):
				print i
				return ;
# solve(13195)
# solve(600851475143)