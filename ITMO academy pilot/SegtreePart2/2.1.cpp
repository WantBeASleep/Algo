#include <iostream>
#include <vector>
#include <algorithm>
#include <utility>
#include <climits>

using namespace std;

typedef long long ll;
typedef pair<ll, ll> p;

class Segtree{
    public:
        Segtree(int n) {
            init(n);
        }

        void add(int l, int r, ll v) {
            add(l, r, v, 0, 0, size);
        }

        ll min(int l, int r) {
            return min(l, r, 0, 0, size);
        }

    private:
        vector<p> a;
        int size;

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, p(0, 0));
        }

        void add(int l, int r, ll v, int x, int lx, int rx) {
            if (lx >= r or rx <= l) return;
            if (lx >= l and rx <= r) {
                a[x].first += v;
                a[x].second += v;
                return;
            }

            int m = (lx + rx) / 2;
            add(l, r, v, 2 * x + 1, lx, m);
            add(l, r, v, 2 * x + 2, m, rx);

            a[x].first = ::min(a[2 * x + 1].first, a[2 * x + 2].first) + a[x].second;
        }

        ll min(int l, int r, int x, int lx, int rx) {
            if (lx >= r or rx <= l) return LLONG_MAX;
            if (lx >= l and rx <= r) {
                return a[x].first;
            }

            int m = (lx + rx) / 2;
            ll left = min(l, r, 2 * x + 1, lx, m);
            ll right = min(l, r, 2 * x + 2, m, rx);

            return ::min(left, right) + a[x].second;
        }

};

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n, m;
    cin >> n >> m;
    Segtree sg(n);
    while (m--) {
        int type;
        cin >> type;
        if (type == 1) {
            int l, r;
            ll v;
            cin >> l >> r >> v;
            sg.add(l, r, v);
        } else {
            int l, r;
            cin >> l >> r;
            cout << sg.min(l, r) << "\n";
        }
    }
    return 0;
}