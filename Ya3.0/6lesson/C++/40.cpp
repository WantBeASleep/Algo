#include <iostream>

#include <vector>
#include <deque>
#include <unordered_set>
#include <unordered_map>
#include <sstream>
#include <string>
#include <utility>

using namespace std;

void setInsrt(unordered_set<int>& line1, unordered_set<int>& line2, int num2, unordered_map<int, unordered_set<int>>& res);

typedef pair<int, int> p;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n ;
    cin >> n; 

    int m;
    cin >> m;
    
    vector<unordered_set<int>> allStat(m + 1, unordered_set<int>(n));
    string line;
    getline(cin, line);

    for (int i = 1; i != m + 1; ++i) {
        getline(cin, line);
        istringstream ss(line);
        int station;

        while (ss >> station) {
            allStat[i].insert(station);
        }
    }

    int start, end;
    cin >> start >> end;

    vector<unordered_map<int, unordered_set<int>>> cross(m + 1);
    //(m + 1, unordered_map<int, vector<int>>(n));
    for (int i = 1; i != m + 1; ++i) {
        for (int j = 1; j != m + 1; ++j) {
            if (i == j) {
                continue;
            }

            setInsrt(allStat[i], allStat[j], j, cross[i]);
        }
    }

    vector<bool> visitedLines(m + 1, false);

    unordered_set<int> endLines(m + 1);
    for (int i = 1; i != m + 1; ++i) {
        if (allStat[i].count(end)) {
            endLines.insert(i);
        }
    }

    deque<p> q;
    for (int i = 1; i != m + 1; ++i) {
        if (allStat[i].count(start)) {

            if (endLines.count(i)) {
                cout << 0 << endl;
                return 0;
            }

            visitedLines[i] = true;

            for (auto const& [key, val] : cross[i]) {
                q.push_back(p(i, key));
            }
        }
    }

    int len = 1;

    while (!q.empty()) {
        int oldS = q.size();
        for (int i = 0; i != oldS; ++i) {
            p curS = q.front();
            for (auto newL : cross[curS.first][curS.second]) {
                if (visitedLines[newL]) {
                    continue;
                }

                if (endLines.count(newL)) {
                    cout << len << endl;
                    return 0;
                }

                visitedLines[newL] = true;

                for (auto const& [newCrossStations, val] : cross[newL]) {
                    if (newCrossStations == curS.second) {
                        continue;
                    }

                    q.push_back(p(newL, newCrossStations));
                }
            }
            q.pop_front();
        }
        len++;
    }

    cout << -1 << endl;

    return 0;
}

void setInsrt(unordered_set<int>& line1, unordered_set<int>& line2, int num2, unordered_map<int, unordered_set<int>>& res) {
    for (auto station : line1) {
        if (line2.count(station)) {
            res[station].insert(num2);
        }
    }
}

/*
5
5
2 1 2
2 1 3
2 2 3
2 3 4
2 4 5
1 5

>>> 1
Неправильно
каво блять? Ладно, задача гавна, какая задача, такие и тесты
*/