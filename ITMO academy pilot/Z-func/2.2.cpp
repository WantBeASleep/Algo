#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int q;
    cin >> q;
    for (int gl = 0; gl < q; gl++) {
        int k, j;
        cin >> k >> j;

        if (j % 2 or j == 0) {
            cout << 0 << endl;
            continue;
        }

        int lx = 0;
        int rx = pow(2, k) - 1;
        while (lx != j) {
            int m = (lx + rx) / 2;
            if (j > m) lx = m + 1;
            else rx = m;
        }
        
        cout << (rx - lx) << endl;
    }
    return 0;
}