#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    string s;
    cin >> s;
    int n = s.length();

    s = s + s;
    vector<int> z(2 * n, 0);
    int l = 0, r = 0;
    for (int i = 1; i < 2 * n; i++) {
        if (i < r) z[i] = min(z[i - l], r - i);
        while (i + z[i] < 2 * n and s[z[i]] == s[i + z[i]]) z[i]++;
        if (r < i + z[i]) {
            r = i + z[i];
            l = i;
        }
    }

    int ans = 1;
    for (int i = 0; i < n; i++) {
        if (z[i] < n and s[i + z[i]] < s[z[i]]) ans++;
    }


    cout << ans << "\n";
    return 0;
}