//This program returns 1 is the given number is an armstrong number and 0
// otherwise.
#include "armstrong_numbers.h"
#include <math.h>


int isArmstrongNumber(int n)
{
  int exponent = floor(log10(n)) + 1;
  int number = n;
  int sum = 0;
  while(n > 0)
    {
      sum += pow(n % 10, exponent);
      n = n /10;
    }
  return number == sum;
}
