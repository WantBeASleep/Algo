#include <bits/stdc++.h>

using namespace std;

vector<int> z(8000 + 5, 0);
int maxzfunc(string s) {
    int l = 0, r = 0;
    int maxz = 0;
    for (int i = 1; i < s.length(); i++) {
        z[i] = 0;
        if (i < r) z[i] = min(z[i - l], r - i);
        while (i + z[i] < s.length() and s[z[i]] == s[i + z[i]]) z[i]++;
        maxz = max(maxz, z[i]);
        if (r < i + z[i]) {
            r = i + z[i];
            l = i;
        }
    }
    return maxz;
}

int sumn (int n) {
    return ((1 + n) * n) / 2;
}

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);

    string s;
    cin >> s;

    if (s.length() == 0) {
        cout << 0 << "\n";
        return 0;
    }

    string t;
    long long ans = 0;

    for (int i = s.length() - 1; i >= 0; i--) {
        t = s[i] + t;
        int x = maxzfunc(t);
        ans += sumn(s.length() - i) - sumn(x);
    }

    cout << ans << "\n";

    return 0;
}