#include <stdio.h>
#include <math.h>
#include <assert.h>

int main() {
    int n, k;
    scanf("%d %d", &n, &k);
    // printf("N=%d K=%d\n", n, k);
    k+=2;
    int spazio_libero = n;

    for (int i=0; i<k; i++) {
        assert(spazio_libero == 0 ? k-i-1 == 0 : 1);
        int max_spazio = (k-i-1 > 0) ? (spazio_libero-1)/(k-i-1) + ((spazio_libero-1)%(k-i-1)>0? 1: 0) : 0;
        if (i == 0) {
            printf("%d\n", max_spazio-1);
        }
        printf("%d %d\n", i, n-spazio_libero);
        // printf("T #%d ST=%d PSL=%d", i, n-spazio_libero, spazio_libero);
        spazio_libero-=(max_spazio);
        // printf(" MS=%d SL=%d RIM_K=%d\n", max_spazio, spazio_libero, k-i-1);
    }

    return 0;
}