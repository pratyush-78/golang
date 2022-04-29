#include<stdio.h>
#include<time.h>
#include<string.h>

int main()
{
    int a=10, b=30;

    time_t t = time(NULL);


    time_t t1;   // not a primitive datatype
    time(&t1);
    printf("%ld %ld\n", t,t1);

    char *dt = ctime(&t1);
    char darr[sizeof(char*)];
    memcpy(darr, dt, sizeof(char*));
    printf("%s", darr);


    printf("\nThis program has been writeen at (date and time): %s", ctime(&t));

    printf("\n\n\t\t\tCoding is Fun !\n\n\n");

    char * arr = "Hello";

    puts(dt);

    



    
    
    return 0;

}