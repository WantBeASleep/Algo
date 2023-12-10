#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int t;
    cin >> t;
    for (int i = 0; i < t; i++) {
        string str;
        cin >> str;
        int ans = 0;
        for (int j = 0; j < str.length(); j++) {
            string pali(str, 0, j + 1);
            string cpypali(pali);
            reverse(cpypali.begin(), cpypali.end());
            if (pali == cpypali) ans = j + 1;
        }
        cout << ans << endl;
    }
    return 0;
}