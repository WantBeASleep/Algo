#include <iostream>
#include <vector>
#include <tuple>

using namespace std;

typedef long long ll;

class Segtree{
    public:
        Segtree(int n) {
            init(n);
        }

        void modify(int l, int r) {
            modify(l, r, 0, 0, size);
        }

        ll calc(int k) {
            return calc(k, 0, 0, size);
        }

    private:

        struct val
        {
            ll count;
        };
        
        typedef tuple<val, bool> t;

        vector<t> a;
        int size;

        val op_modify(val a, int len) {
            return {len - a.count};
        }

        val op_calc(val a, val b) {
            return {a.count + b.count};
        }

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, t({0}, false));
        }

        void propagation(int x, int lx, int rx) {
            if (get<1>(a[x]) == false) return;
            if (rx == lx + 1) return;

            int half = (rx - lx) / 2;

            get<0>(a[2 * x + 1]) = op_modify(get<0>(a[2 * x + 1]), half);
            get<0>(a[2 * x + 2]) = op_modify(get<0>(a[2 * x + 2]), half);
            
            get<1>(a[x]) = false;
            get<1>(a[2 * x + 1]) = !(get<1>(a[2 * x + 1]));
            get<1>(a[2 * x + 2]) = !(get<1>(a[2 * x + 2]));
        }

        void modify(int l, int r, int x, int lx, int rx) {
            propagation(x, lx, rx);
            if (lx >= r or rx <= l) return;
            if (lx >= l and rx <= r) {
                get<0>(a[x]) = op_modify(get<0>(a[x]), (rx - lx));
                get<1>(a[x]) = true;
                return;
            }

            int m = (lx + rx) / 2;
            modify(l, r, 2 * x + 1, lx, m);
            modify(l, r, 2 * x + 2, m, rx);

            get<0>(a[x]) = op_calc(get<0>(a[2 * x + 1]), get<0>(a[2 * x + 2]));
            if (get<1>(a[x]) == true) get<0>(a[x]) = op_modify(get<0>(a[x]), (rx - lx));
        }

        int calc(int k, int x, int lx, int rx) {
            if (rx == lx + 1) return lx;
            propagation(x, lx, rx);
            int m = (lx + rx) / 2;
            if (get<0>(a[2 * x + 1]).count >= k) {
                return calc(k, 2 * x + 1, lx, m);
            } else {
                return calc(k - get<0>(a[2 * x + 1]).count, 2 * x + 2, m, rx);
            }
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
            cin >> l >> r;
            sg.modify(l, r);
        } else {
            int k;
            cin >> k;
            ++k;
            cout << sg.calc(k) << endl;
        }
    }
}