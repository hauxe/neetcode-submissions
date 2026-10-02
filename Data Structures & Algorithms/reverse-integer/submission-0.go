func reverse(x int) int {
    // --- Extract sign using arithmetic right shift ---
    // For a 32-bit signed int, shifting right by 31 gives all 1s (-1) if negative,
    // all 0s (0) if non-negative. This is the two's complement sign-extend trick.
    sign := x >> 31

    // --- Get absolute value using XOR + subtract (no branching, no abs()) ---
    // If sign = 0:  (x ^ 0) - 0 = x
    // If sign = -1: (x ^ -1) - (-1) = (~x) + 1 = -x  (classic two's complement negation)
    absX := (x ^ sign) - sign

    var result int64 = 0
    for absX != 0 {
        digit := int64(absX % 10)
        absX /= 10

        // --- Multiply by 10 using shifts instead of '*' ---
        // x * 10 = x * 8 + x * 2 = (x << 3) + (x << 1)
        result = (result << 3) + (result << 1) + digit

        if result > math.MaxInt32 || result < math.MinInt32 {
            return 0
        }
    }

    // --- Reapply sign using the same XOR/subtract trick ---
    final := (result ^ int64(sign)) - int64(sign)
    return int(final)
}