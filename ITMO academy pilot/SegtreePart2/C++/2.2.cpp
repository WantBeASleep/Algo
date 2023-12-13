#include <iostream>
#include <vector>
#include <algorithm>
#include <utility>
#include <climits>

using namespace std;

typedef long long ll;
typedef pair<ll, ll> p;

#define MODULE (ll)(1e9 + 7)

class Segtree{
    public:
        Segtree(int n) {
            init(n);
            build(0, 0, size);
        }

        void add(int l, int r, ll v) {
            add(l, r, v, 0, 0, size);
        }

        ll sum(int l, int r) {
            return sum(l, r, 0, 0, size);
        }

    private:
        vector<p> a;
        int size;

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, p(1, 1));
        }

        void build(int x, int lx, int rx) {
            if (rx == lx + 1) return;

            int m = (lx + rx) / 2;
            build(2 * x + 1, lx, m);
            build(2 * x + 2, m, rx);

            a[x].first = (a[2 * x + 1].first + a[2 * x + 2].first) % MODULE;
        }

        void add(int l, int r, ll v, int x, int lx, int rx) {
            if (lx >= r or rx <= l) return;
            if (lx >= l and rx <= r) {
                a[x].first = (a[x].first * v) % MODULE;
                a[x].second = (a[x].second * v) % MODULE;
                return;
            }

            int m = (lx + rx) / 2;
            add(l, r, v, 2 * x + 1, lx, m);
            add(l, r, v, 2 * x + 2, m, rx);

            a[x].first = (((a[2 * x + 1].first + a[2 * x + 2].first) % MODULE) * a[x].second) % MODULE;
        }

        ll sum(int l, int r, int x, int lx, int rx) {
            if (lx >= r or rx <= l) return 0;
            if (lx >= l and rx <= r) {
                return a[x].first;
            }

            int m = (lx + rx) / 2;
            ll left = sum(l, r, 2 * x + 1, lx, m);
            ll right = sum(l, r, 2 * x + 2, m, rx);

            return (((left + right) % MODULE) * a[x].second) % MODULE;
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
            cout << sg.sum(l, r) << "\n";
        }
    }
    return 0;
}