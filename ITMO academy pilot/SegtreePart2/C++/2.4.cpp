#include <iostream>
#include <vector>
#include <utility>

using namespace std;

typedef long long ll;
typedef pair<ll, ll> p;

class Segtree{
    public:
        Segtree(int n) {
            init(n);
        }

        void modify(int l, int r, int v) {
            modify(l, r, v, 0, 0, size);
        }

        ll calc(int l, int r) {
            return calc(l, r, 0, 0, size);
        }

    private:
        vector<p> a;
        int size;

        ll op_modify(ll a, ll v) {
            return a + v;
        }

        ll op_calc(ll a, ll b) {
            return a + b;
        }

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, p(0, 0));
        }

        void modify(int l, int r, int v, int x, int lx, int rx) {
            if (x >= 2 * size - 1) return;
            if (lx >= r or rx <= l) return;
            if (lx >= l and rx <= r) {
                a[x].first = op_modify(a[x].first, v * (rx - lx));
                a[x].second = op_modify(a[x].second, v);
                return;
            }

            int m = (lx + rx) / 2;
            modify(l, r, v, 2 * x + 1, lx, m);
            modify(l, r, v, 2 * x + 2, m, rx);

            ll tmp = op_calc(a[2 * x + 1].first, a[2 * x + 2].first);
            a[x].first = op_modify(tmp, a[x].second * (rx - lx));

            // a[x].first = op_modify(op_calc(a[2 * x + 1].first, a[2 * x + 2].first), a[x].second * (rx - lx));
        }

        ll calc(int l, int r, int x, int lx, int rx) {
            if (x >= 2 * size - 1) return 0;
            if (lx >= r or rx <= l) return (ll)0;
            if (lx >= l and rx <= r) return a[x].first;

            int m = (lx + rx) / 2;
            ll left = calc(l, r, 2 * x + 1, lx, m);
            ll right = calc(l, r, 2 * x + 2, m, rx);

            ll tmp = op_calc(left, right);
            ll res = op_modify(tmp, a[x].second * (min(r, rx) - max(l, lx)));

            // return op_modify(op_calc(left, right), a[x].second * (min(r, rx) - max(l, lx)));
            return res;
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
            int v;
            cin >> l >> r >> v;
            sg.modify(l, r, v);
        } else {
            int l, r;
            cin >> l >> r;
            cout << sg.calc(l, r) << "\n";
        }
    }
    return 0;
}