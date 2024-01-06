/**
 * 300. Longest Increasing Subsequence
Given an integer array nums, return the length of the longest strictly increasing 
subsequence.
 
Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Example 2:

Input: nums = [0,1,0,3,2,3]
Output: 4
Example 3:

Input: nums = [7,7,7,7,7,7,7]
Output: 1
 */

function lengthOfLIS(nums: number[]): number {
    // basic case
    if (!nums) return 0;

    // initiate empty array as long as nums
    let subSequence: number[] = new Array<number>(nums.length);
    // initiate first element to be compared
    subSequence[0] = nums[0];
    // for tracking the step
    let step: number = 1;

    nums.forEach(num => {
        // check iterate num < first subSeq
        if (num < subSequence[0]) {
            subSequence[0] = num;
        // check next num > first/prev subSeq
        // then increment the step
        } else if (num > subSequence[step - 1]) {
            subSequence[step++] = num;  
        // otherwise, check with binary search
        } else {
            let left: number = 0;
            let right: number = step - 1;
            while (left < right) {
                let mid: number = Math.floor((left + right) / 2);

                if (subSequence[mid] < num) left = mid + 1;
                else right = mid
            }

            subSequence[right] = num;
        }
    });
    return step;
};

console.log(lengthOfLIS([10,9,2,5,3,7,101,18, 102])) // 4 [2, 3, 7, 101, 102]

console.log(lengthOfLIS([0,1,0,3,2,3]))