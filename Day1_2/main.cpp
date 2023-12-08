#include <iostream>
#include <fstream>
#include <string>
#include <cctype>
#include <unordered_map>

using namespace std;

unordered_map<string, int> digit_map {
    {"one", 1},
    {"two", 2},
    {"three", 3},
    {"four", 4},
    {"five", 5},
    {"six", 6},
    {"seven", 7},
    {"eight", 8},
    {"nine", 9},
    {"1", 1},
    {"2", 2},
    {"3", 3},
    {"4", 4},
    {"5", 5},
    {"6", 6},
    {"7", 7},
    {"8", 8},
    {"9", 9}
};

vector<string> digit_strings {
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
    "1",
    "2",
    "3",
    "4",
    "5",
    "6",
    "7",
    "8",
    "9"
};

int getFirstNum(string line) {
    for (int i = 0; i < line.size(); i++)
    {
        for (auto match : digit_strings) {
            int find_pos = line.rfind(match,i);
            if (find_pos == i) {
                return digit_map[match];
            }
        }
    }

    return 0;
}

int getLastNum(string line) {
    for (int i = line.size(); i >=0 ; i--)
    {
        for (auto match : digit_strings) {
            int find_pos = line.rfind(match,i);
            if (find_pos == i) {
                return digit_map[match];
            }
        }
    }

    return 0;
}

int getDigits(string line)
{

    bool done = false;
    int digits = 0;
    digits += 10 * getFirstNum(line);
    digits += getLastNum(line);

    return digits;
}

int main()
{
    ifstream inputFile("input_file");

    string line;
    int sum = 0;
    while ( getline(inputFile,line) )
    {
        sum += getDigits(line);;
    }

    cout << "Sum is: " << sum;

    return 0;
}