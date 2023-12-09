#include <iostream>
#include <algorithm>

using namespace std;

typedef long long ll;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    ll a, k, b, m, x;
    cin >> a >> k >> b >> m >> x;

    ll l = 0;
    ll r = 2 * x / max(a, b) + 1;

    while (r != l + 1) {
        ll mid = (l + r) / 2;
        if ((a * (mid - (mid / k)) + b * (mid - (mid / m))) >= x) {
            r = mid;
        } else {
            l = mid;
        }
    }

    cout << r << endl;

    return 0;
}
