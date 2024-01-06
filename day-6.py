"""
1235. Maximum Profit in Job Scheduling

We have n jobs, where every job is scheduled to be done from startTime[i] to endTime[i], obtaining a profit of profit[i].
You're given the startTime, endTime and profit arrays, return the maximum profit you can take such that there are no two jobs in the subset with overlapping time range.
If you choose a job that ends at time X you will be able to start another job that starts at time X.

Example 1:
Input: startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
Output: 120
Explanation: The subset chosen is the first and fourth job. 
Time range [1-3]+[3-6] , we get profit of 120 = 50 + 70.

Example 2:
Input: startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,60]
Output: 150
Explanation: The subset chosen is the first, fourth and fifth job. 
Profit obtained 150 = 20 + 70 + 60.

Example 3:
Input: startTime = [1,1,1], endTime = [2,3,4], profit = [5,6,4]
Output: 6

Constraints:
1 <= startTime.length == endTime.length == profit.length <= 5 * 104
1 <= startTime[i] < endTime[i] <= 109
1 <= profit[i] <= 104

[(1, 3, 50), (2, 4, 10), (3, 5, 40), (3, 6, 70)]
[(1, 3, 20), (2, 5, 20), (4, 6, 70), (6, 9, 60), (3, 10, 100)]
[(1, 2, 5), (1, 3, 6), (1, 4, 4)]
"""

# import pdb


class Solution:
    def jobScheduling(
        self, startTime: list[int], endTime: list[int], profit: list[int]
    ) -> int:
        # sort by end time
        job_details = sorted(zip(startTime, endTime, profit), key=lambda x: x[1])
        # initiate list to store the possibility of max profit
        max_profit = [0] * len(profit)

        for idx, (start, _, pro) in enumerate(job_details):
            # backtrack to search index that can chaining with job[x]
            # using binary search
            prev_profit_idx = self.binary_search(start, idx, job_details)
            prev_profit = max_profit[prev_profit_idx] if prev_profit_idx >= 0 else 0
            # to compare with latest added profit
            last_profit = max_profit[idx - 1]
            # calculate and insert maximum profit from the current state
            max_profit[idx] = max(last_profit, prev_profit + pro)

        return max_profit[-1]

    def binary_search(self, start: int, idx: int, job_details: list[tuple]):
        left = 0
        right = idx - 1
        while left <= right:
            mid = (left + right) // 2
            if job_details[mid][1] <= start:
                left = mid + 1
            else:
                right = mid - 1

        return right


sol = Solution()
# print(sol.jobScheduling([1, 2, 3, 3], [3, 4, 5, 6], [50, 10, 40, 70]))  # output: 120
print(
    sol.jobScheduling([1, 2, 3, 4, 6], [3, 5, 10, 6, 9], [20, 20, 100, 70, 60])
)  # output: 150
# print(sol.jobScheduling([1, 1, 1], [2, 3, 4], [5, 6, 4]))  # Output: 6
