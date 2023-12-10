#include <bits/stdc++.h>

using namespace std;

int func(int n) {
    int ans = 0;
    for (int i = 1; i <= n; i++) ans += i;
    return ans;
}

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int q;
    cin >> q;
    for (int i = 0; i < q; i++) {
        string t, p;
        cin >> t >> p;
        vector<int> pos;
        int ans = 0;
        int diff = t.length() - p.length();

        for (int j = 0; j < diff + 1; j++) {
            bool flag = true;
            for (int k = 0; k < p.length(); k++) {
                int m = k + j;
                if (t[m] != p[k]) flag = false;
            }
            if (flag) pos.push_back(j);
        }

        if (pos.size() > 0) {
            ans += func(pos[0] + p.length() - 1);
            for (int j = 1; j < pos.size(); j++) {
                int n = pos[j] + p.length() - 2 - (pos[j - 1] + 1) + 1;
                ans += func(n);
            }
            ans += func(t.size() - (pos[pos.size() - 1] + 1));
        } else if (pos.size() == 0) ans = func(t.size());
        if (p.length() >= 3) {
            int diffr = func(p.length() - 2);
            ans -= pos.size() * diffr;
        }
        cout << ans << endl;
    }
    return 0;
}