#include <bits/stdc++.h>

using namespace std;

vector<int> z(10e6 + 10, 0);

int zfunc(string s, int n) {
    int l = 0, r = 0;
    for (int i = 1; i < s.length(); i++) {
        z[i] = 0;
        if (i < r) z[i] = min(z[i - l], r - i);
        while (i + z[i] < s.length() and s[z[i]] == s[i + z[i]]) z[i]++;
        if (z[i] == n) return -1;
        if (i + z[i] == s.length()) return z[i];
        if (r < i + z[i]) {
            r = i + z[i];
            l = i;
        }
    }
    return 0;
}

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int q;
    cin >> q;

    while (q--) {
        string s, t;
        cin >> s >> t;

        int st = zfunc(s + "$" + t, s.length());
        if (st == -1) {
            cout << t << "\n";
            continue;
        }

        int ts = zfunc(t + "$" + s, t.length());
        if (ts == -1) {
            cout << s << "\n";
            continue;
        }

        if (st > ts) cout << t << s.substr(st) << "\n";
        else cout << s << t.substr(ts) << "\n";
    }
    return 0;
}