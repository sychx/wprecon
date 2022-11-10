package text

func FormatConfidence(value0 int8, value1 int8) int8 {
    if value0 >= 100 { return 100 }

    return value0 + value1
}
