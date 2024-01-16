package main

import (
	"fmt"
	"math/rand"
)

/*
380. Insert Delete GetRandom O(1)

Implement the RandomizedSet class:

RandomizedSet() Initializes the RandomizedSet object.
bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.
You must implement the functions of the class such that each function works in average O(1) time complexity.



Example 1:

Input
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
Output
[null, true, false, true, 2, true, false, 2]

Explanation
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
randomizedSet.insert(2); // 2 was already in the set, so return false.
randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.


["RandomizedSet","insert","insert","remove","insert","insert","insert","remove","remove","insert","remove","insert","insert","insert","insert","insert","getRandom","insert","remove","insert","insert"]
[[],[3],[-2],[2],[1],[-3],[-2],[-2],[3],[-1],[-3],[1],[-2],[-2],[-2],[1],[],[-2],[0],[-3],[1]]


Constraints:

-231 <= val <= 231 - 1
At most 2 * 105 calls will be made to insert, remove, and getRandom.
There will be at least one element in the data structure when getRandom is called.
*/

type RandomizedSet struct {
	elements []int
	mapIndex map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{elements: make([]int, 0), mapIndex: make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, found := this.mapIndex[val]; found {
		return false
	}

	this.elements = append(this.elements, val)
	this.mapIndex[val] = len(this.elements) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	found_idx, found := this.mapIndex[val]
	if !found {
		return false
	}

	// find latest element
	lastElement := this.elements[len(this.elements)-1]
	// replace index on latest index map to current index that removed element
	// and replace element on current index from latest element
	this.elements[found_idx] = lastElement
	this.mapIndex[lastElement] = found_idx

	delete(this.mapIndex, val)
	// rearrange the elements with exlcuding the latest element
	this.elements = this.elements[:len(this.elements)-1]

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.elements[rand.Intn(len(this.elements))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func generateResult(rules []string, vals [][]int) []interface{} {
	obj := Constructor()
	result := make([]interface{}, 0)
	for idx, rule := range rules {
		fmt.Println(idx)
		fmt.Println("RULE")
		fmt.Println(rule)
		fmt.Println("VALS")
		fmt.Println(vals[idx])
		if rule == "RandomizedSet" {
			result = append(result, nil)
			continue
		} else if rule == "insert" {
			x := obj.Insert(vals[idx][0])
			result = append(result, x)
		} else if rule == "remove" {
			x := obj.Remove(vals[idx][0])
			result = append(result, x)
		} else {
			x := obj.GetRandom()
			result = append(result, x)
		}
		fmt.Println("INDEXES")
		fmt.Println(obj.elements)
		fmt.Println("MAPINDEX")
		fmt.Println(obj.mapIndex)
	}
	return result
}

func main() {
	// inputRules := []string{"RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"}
	// inputVals := [][]int{{}, {1}, {2}, {2}, {}, {1}, {2}, {}}

	// fmt.Println(generateResult(inputRules, inputVals))

	// inputRules2 := []string{"RandomizedSet", "insert", "insert", "remove", "insert", "remove", "getRandom"}
	// inputVals2 := [][]int{{}, {0}, {1}, {0}, {2}, {1}, {}}

	// fmt.Println(generateResult(inputRules2, inputVals2))

	inputRules3 := []string{"RandomizedSet", "insert", "insert", "remove", "insert", "insert", "insert", "remove", "remove", "insert", "remove", "insert", "insert", "insert", "insert", "insert", "getRandom", "insert", "remove", "insert", "insert"}
	inputVals3 := [][]int{{}, {3}, {-2}, {2}, {1}, {-3}, {-2}, {-2}, {3}, {-1}, {-3}, {1}, {-2}, {-2}, {-2}, {1}, {}, {-2}, {0}, {-3}, {1}}
	fmt.Println(generateResult(inputRules3, inputVals3))
}
