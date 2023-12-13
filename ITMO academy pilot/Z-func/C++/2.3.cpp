#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);

    int t;
    cin >> t;
    for (int glbl = 0; glbl < t; glbl++) {
        int n;
        cin >> n;
        vector<int> z(n);
        for (int i = 0; i < n; i++) cin >> z[i];


     
        string s;
        char symb = 97;
        int prefixLen = 0;
        int j = 0;
        for (int i = 0; i < n; i++) {
            if (z[i] == 0 and prefixLen == 0) {
                s += symb;
                if (symb < 122) symb++;
            }
            if (z[i] > prefixLen) {
                prefixLen = z[i];
                j = 0;
            }
            if (prefixLen > 0) {
                s += s[j];
                j++;
                prefixLen--;
            }
        }

        vector<int> znew(s.length(), 0);
        for (int i = 1; i < s.length(); i++) {
            while (i + znew[i] < s.length() and s[i + znew[i]] == s[znew[i]]) znew[i]++;
        }
        bool flag = true;
        if (znew.size() != z.size()) flag = false;
        for (int i = 0; i < z.size(); i++) {
            if (z[i] != znew[i]) flag = false;
        }

        if (flag) cout << s << endl ;
        else cout << "!" << endl;
    }
    
    return 0;
}