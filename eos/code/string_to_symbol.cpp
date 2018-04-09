#include <iostream>

using namespace std;

// 精度+代币符号转换成一串 uint64_t 数字.
constexpr uint64_t string_to_symbol(uint8_t precision, const char *str) {
  uint32_t len = 0;
  while (str[len]) {
    ++len;
  }

  uint64_t result = 0;
  for (uint32_t i = 0; i < len; ++i) {
    if (str[i] < 'A' || str[i] > 'Z') {
      /// ERRORS?
    } else {
      result |= (uint64_t(str[i]) << (8 * (1 + i)));
    }
  }

  result |= uint64_t(precision);
  return result;
}

int main() {
  auto symbol = string_to_symbol(4, "EOS");

  cout << symbol << endl;

  return 0; 
}