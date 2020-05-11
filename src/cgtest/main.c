#include <stdio.h>
#include "libgoutil.h"

int main(int argc, char **argv) {
    printf("calling greet\n");
    Greet();

    printf("what is 1+2?\n");
    printf("--> %d\n", Add(1, 2));

    printf("what is 5-3?\n");
    printf("--> %d\n", Minus(5, 3));

    printf("bye!\n");
    return 0;
}
