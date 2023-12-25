import z3
import sys

# constraints variables (i.e. x, y, z and velocities for the rock throw we're trying to find)
x = z3.Int('x')
y = z3.Int('y')
z = z3.Int('z')
vx = z3.Int('vx')
vy = z3.Int('vy')
vz = z3.Int('vz')

s = z3.Solver()
i = 0
for line in sys.stdin:
    x_hs_0, y_hs_0, z_hs_0, x_hs_v, y_hs_v, z_hs_v = line.strip().split(" ")
    t = z3.Int('t' + str(i))
    s.add(x + vx * t == x_hs_0 + x_hs_v * t)
    s.add(y + vy * t == y_hs_0 + y_hs_v * t)
    s.add(z + vz * t == z_hs_0 + z_hs_v * t)
    i += 1

r = s.check()
if r == z3.unsat:
    print('no solution', file=sys.stderr)
    sys.exit(1)
elif r == z3.unknown:
    print('failed to solve', file=sys.stderr)
    sys.exit(1)
m = s.model()
print(m[x], m[y], m[z])
