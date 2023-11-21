#include <iostream>
#include <limits>
using namespace std;

bool isPalindrome(int x)
{
    if (x < 0)
    {
        return false;
    }

    long long reversedNum = 0;
    int num = x;

    while (num != 0)
    {
        int digit = num % 10;

        if ((reversedNum > numeric_limits<int>::max() / 10) ||
            (reversedNum < numeric_limits<int>::min() / 10))
        {
            return false;
        }

        reversedNum = reversedNum * 10 + digit;
        num /= 10;
    }

    return (reversedNum == x);
}

int main()
{
    bool isPal = isPalindrome(10);
    cout << isPal << endl;
    return 0;
}
