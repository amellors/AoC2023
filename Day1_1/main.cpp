#include <iostream>
#include <fstream>
#include <string>
#include <cctype>

using namespace std;

int getDigits(string line)
{

    bool done = false;
    int l;
    int r;
    int digits = 0;
    for(l = 0; !isdigit(line[l]); l++);
    digits += 10 * (line[l] - '0');
    for(r = line.size() - 1; !isdigit(line[r]); r--);
    digits += (line[r] - '0');

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