#include <iostream>
#include <vector>
#include <algorithm>
#include <utility>
#include <climits>

using namespace std;

typedef pair<int, int> p;

class Segtree{
    public:
        Segtree(int n) {
            init(n);
        }

        void modify(int l, int r, int v) {
            modify(l, r, v, 0, 0, size);
        }

        int calc(int l, int r) {
            return calc(l, r, 0, 0, size);
        }

    private:
        vector<p> a;
        int size;

        int op_modify(int a, int b) {
            return a | b;
        }

        int op_calc(int a, int b) {
            return a & b;
        }

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, p(0, 0));
        }

        void modify(int l, int r, int v, int x, int lx, int rx) {
            if (lx >= r or rx <= l) return;
            if (lx >= l and rx <= r) {
                a[x].first = op_modify(a[x].first, v);
                a[x].second = op_modify(a[x].second, v);
                return;
            }

            int m = (lx + rx) / 2;
            modify(l, r, v, 2 * x + 1, lx, m);
            modify(l, r, v, 2 * x + 2, m, rx);

            a[x].first = op_modify(op_calc(a[2 * x + 1].first, a[2 * x + 2].first), a[x].second);
        }

        int calc(int l, int r, int x, int lx, int rx) {
            if (lx >= r or rx <= l) return (1LL << 30) - 1;
            if (lx >= l and rx <= r) {
                return a[x].first;
            }

            int m = (lx + rx) / 2;
            int left = calc(l, r, 2 * x + 1, lx, m);
            int right = calc(l, r, 2 * x + 2, m, rx);

            return op_modify(op_calc(left, right), a[x].second);
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