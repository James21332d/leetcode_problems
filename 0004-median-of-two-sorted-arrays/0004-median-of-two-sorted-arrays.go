func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) > len(nums2) {
        return findMedianSortedArrays(nums2, nums1)
    }

    m, n := len(nums1), len(nums2)
    lo, hi := 0, m
    half := (m + n + 1) / 2

    for lo <= hi {
        i := (lo + hi) / 2
        j := half - i

        left1 := -1 << 60
        if i > 0 {
            left1 = nums1[i-1]
        }
        right1 := 1<<60 - 1
        if i < m {
            right1 = nums1[i]
        }

        left2 := -1 << 60
        if j > 0 {
            left2 = nums2[j-1]
        }
        right2 := 1<<60 - 1
        if j < n {
            right2 = nums2[j]
        }

        if left1 <= right2 && left2 <= right1 {
            if (m+n)%2 == 1 {
                return float64(max(left1, left2))
            }
            return float64(max(left1, left2)+min(right1, right2)) / 2.0
        }

        if left1 > right2 {
            hi = i - 1
        } else {
            lo = i + 1
        }
    }

    return 0.0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
