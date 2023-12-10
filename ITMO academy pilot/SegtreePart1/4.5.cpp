/*
    ЭТО ПРОСТО ROFLS, ПО ФАКТУ РЕШЕНИЕ O(NlOGN) Я НЕ ЕБУ ПОЧЕМУ ОНО НА ИЗИ ПРОХОДИТ ЗА 233МС.
    ВЫЙГРЫШ ТОЛЬКО В ОТСЕЧЕНИИ MAX/MIN, НО
    МОЖЕТ БЫТЬ СИТУАЦИИ ГДЕ МЫ ЗАЙДЕМ ПОЧТИ В КАЖДУЮ ВЕРШИНУ НА ОТРЕЗКЕ, ТОГДА ОБХОД БУДЕТ NLOGN, ЧТО КАК БЫ ПИЗДЕЦ
*/

#include <iostream>
#include <vector>
#include <tuple>
#include <climits>

using namespace std;

typedef tuple<int, int, int> pr;

const pr zero = pr(INT_MIN, INT_MAX, 0);

class Segtree {
    public:
        Segtree(int n) {
            init(n);
        }

        void set(int i, int v) {
            set(i, v, 0, 0, size);
        }

        int sharking(int l, int r, int p) {
            return sharking(l, r, p, 0, 0, size);
        }
    private:
        vector<pr> a;
        int size;

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, zero);
        }

        pr repair(pr v1, pr v2) {
            return pr(max(get<0>(v1), get<0>(v2)), min(get<1>(v1), get<1>(v2)), get<2>(v1) + get<2>(v2));
        }

        void set(int i, int v, int x, int lx, int rx) {
            if (rx == lx + 1) {
                a[x] = pr(v, v, 1);
                return;
            }

            int m = (lx + rx) >> 1;
            if (i < m) set(i, v, 2 * x + 1, lx, m);
            else set(i, v, 2 * x + 2, m, rx);

            a[x] = repair(a[2 * x + 1], a[2 * x + 2]);
        }

        void to_zero(int x, int lx, int rx) {
            if (rx == lx + 1) {
                a[x] = zero;
                return;
            }

            int m = (lx + rx) >> 1;
            to_zero(2 * x + 1, lx, m);
            to_zero(2 * x + 2, m, rx);
            a[x] = zero;
        }

        int sharking(int l, int r, int p, int x, int lx, int rx) {
            if (rx <= l or lx >= r) return 0;
            if (get<1>(a[x]) > p or get<2>(a[x]) == 0) return 0;
            
            if (get<0>(a[x]) <= p and lx >= l and rx <= r) {
                int ans = get<2>(a[x]);
                to_zero(x, lx, rx);
                return ans;
            }

            int m = (lx + rx) >> 1;
            int lshark = sharking(l, r, p, 2 * x + 1, lx, m);
            int rshark = sharking(l, r, p, 2 * x + 2, m, rx);
            a[x] = repair(a[2 * x + 1], a[2 * x + 2]);
            return lshark + rshark;
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
            int i, v;
            cin >> i >> v;
            sg.set(i, v);
        } else {
            int l, r, p;
            cin >> l >> r >> p;
            cout << sg.sharking(l, r, p) << endl;
            fflush(stdout);
        }
    }
    return 0;
}