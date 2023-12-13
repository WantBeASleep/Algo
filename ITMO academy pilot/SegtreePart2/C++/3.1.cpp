#include <iostream>
#include <vector>
#include <algorithm>
#include <tuple>
#include <climits>

using namespace std;

typedef long long ll;

class Segtree{
    public:
        Segtree(int n) {
            init(n);
        }

        void modify(int l, int r, ll v) {
            modify(l, r, v, 0, 0, size);
        }

        // ll calc(int l, int r) {
        //     return calc(l, r, 0, 0, size).seg;
        // }

        ll calc() {
            return fastcalc();
        }

    private:

        struct val
        {
            ll seg;
            ll pref;
            ll suff;
            ll sum;
        };
        
        typedef tuple<val, ll, bool> t;

        vector<t> a;
        int size;

        val op_modify(ll v) {
            val a;
            if (v >= 0) {
                a.seg = v;
                a.pref = v;
                a.suff = v;
                a.sum = v;
            } else {
                a.seg = 0;
                a.pref = 0;
                a.suff = 0;
                a.sum = v;
            }
            return a;
        }

        val op_calc(val a, val b) {
            val res;
            res.seg = max(max(a.seg, b.seg), (a.suff + b.pref));
            res.pref = max(a.pref, (a.sum + b.pref));
            res.suff = max(b.suff, (a.suff + b.sum));
            res.sum = a.sum + b.sum;
            return res;
        }

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, t({0, 0, 0, 0}, 0, false));
        }

        void propagation(int x, int lx, int rx) {
            if (get<2>(a[x]) == false) return;
            if (rx == lx + 1) return;

            int half = (rx - lx) / 2;

            get<1>(a[2 * x + 1]) = get<1>(a[x]);
            get<1>(a[2 * x + 2]) = get<1>(a[x]);

            get<0>(a[2 * x + 1]) = op_modify(get<1>(a[2 * x + 1]) * half);
            get<0>(a[2 * x + 2]) = op_modify(get<1>(a[2 * x + 2]) * half);

            get<2>(a[x]) = false;
            get<2>(a[2 * x + 1]) = true;
            get<2>(a[2 * x + 2]) = true;
        }

        void modify(int l, int r, ll v, int x, int lx, int rx) {
            propagation(x, lx, rx);
            if (lx >= r or rx <= l) return;
            if (lx >= l and rx <= r) {
                get<0>(a[x]) = op_modify(v * (rx - lx));
                get<1>(a[x]) = v;
                get<2>(a[x]) = true;
                return;
            }

            int m = (lx + rx) / 2;
            modify(l, r, v, 2 * x + 1, lx, m);
            modify(l, r, v, 2 * x + 2, m, rx);

            get<0>(a[x]) = op_calc(get<0>(a[2 * x + 1]), get<0>(a[2 * x + 2]));
            if (get<2>(a[x]) == true) get<0>(a[x]) = op_modify(get<1>(a[x]) * (rx - lx));
        }

        // val calc(int l, int r, int x, int lx, int rx) {
        //     if (lx >= r or rx <= l) return {LLONG_MIN, LLONG_MIN, LLONG_MIN, LLONG_MIN};
        //     if (lx >= l and rx <= r) {
        //         return get<0>(a[x]);
        //     }

        //     int m = (lx + rx) / 2;
        //     val left = calc(l, r, 2 * x + 1, lx, m);
        //     val right = calc(l, r, 2 * x + 2, m, rx);

        //     val subtree = op_calc(left, right);
        //     if (get<2>(a[x]) == true) subtree = op_modify(get<1>(a[x]) * (min(r, rx) - max(l, lx)));

        //     return subtree;
        // }

        ll fastcalc() {
            return get<0>(a[0]).seg;
        }

};

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n, m;
    cin >> n >> m;
    Segtree sg(n);
    vector<ll> ans;
    while (m--) {
        int l, r;
        ll v;
        cin >> l >> r >> v;
        sg.modify(l, r, v);
        ans.push_back(sg.calc());
    }

    for (auto x : ans) {
        cout << x << endl;
    }
}