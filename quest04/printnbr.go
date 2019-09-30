package piscine

import "github.com/01-edu/z01"

func intToRune(num int) rune {
  if num < 0 {
    num *= -1
  }
  if num == 0 {
    return '0'
  } else if num == 1 {
    return '1'
  } else if num == 2 {
    return '2'
  } else if num == 3 {
    return '3'
  } else if num == 4 {
    return '4'
  } else if num == 5 {
    return '5'
  } else if num == 6 {
    return '6'
  } else if num == 7 {
    return '7'
  } else if num == 8 {
    return '8'
  } else {
    return '9'
  }
}

func PrintNbr(n int) {
  var neg int = 1

  if n < 0 {
    neg = -1
    z01.PrintRune('-')
  }
  if n >= 10 || n <= -10 {
    PrintNbr(n / 10 * neg)
  }
  z01.PrintRune(intToRune(n % 10))
}
