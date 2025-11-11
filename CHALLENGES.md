# Challenges Checklist

A running checklist of problems to tackle. We will add and refine this list as we go.

## How to use

- Add each challenge as a new checklist item in the section below.
- Prefer kebab-case slugs (e.g., `two-sum`).
- Keep short notes/tags inline. Longer notes can be placed in a separate line below the item.

Template to copy:

```markdown
- [ ] <title> (slug: <slug>) — <tags/notes>
  - Notes: <optional details or links>
```

Status legend:

- [ ] Pending / in progress
- [x] Completed

## Checklist

- [x] Two Sum (slug: two-sum)
  - Description: Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
  - Problem: `func TwoSum(nums []int, target int) []int`
  - Rationale: The #1 icebreaker. Tests your knowledge of maps (hash tables) for O(n) efficiency vs. the O(n^2) brute-force.
  - Key Concepts: Maps (`map[int]int`), slices (`[]int`), loops (`for range`).
  - Link: [LeetCode – Two Sum](https://leetcode.com/problems/two-sum/)
  - Dir: [challenges/two-sum/](challenges/two-sum/)

- [x] Valid Palindrome (slug: valid-palindrome)
  - Description: Return true if a string is a palindrome after lowercasing and removing non-alphanumeric characters.
  - Problem: `func isPalindrome(s string) bool`
  - Rationale: Tests string normalization and two-pointer technique.
  - Key Concepts: String iteration, `unicode.IsLetter`, `unicode.IsNumber`, `unicode.ToLower`, two-pointers.
  - Dir: [challenges/valid-palindrome/](challenges/valid-palindrome/)

- [ ] Valid Anagram (slug: valid-anagram)
  - Description: Return true if t is an anagram of s.
  - Problem: `func isAnagram(s string, t string) bool`
  - Rationale: Map-based character frequency counting.
  - Key Concepts: Maps (`map[rune]int`), length check, loops.
  - Dir: [challenges/valid-anagram/](challenges/valid-anagram/)

- [ ] Contains Duplicate (slug: contains-duplicate)
  - Description: Return true if any value appears at least twice.
  - Problem: `func containsDuplicate(nums []int) bool`
  - Rationale: Use a map as a set to detect duplicates.
  - Key Concepts: Maps as sets (`map[int]struct{}` or `map[int]bool`).
  - Dir: [challenges/contains-duplicate/](challenges/contains-duplicate/)

- [ ] Reverse String (slug: reverse-string)
  - Description: Reverse a string in place with O(1) extra memory given `[]byte`.
  - Problem: `func reverseString(s []byte)`
  - Rationale: In-place algorithms and two-pointer swaps.
  - Key Concepts: In-place operations, two-pointers, slice manipulation.
  - Dir: [challenges/reverse-string/](challenges/reverse-string/)

- [ ] Group Anagrams (slug: group-anagrams)
  - Description: Group strings that are anagrams of each other.
  - Problem: `func groupAnagrams(strs []string) [][]string`
  - Rationale: Build a canonical key (sorted string or char counts) for a map.
  - Key Concepts: Maps (`map[string][]string`), sorting (`sort.Slice`), byte conversion.
  - Dir: [challenges/group-anagrams/](challenges/group-anagrams/)

- [ ] Top K Frequent Elements (slug: top-k-frequent-elements)
  - Description: Return the k most frequent elements in an array.
  - Problem: `func topKFrequent(nums []int, k int) []int`
  - Rationale: Frequency counting and selection by sorting or heap/bucket.
  - Key Concepts: Frequency map, sorting (`sort.SliceStable`), bucket sort or heap.
  - Dir: [challenges/top-k-frequent-elements/](challenges/top-k-frequent-elements/)

- [ ] Basic REST API JSON Marshalling (slug: rest-api-json-marshalling)
  - Description: Define structs to marshal into a given JSON shape with nested fields.
  - Problem: `type User struct { ... }`, `type Geo struct { ... }`
  - Rationale: Models closely match API payloads using struct tags.
  - Key Concepts: Structs, JSON tags (e.g., ``json:"user_name"``), nested structs.
  - Dir: [challenges/rest-api-json-marshalling/](challenges/rest-api-json-marshalling/)

- [ ] Basic REST API JSON Unmarshalling (slug: rest-api-json-unmarshalling)
  - Description: Unmarshal an error JSON and return the message.
  - Problem: `func getErrorMessage(jsonInput string) (string, error)`
  - Rationale: Practice `json.Unmarshal` and error handling.
  - Key Concepts: `encoding/json`, structs, error handling.
  - Dir: [challenges/rest-api-json-unmarshalling/](challenges/rest-api-json-unmarshalling/)

- [ ] Valid Parentheses (slug: valid-parentheses)
  - Description: Determine if brackets in a string are valid and properly nested.
  - Problem: `func isValid(s string) bool`
  - Rationale: Classic stack problem for parsers.
  - Key Concepts: Slice as stack, maps for matching pairs.
  - Dir: [challenges/valid-parentheses/](challenges/valid-parentheses/)

- [ ] K Closest Points to Origin (slug: k-closest-points-to-origin)
  - Description: Return the k points closest to (0,0).
  - Problem: `func kClosest(points [][]int, k int) [][]int`
  - Rationale: Sort or use a heap by distance; relevant to geo services.
  - Key Concepts: Sorting (`sort.Slice`), distance metric, optional max-heap.
  - Dir: [challenges/k-closest-points-to-origin/](challenges/k-closest-points-to-origin/)

- [ ] LRU Cache (slug: lru-cache)
  - Description: Design an LRU cache supporting `Get` and `Put` in O(1).
  - Problem: `type LRUCache struct { ... }`, `func Constructor(capacity int) LRUCache`, `func (c *LRUCache) Get(key int) int`, `func (c *LRUCache) Put(key int, value int)`
  - Rationale: Combine map + doubly-linked list to meet performance needs.
  - Key Concepts: Structs, maps, doubly-linked list.
  - Dir: [challenges/lru-cache/](challenges/lru-cache/)

- [ ] Subarray Sum Equals K (slug: subarray-sum-equals-k)
  - Description: Count subarrays whose sum equals k.
  - Problem: `func subarraySum(nums []int, k int) int`
  - Rationale: Uses prefix sums with a map to count occurrences efficiently.
  - Key Concepts: Prefix sums, maps.
  - Dir: [challenges/subarray-sum-equals-k/](challenges/subarray-sum-equals-k/)

- [ ] Merge Intervals (slug: merge-intervals)
  - Description: Merge overlapping intervals.
  - Problem: `func merge(intervals [][]int) [][]int`
  - Rationale: Sort by start, then greedily merge.
  - Key Concepts: Sorting (`sort.Slice`), greedy, slice manipulation.
  - Dir: [challenges/merge-intervals/](challenges/merge-intervals/)

- [ ] Longest Substring Without Repeating Characters (slug: longest-substring-without-repeating-characters)
  - Description: Length of the longest substring without repeating characters.
  - Problem: `func lengthOfLongestSubstring(s string) int`
  - Rationale: Sliding window over string with last-seen indices.
  - Key Concepts: Sliding window, maps as sets.
  - Dir: [challenges/longest-substring-without-repeating-characters/](challenges/longest-substring-without-repeating-characters/)

- [ ] Web Crawler / Worker Pool (slug: web-crawler-worker-pool)
  - Description: Crawl a list of URLs concurrently using a fixed-size worker pool.
  - Problem: `func crawlUrls(urls []string, workerCount int) map[string]string`
  - Rationale: Tests concurrency, bounded parallelism, and I/O integrations.
  - Key Concepts: Goroutines, channels, `sync.WaitGroup`, `http.Get`.
  - Dir: [challenges/web-crawler-worker-pool/](challenges/web-crawler-worker-pool/)

- [ ] Find First Duplicate in a Stream (slug: first-duplicate-stream)
  - Description: From a stream (channel) of ints, return the first duplicate and stop.
  - Problem: `func findFirstDuplicate(stream <-chan int) int`
  - Rationale: Manage state in a concurrent stream processor.
  - Key Concepts: Read-only channels (`<-chan`), maps as sets, `select`.
  - Dir: [challenges/first-duplicate-stream/](challenges/first-duplicate-stream/)

- [ ] Bounded Parallelism (Worker Pool) (slug: bounded-parallelism-worker-pool)
  - Description: Run a list of tasks with a maximum concurrency limit.
  - Problem: `func runTasks(tasks []func(), maxConcurrency int)`
  - Rationale: Protects resources and respects rate limits.
  - Key Concepts: Buffered channels (semaphore), `sync.WaitGroup`, goroutines.
  - Dir: [challenges/bounded-parallelism-worker-pool/](challenges/bounded-parallelism-worker-pool/)

- [ ] Rate Limiter (slug: rate-limiter)
  - Description: Allow up to N requests per second.
  - Problem: `type RateLimiter struct { ... }`, `func (rl *RateLimiter) Allow() bool`
  - Rationale: Fundamental for high-performance APIs.
  - Key Concepts: `time.Ticker`, token bucket, channels or `sync.Mutex`.
  - Dir: [challenges/rate-limiter/](challenges/rate-limiter/)

- [ ] K-Nearest Neighbors (Geo-Search) (slug: k-nearest-neighbors-geo-search)
  - Description: Return IDs of the k nearest stores to a user location.
  - Problem: `type Store struct { ... }`, `func findNearestStores(userLat, userLon float64, stores []Store, k int) []string`
  - Rationale: Geo-service style problem measuring distances and sorting.
  - Key Concepts: Structs, sort by distance (Haversine or Euclidean).
  - Dir: [challenges/k-nearest-neighbors-geo-search/](challenges/k-nearest-neighbors-geo-search/)

- [ ] Message Processing Pipeline (Channels) (slug: message-processing-pipeline)
  - Description: 3-stage pipeline: produce 1..100, square, then consume/print.
  - Problem: `func main()` to wire the stages with channels.
  - Rationale: Demonstrates channel-based pipelines and synchronization.
  - Key Concepts: Buffered/unbuffered channels, goroutines, `sync.WaitGroup`, `close`.
  - Dir: [challenges/message-processing-pipeline/](challenges/message-processing-pipeline/)

- [ ] API Aggregator Service (slug: api-aggregator-service)
  - Description: Fetch user and orders in parallel and aggregate; fail on any error.
  - Problem: `func getAggregatedData(userID string) (AggregatedData, error)`
  - Rationale: Concurrent I/O, error propagation, data aggregation.
  - Key Concepts: Goroutines, `sync.WaitGroup`, channels for results/errors, `context.Context`.
  - Dir: [challenges/api-aggregator-service/](challenges/api-aggregator-service/)

- [ ] SQL-like Group By (slug: sql-like-group-by)
  - Description: Group sales by region and sum amounts.
  - Problem: `func groupSalesByRegion(sales []Sale) map[string]int`
  - Rationale: Implement a simple GROUP BY in Go.
  - Key Concepts: Maps, structs, loops.
  - Dir: [challenges/sql-like-group-by/](challenges/sql-like-group-by/)

- [ ] In-Memory Search Index (slug: in-memory-search-index)
  - Description: Build an inverted index with `Add` and `Search`.
  - Problem: `type Index struct { ... }`, `func (idx *Index) Add(...)`, `func (idx *Index) Search(...) []string`
  - Rationale: Architecture + data structures for fast search.
  - Key Concepts: Maps (`map[string][]string`), `strings.Split`, optional `sync.RWMutex`.
  - Dir: [challenges/in-memory-search-index/](challenges/in-memory-search-index/)

- [ ] Find All Duplicates in an Array (slug: find-all-duplicates-in-array)
  - Description: Find elements appearing twice in O(n) time and O(1) extra space.
  - Problem: `func findDuplicates(nums []int) []int`
  - Rationale: Use the array itself as a hash by negating indices.
  - Key Concepts: In-place modification, indexing tricks.
  - Dir: [challenges/find-all-duplicates-in-array/](challenges/find-all-duplicates-in-array/)
