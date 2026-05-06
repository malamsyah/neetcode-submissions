func longestConsecutive(nums []int) int {
    set := make(map[int]struct{})

    for _, v := range nums {
        set[v] = struct{}{}
    }

    res := 0
    for _, v := range nums {
        curr := v
        a := 1
        if _, found := set[curr - 1]; found {
            continue
        }

        for {
            if _, found := set[curr + 1]; found {
                curr++
                a++
            } else {
                break
            }
        }
        if a > res {
            res = a
        }
    }

    return res
}
