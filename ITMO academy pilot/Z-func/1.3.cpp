#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int q;
    cin >> q;
    for (int i = 0; i < q; i++) {
        string t, p;
        vector<int> idxs;
        int ans = 0;
        cin >> t >> p;
        if (p.length() > t.length()) {
            cout << 0 << endl << endl;
            continue;
        }
        for (int j = 0; j < t.length() - p.length() + 1; j++) {
            bool flag = true;
            for (int k = 0; k < p.length(); k++) {
                int m = j + k;
                if (t[m] != p[k] and p[k] != '?') flag = false;
            }
            if (flag) {
                ans++;
                idxs.push_back(j);
            }
        }
        cout << ans << endl;
        for (int j = 0; j < idxs.size(); j++) cout << idxs[j] << " ";
        cout << endl;
    }

    return 0;
}