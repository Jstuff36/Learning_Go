package main

func twoSum(nums []int, target int) []int {
    potentialSolutions := make(map[int]int)
    for i, num := range nums {
        testVal := target - num
        if testValIndex, ok := potentialSolutions[testVal]; ok {
            return []int{testValIndex, i}
        } else {
            potentialSolutions[num] = i
        }
    }
    return nil
}