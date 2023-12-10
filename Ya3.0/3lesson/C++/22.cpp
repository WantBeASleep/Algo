#include <iostream>

#include <vector>
#include <deque>

using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);

    int n, k;
    cin >> n >> k;

    int sum = 1;
    vector<int> dp(n);
    dp[0] = sum;
    deque<int> lastk;
    lastk.push_back(dp[0]);

    for (int i = 1; i != n; ++i) {
        dp[i] = sum;
        lastk.push_back(sum);
        sum += dp[i];

        if (lastk.size() > k) {
            sum -= lastk.front();
            lastk.pop_front();
        }
    }

    cout << dp[n - 1] << endl;
    return 0;
}