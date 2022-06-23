#
# ---------------------------------------------
# Primes Search - Python3 version
# ---------------------------------------------
# Author: Adam Blaszczyk
#         http://wyciekpamieci.blogspot.com
# Date:   2022-06-23
# ---------------------------------------------
#
# Requirements:
#         - Python 3.x
#

import math
import time

def is_prime(n):
  if n < 2:
    return False
  elif n == 2:
    return True
  else:
    # range = 2..math.ceil(math.sqrt(n))
    for i in range(2, (int)(math.ceil(math.sqrt(n))) + 1):
      if n % i == 0:
        return False
    return True
    

max = 8000000

print("Checking %d numbers. Please wait..." % max)

t1 = time.time()

for i in range(1, max):
  is_prime(i)
  #print("%d   %d" % (i, is_prime(i)))
  
t2 = time.time()
print("Done in %d seconds." % (t2-t1)) 
