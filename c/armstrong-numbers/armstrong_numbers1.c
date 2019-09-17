#include <stdio.h>
#include <math.h>

int main()
{
  int n = 9475;
  int count = 0;
  unsigned int sum = 0;
  int tmp = n;
  
  while (tmp /= 10) {
    count++;
  }
  count++;
  int dig[count];
  int digSize = count;
  printf("count %d\n", digSize);
  tmp = n;
  while(count--)
    {
      dig[count] = tmp % 10;
      tmp /= 10;
    }
  for(int i = 0; i < digSize; i++)
    {
      printf("dig[i] %d\n", dig[i]);
      sum += pow(dig[i], digSize);
    }
  printf("sum %d\n", sum);
  if ( sum == n )
    {
      printf("True");
    }
  else
    {
      printf("False");
    }
  return 0;
}
