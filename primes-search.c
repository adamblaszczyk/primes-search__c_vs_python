/*
 *  -----------------------------------------
 *  Primes Search - C version
 *  -----------------------------------------
 *  Author: Adam Blaszczyk
 *          http://wyciekpamieci.blogspot.com
 *  Date:   2022-06-23
 *  -----------------------------------------
 *  
 *  Compilation (Linux, Windows):
 *         gcc primes-search.c -o primes-search -lm
 *  Compilation (FreeBSD):
 *         cc primes-search.c -o primes-search -lm
 *
 *  Usage:
 *         primes-search
 *
 */
 
#include <stdio.h>
#include <math.h>
#include <time.h>

#define FALSE 0
#define TRUE  1
  
int is_prime(int n) {
  if (n<2) {
    return FALSE;
  }
  else if (n==2) {
    return TRUE;
  }
  else {
    for (int i=2; i<=(int)ceil(sqrt(n)); i++) {
      if (n % i == 0) {
        return FALSE;
      }
    }
    return TRUE;
  }
}

int main() {
  int max = 8000000;
  time_t t1, t2;
  
  printf("Checking %d numbers. Please wait...\n", max);
  
  t1 = time(NULL);
  
  for(int i=1; i<=max; i++) {
    is_prime(i);
    //printf("%d   %d\n", i, is_prime(i));
  }
  
  t2 = time(NULL);
  printf("Done in %ld seconds.\n", t2-t1);
}

