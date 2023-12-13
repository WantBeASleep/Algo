#include <bits/stdc++.h>
 
using namespace std;
 
int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int q;
    cin >> q;
    for (int gl = 0; gl < q; gl++) {
        string s, t;
        cin >> s >> t;
        string news = t + "$" + s + s;
 
        vector<int> z(news.length(), 0);
        int l = 0, r = 0;
        for (int i = 1; i < news.length(); i++) {
            if (i < r) z[i] = min(z[i - l], r - i);
            while (z[i] + i < news.length() and news[z[i]] == news[i + z[i]]) z[i]++;
            if (r < i + z[i]) {
                r = i + z[i];
                l = i;
            }
        }
 
        int ans = -1;
        for (int i = t.length() + 1; i < news.length(); i++) {
            if (z[i] == t.length()) {
                ans = i - (t.length() + 1);
                break;
            }
        }
        cout << ans << endl;
    }
    return 0;
}