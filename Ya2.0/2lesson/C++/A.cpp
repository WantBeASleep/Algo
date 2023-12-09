#include <iostream>

using namespace std;

typedef long long ll;

ll getmax() {
    ll maxelem, curelem;
    int count = 0;
    cin >> maxelem;
    curelem = maxelem;

    while (curelem != 0) {
        if (curelem == maxelem) {
            ++count;
        } else if (curelem > maxelem) {
            maxelem = curelem;
            count = 1;
        }
        cin >> curelem;
    }

    return count;
}

int main()
{
    // ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    cout << getmax() << endl;
    return 0;
}