#include<stdio.h>
#include<math.h>

int multiply(int, int[], int);
void factorial(int);

int main() {
    int num;
    printf("Enter a number ..\n");
    scanf("%d", &num);

    factorial(num);
    return 0;
}

void factorial(int n) {
    int a[160];
    a[0] = 1;
    int len_res = 1;
    
    for(int i=2; i<=n; i++) {
      len_res =  multiply(i, a, len_res);
    }

    for(int i=len_res-1; i>=0; i--) {
       printf("%d", a[i]);
    }

    return;
}

int multiply(int n, int res[], int len_res) {
    int carry = 0;

    for(int i=0; i< len_res; i++)  {
        int prod = res[i] * n + carry;
	res[i] = prod % 10;
	carry = prod / 10;
    }

    while(carry) {
       res[len_res] = carry % 10;
       carry = carry/10;
       len_res++;
    }

    return len_res;
}




