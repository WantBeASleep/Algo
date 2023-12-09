#include <iostream>

#include <map>

using namespace std;

typedef long long ll;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n;
    cin >> n;
    map<ll, ll> m;
    for (int i = 0; i != n; ++i) {
        ll key, val;
        cin >> key >> val;
        m[key] += val;
    }

    map<ll, ll>::iterator it;
    for (it = m.begin(); it != m.end(); ++it) {
        cout << it->first << " " << it->second << endl;
    }
    return 0;
}