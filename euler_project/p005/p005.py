def gcd(a, b):
	if a % b == 0:
		return b
	return gcd(b, a % b)

assert gcd(10, 8) == 2

ans = 1;
for i in range(2, 21):
	ans = ans * i / gcd(ans, i)

print ans