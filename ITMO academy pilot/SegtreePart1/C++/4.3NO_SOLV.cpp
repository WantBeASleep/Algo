#include <iostream>
#include <vector>

using namespace std;

int n;

class Segtree {
    public:
        Segtree() {
            init(n);
        }

        Segtree(int k) {
            init(k);
        }

        void set(int i, int v) {
            set(i, v, 0, 0, size);
        }

        int sum(int l, int r) {
            return sum(l, r, 0, 0, size);
        }

        int sum(int l) {
            return sum(l, size, 0, 0, size);
        }

    private:
        vector<int> a;
        int size;

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, 0);
        }

        void set(int i, int v, int x, int lx, int rx) {
            if (rx == lx + 1) {
                a[x] += v;
                return;
            }

            int m = (lx + rx) >> 1;
            if (i < m) set(i, v, 2 * x + 1, lx, m);
            else set(i, v, 2 * x + 2, m, rx);
            a[x] = a[2 * x + 1] + a[2 * x + 2];
        }

        int sum(int l, int r, int x, int lx, int rx) {
            if (lx >= r or rx <= l) return 0;
            if (lx >= l and rx <= r) return a[x];

            int m = (lx + rx) >> 1;
            int lsum = sum(l, r, 2 * x + 1, lx, m);
            int rsum = sum(l, r, 2 * x + 2, m, rx);
            return lsum + rsum;
        }
};

int main() {
    int m;
    cin >> n >> m;
    
    vector<Segtree> stats(41);
    vector<int> a(n);

    for (int i = 0; i < n; i++) {
        int num;
        cin >> num;

        a[i] = num;

        for (int j = 1; j < num; j++) {
            stats[j].set(i, 1);
        }
    }

    while (m--) {
        int type;
        cin >> type;

        if (type == 1) {
            int l, r;
            cin >> l >> r;
            l--;

            int ans = 0;
            for (int i = l; i < r; i++) {
                ans += stats[a[i]].sum(l, i);
            }
            cout << ans << "\n";
        } else {
            int x, y;
            cin >> x >> y;
            x--;
            
            for (int j = 1; j < a[x]; j++) {
                stats[j].set(x, -1);
            }

            for (int j = 1; j < y; j++) {
                stats[j].set(x, 1);
            }

            a[x] = y;
        }
    }
}