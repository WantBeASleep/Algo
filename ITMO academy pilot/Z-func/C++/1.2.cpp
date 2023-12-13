#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int t;
    cin >> t;
    
    for (int i = 0; i < t; i++) {
        int ans = 0;
        string str;
        cin >> str;
        unordered_set<string> prefix;
        vector<string> prefixlist;
        unordered_set<string> postfix;
        vector<string> postfixlist;

        for (int j = 0; j < str.length(); j++) {
            prefix.insert(str.substr(0, j + 1));
            prefixlist.push_back(str.substr(0, j + 1));
        }
        for (int j = 0; j < str.length(); j++) {
            postfix.insert(str.substr(str.length() - 1 - j, j + 1));
            postfixlist.push_back(str.substr(str.length() - 1 - j, j + 1));
        }

        for (int k = 0; k < prefixlist.size(); k++) {
            for (int m = 0; m < str.length(); m++) {
                if (str.substr(m, prefixlist[k].length()) == prefixlist[k]) {
                    if (!postfix.count(str.substr(m, prefixlist[k].length()))) ans++;
                }
            }
        }

        for (int k = 0; k < postfixlist.size(); k++) {
            for (int m = 0; m < str.length(); m++) {
                if (str.substr(m, postfixlist[k].length()) == postfixlist[k]) {
                    if (!prefix.count(str.substr(m, postfixlist[k].length()))) ans++;
                }
            }
        }

        cout << ans << endl;
        
    }
    return 0;
}