package main

import (
	"fmt"
)

// https://leetcode.com/problems/design-twitter/description/

// Design a simplified version of Twitter where users can post tweets, follow/unfollow another user, and is able to see the 10 most recent tweets in the user's news feed.
//
// Implement the Twitter class:
//
// Twitter() Initializes your twitter object.
// void postTweet(int userId, int tweetId) Composes a new tweet with ID tweetId by the user userId. Each call to this function will be made with a unique tweetId.
// List<Integer> getNewsFeed(int userId) Retrieves the 10 most recent tweet IDs in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user themself. Tweets must be ordered from most recent to least recent.
// void follow(int followerId, int followeeId) The user with ID followerId started following the user with ID followeeId.
// void unfollow(int followerId, int followeeId) The user with ID followerId started unfollowing the user with ID followeeId.
//
// 1 <= userId, followerId, followeeId <= 500
// 0 <= tweetId <= 10^4
// All the tweets have unique IDs.
// At most 3 * 10^4 calls will be made to postTweet, getNewsFeed, follow, and unfollow.
// A user cannot follow himself.
const FEED_LENGTH = 10

type Twitter struct {
	tweetCount int
	users      map[int]map[int]any // user id to followee id
	tweets     map[int][]Tweet     // user id to user tweets
}

type (
	Tweet struct {
		id    int
		count int
	}
	TweetHeap []Tweet
)

func (t *TweetHeap) Push(x Tweet) {
	*t = append(*t, x)
	t.heapifyUp(len(*t) - 1)
}

func (t *TweetHeap) Pop() Tweet {
	top := (*t)[0]
	(*t)[0] = (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	if len(*t) > 0 {
		t.heapifyDown(len(*t), 0)
	}
	return top
}

func (t *TweetHeap) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2

		if (*t)[parent].count < (*t)[i].count {
			break
		}

		(*t)[parent], (*t)[i] = (*t)[i], (*t)[parent]
		i = parent
	}
}

func (t *TweetHeap) heapifyDown(n, i int) {
	for {
		smallest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && (*t)[left].count < (*t)[smallest].count {
			smallest = left
		}

		if right < n && (*t)[right].count < (*t)[smallest].count {
			smallest = right
		}

		if smallest == i {
			break
		}

		(*t)[smallest], (*t)[i] = (*t)[i], (*t)[smallest]
		i = smallest
	}
}

func Constructor() Twitter {
	return Twitter{users: make(map[int]map[int]any), tweets: make(map[int][]Tweet)}
}

func (this *Twitter) checkUser(userId int) {
	if _, exists := this.users[userId]; !exists {
		this.users[userId] = map[int]any{userId: struct{}{}}
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.checkUser(userId)

	this.tweets[userId] = append(this.tweets[userId], Tweet{tweetId, this.tweetCount})
	this.tweetCount++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	this.checkUser(userId)
	tweetHeap := make(TweetHeap, 0, FEED_LENGTH) // minHeap of FEED_LENGTH of tweets with largest counts

	for followee := range this.users[userId] {
		for _, tweet := range this.tweets[followee] {
			if len(tweetHeap) < FEED_LENGTH {
				tweetHeap.Push(tweet)
			} else {
				if tweet.count > tweetHeap[0].count {
					tweetHeap[0] = tweet
					tweetHeap.heapifyDown(FEED_LENGTH, 0)
				}
			}
		}
	}

	tweetIds := make([]int, len(tweetHeap))
	for i := len(tweetHeap) - 1; i >= 0; i-- {
		tweetHeap[0], tweetHeap[i] = tweetHeap[i], tweetHeap[0]
		tweetIds[i] = tweetHeap[i].id
		tweetHeap.heapifyDown(i, 0)
	}

	return tweetIds
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	this.checkUser(followerId)
	this.users[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if _, exists := this.users[followerId][followeeId]; exists {
		delete(this.users[followerId], followeeId)
	}
}

/**
* Your Twitter object will be instantiated and called as such:
* obj := Constructor();
* obj.PostTweet(userId,tweetId);
* param_2 := obj.GetNewsFeed(userId);
* obj.Follow(followerId,followeeId);

 * obj.Unfollow(followerId,followeeId);
*/

func main() {
	twitter := Constructor()
	twitter.PostTweet(1, 5) // User 1 posts a new tweet (id = 5).
	fmt.Println(
		twitter.GetNewsFeed(1),
		"==",
		5,
	) // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
	twitter.Follow(1, 2)    // User 1 follows user 2.
	twitter.PostTweet(2, 6) // User 2 posts a new tweet (id = 6).
	fmt.Println(
		twitter.GetNewsFeed(1),
		"==",
		"[6,5]",
	) // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
	twitter.Unfollow(1, 2) // User 1 unfollows user 2.
	fmt.Println(
		twitter.GetNewsFeed(1),
		"==",
		"[5]",
	) // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.
}
