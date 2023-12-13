#include <bits/stdc++.h>

using namespace std;

vector<int> zfunc(vector<int>& s) {
    vector<int> z(s.size(), 0);
    int l = 0, r = 0;
    for (int i = 0; i < z.size(); i++) {
        if (i < r) z[i] = min(z[i - l], r - i);
        while (i + z[i] < z.size() and s[z[i]] == s[i + z[i]]) z[i]++;
        if (r < i + z[i]) {
            r = i + z[i];
            l = i;
        }
    }
    return z;
}

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n, m;
    cin >> n >> m;
    vector<int> inpt(2 * n + 1);
    for (int i = 0; i < n; i++) {
        cin >> inpt[i];
        inpt[2 * n - i] = inpt[i];
    }
    inpt[n] = m + 1;

    vector<int> z(2 * n + 1, 0);
    int l = 0, r = 0;
    for (int i = n + 1; i < 2 * n + 1; i++) {
        if (i < r) z[i] = min(z[i - l], r - i);
        while (i + z[i] < z.size() and inpt[z[i]] == inpt[i + z[i]]) z[i]++;
        if (r < i + z[i]) {
            r = i + z[i];
            l = i;
        }
    }
    
    cout << n << " ";
    for (int i = (2 * n + 1) - 2, j = 1; i > n and j <= n/2 ; i -= 2, j++) {
        if (z[i] >= j) cout << n - j << " ";
    }

    return 0;
}