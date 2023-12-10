#include <bits/stdc++.h>

using namespace std;

class Segtree {
    public:
        Segtree(vector<int>& data) {
            init(data.size());
            build(data, 0, 0, size);
        }

        void swap(int i) {
            swap(i, 0, 0, size);
        }

        int getidx(int k) {
            return getidx(k, 0, 0, size);
        }



    private:
        vector<int> a;
        int size;

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, 0);
        }

        void build(vector<int>& data, int x, int lx, int rx) {
            if (rx == lx + 1) {
                if (lx < data.size()) a[x] = data[lx];
                return;
            }

            int m = (lx + rx) / 2;
            build(data, 2 * x + 1, lx, m);
            build(data, 2 * x + 2, m, rx);
            a[x] = a[2 * x + 1] + a[2 * x + 2];
        }

        void swap(int i, int x, int lx, int rx) {
            if (rx == lx + 1) {
                a[x] = (a[x] + 1) % 2;
                return;
            }

            int m = (lx + rx) / 2;
            if (i < m) swap(i, 2 * x + 1, lx, m);
            else swap(i, 2 * x + 2, m, rx);
            a[x] = a[2 * x + 1] + a[2 * x + 2];
        }

        int getidx(int k, int x, int lx, int rx) {
            if (rx == lx + 1) return lx;

            int m = (lx + rx) / 2;
            if (k <= a[2 * x + 1]) return getidx(k, 2 * x + 1, lx, m);
            else return getidx(k - a[2 * x + 1], 2 * x + 2, m, rx);
        }

};

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n, m;
    cin >> n >> m;
    vector<int> data(n);
    for (int i = 0; i < n; i++) cin >> data[i];

    Segtree sg(data);

    while(m--) {
        int type;
        cin >> type;

        if (type == 1) {
            int i;
            cin >> i;
            sg.swap(i);
        } else {
            int k;
            cin >> k;
            k++;
            cout << sg.getidx(k) << "\n";
        }
    }

    return 0;
}