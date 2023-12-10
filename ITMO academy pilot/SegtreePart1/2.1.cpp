#include <bits/stdc++.h>

using namespace std;

typedef long long ll;
typedef tuple<ll, ll, ll, ll> t;

// 0 - seg
// 1 - pref
// 2 - suff
// 3 - summ

class Segtree {
    public:
        Segtree(vector<int>& data) {
            init(data.size());
            build(data, 0, 0, size);
        }

        void set(int i, int v) {
            set(i, v, 0, 0, size);
        }

        ll max_segment(int l, int r) {
            return max_segment(l, r, 0, 0, size);
        }

        ll max_segment() {
            return get<0>(a[0]);
        }

    private:
        vector<t> a;
        int size;

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, make_tuple(0, 0, 0, 0));
        }

        t comparation(t s1, t s2) {
            ll seg, preff, suff, summ;
            ll seg1, preff1, suff1, summ1;
            ll seg2, preff2, suff2, summ2;
            tie(seg1, preff1, suff1, summ1) = s1;
            tie(seg2, preff2, suff2, summ2) = s2;

            seg = max({seg1, seg2, suff1 + preff2});
            preff = max({preff1, summ1 + preff2});
            suff = max({suff2, suff1 + summ2});
            summ = summ1 + summ2;

            return make_tuple(seg, preff, suff, summ);
        }

        void build(vector<int>& data, int x, int lx, int rx) {
            if (rx == lx + 1) {
                if (lx < data.size()) a[x] = make_tuple(data[lx], data[lx], data[lx], data[lx]);
                return;
            }
            int m = (lx + rx) / 2;
            build(data, 2 * x + 1, lx, m);
            build(data, 2 * x + 2, m, rx);

            a[x] = comparation(a[2 * x + 1], a[2 * x + 2]);
        }

        void set(int i, int v, int x, int lx, int rx) {
            if (rx == lx + 1) {
                a[x] = make_tuple(v, v, v, v);
                return;
            }
            int m = (lx + rx) / 2;
            if (i < m) set(i, v, 2 * x + 1, lx, m);
            else set(i, v, 2 * x + 2, m, rx);

            a[x] = comparation(a[2 * x + 1], a[2 * x + 2]);
        }

        ll max_segment(int l, int r, int x, int lx, int rx) {
            if (rx <= l or lx >= r) return 0;
            if (lx >= l and rx <= r) return get<0>(a[x]);

            int m = (lx + rx) / 2;
            ll lmax = max_segment(l, r, 2 * x + 1, lx, m);
            ll rmax = max_segment(l, r, 2 * x + 2, m, rx);
            return max(lmax, rmax);
        }
};

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n, m;
    cin >> n >> m;
    vector<int> a(n);
    for (int i = 0; i < n; i++) cin >> a[i];

    Segtree sg(a);
    cout << max(sg.max_segment(), (ll)0) << "\n";
    while (m--) {
        int i, v;
        cin >> i >> v;

        sg.set(i, v);
        cout << max(sg.max_segment(), (ll)0) << "\n";
    }
    return 0;
}