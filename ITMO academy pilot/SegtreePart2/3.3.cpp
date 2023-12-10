#include <iostream>
#include <tuple>
#include <vector>
#include <algorithm>

using namespace std;


typedef long long ll;
typedef tuple<ll, ll, bool> t;

class Segtree{
    public:
        Segtree(int n) {
            init(n);
        }

        void modify(int l, int r, int v) {
            modify(l, r, v, 0, 0, size);
        }

        int calc(ll v, int l) {
            return calc(l, v, 0, 0, size);
        }

    private:
        vector<t> a;
        int size;

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, t(0, 0, false));
        }

        ll op_modify(ll a, ll v) {
            return a + v;
        }

        ll op_calc(ll a, ll b) {
            return max(a, b);
        }

        void propagate(int x, int lx, int rx) {
            if (get<2>(a[x]) == false) return;
            if (rx == lx + 1) return;

            get<1>(a[2 * x + 1]) = op_modify(get<1>(a[2 * x + 1]), get<1>(a[x]));
            get<1>(a[2 * x + 2]) = op_modify(get<1>(a[2 * x + 2]), get<1>(a[x]));

            get<0>(a[2 * x + 1]) = op_modify(get<0>(a[2 * x + 1]), get<1>(a[x]));
            get<0>(a[2 * x + 2]) = op_modify(get<0>(a[2 * x + 2]), get<1>(a[x]));

            get<2>(a[x]) = false;
            get<1>(a[x]) = 0;
            get<2>(a[2 * x + 1]) = true;
            get<2>(a[2 * x + 2]) = true;
        }

        void modify(int l, int r, ll v, int x, int lx, int rx) {
            propagate(x, lx, rx);
            if (rx <= l or lx >= r) return;
            if (lx >= l and rx <= r) {
                get<0>(a[x]) = op_modify(get<0>(a[x]), v);
                get<1>(a[x]) = op_modify(get<1>(a[x]), v);
                get<2>(a[x]) = true;
                return;
            }

            int m = (lx + rx) / 2;
            modify(l, r, v, 2 * x + 1, lx, m);
            modify(l, r, v, 2 * x + 2, m, rx);

            get<0>(a[x]) = op_calc(get<0>(a[2 * x + 1]), get<0>(a[2 * x + 2]));
        }

        int calc(int l, ll v, int x, int lx, int rx) {
            propagate(x, lx, rx);
            if (get<0>(a[x]) < v or (rx <= l)) return -1;
            if (rx == lx + 1) return lx;

            int m = (lx + rx) / 2;  
            int res = calc(l, v, 2 * x + 1, lx, m);
            if (res == -1) return calc(l, v, 2 * x + 2, m, rx);
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
            int l, r, v;
            cin >> l >> r >> v;
            sg.modify(l, r, v);
        } else {
            int x, l;
            cin >> x >> l;
            cout << sg.calc(x, l) << "\n";
        }
    }
}