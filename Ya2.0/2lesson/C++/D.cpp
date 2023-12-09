#include <iostream>
#include <vector>

using namespace std;

bool iscentral(int len, int v);

void lavka(vector<int>& a, int len) {
    int central = -1;
    int left = -1, right = -1;

    for (int i = 0; i != a.size(); ++i) {
        if (iscentral(len, a[i])) central = i;
        if (right == -1 and a[i] >= (len / 2)) {
            right = i;
            left = i-1;
        }
    }

    if (central != -1) {
        cout << a[central] << endl;
    } else {
        cout << a[left] << " " << a[right] << endl;
    }
}

bool iscentral(int len, int v) {
    if (!(len % 2)) return false;
    if (v == len / 2) return true;
    return false;
}

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int len, k;
    cin >> len >> k;
    vector<int> a(k);
    for (int i = 0; i != k; ++i) cin >> a[i];
    lavka(a, len);
    return 0;
}