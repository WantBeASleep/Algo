#include <bits/stdc++.h>
 
using namespace std;
 
int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int q;
    cin >> q;
    for (int glbl = 0; glbl < q; glbl++) {
        string t;
        cin >> t;
 
        vector<int> z(t.length(), 0);
        int l = 0; int r = 0;
        for (int i = 1; i < t.length(); i++) {
            if (i < r) z[i] = min(z[i - l], r - i);
            while (i + z[i] < t.length() and t[z[i]] == t[i + z[i]]) z[i]++;
            if (r < i + z[i]) {
                r = i + z[i];
                l = i;
            }
        }
 
        for (int i = 0; i < t.length(); i++) {
            if (i + z[i] == t.length()) {
                string ans = t.substr(0, i);
                cout << ans << endl;
                break;
            }
            if (i == t.length() - 1) cout << t << endl;
        }
    }
    return 0;
}