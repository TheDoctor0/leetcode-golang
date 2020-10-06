package solutions

import (
    "container/heap"
)

type Twitter struct {
    userList  map[int]User
    timestamp int
}

type User struct {
    userId int
    follow map[int]struct{}
    fans map[int]struct{}
    news *MaxHeap
    posts *[]Tweet
}

type MaxHeap []Tweet

type Tweet struct {
    userId    int
    tweetId   int
    timestamp int
}

func Constructor() Twitter {
    return Twitter{userList: make(map[int]User)}
}

func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.checkUser(userId)
    this.timestamp++

    time := this.timestamp
    tweet := Tweet{
        userId: userId,
        tweetId:   tweetId,
        timestamp: time,
    }

    user := this.userList[userId]
    *user.posts = append(*user.posts, tweet)

    if len(*user.posts) > 10 {
        *user.posts = (*user.posts)[1:]
    }
}

func (this *Twitter) GetNewsFeed(userId int) []int {
    var result []int

    this.checkUser(userId)

    user := this.userList[userId]

    for _, post := range *user.posts {
        heap.Push(user.news, post)
    }

    for followeeId := range user.follow {
        for _, post := range *(this.userList[followeeId].posts) {
            heap.Push(user.news, post)
        }
    }

    for i := 0; i < 10 && user.news.Len() > 0; {
        t := heap.Pop(user.news).(Tweet)
        result = append(result, t.tweetId)
        i++
    }

    *user.news = MaxHeap{}

    return result
}

func (this *Twitter) Follow(followerId int, followeeId int)  {
    if followerId == followeeId {
        return
    }

    this.checkUser(followerId)
    this.checkUser(followeeId)

    follower := this.userList[followerId]
    followee := this.userList[followeeId]
    follower.follow[followeeId] = struct{}{}
    followee.fans[followerId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if followerId == followeeId {
        return
    }

    this.checkUser(followerId)
    this.checkUser(followeeId)

    follower := this.userList[followerId]
    followee := this.userList[followeeId]

    if _, ok := follower.follow[followeeId]; ok {
        delete(follower.follow, followeeId)
    }

    if _, ok := followee.fans[followerId]; ok {
        delete(followee.fans, followerId)
    }
}

func (this *Twitter) checkUser(userId int) {
    if _, ok := this.userList[userId]; !ok {
        this.userList[userId] = User{userId: userId,
            follow: make(map[int]struct{}),
            fans: make(map[int]struct{}),
            news: &MaxHeap{},
            posts: &[]Tweet{}}
    }
}

func (heap MaxHeap) Len() int {
    return len(heap)
}

func (heap MaxHeap) Less(i int, j int) bool {
    return heap[i].timestamp > heap[j].timestamp
}

func (heap MaxHeap) Swap(i int, j int) {
    heap[i], heap[j] = heap[j], heap[i]
}

func (heap *MaxHeap) Push(a interface{}) {
    *heap = append(*heap, a.(Tweet))
}

func (heap *MaxHeap) Pop() interface{} {
    length := len(*heap)
    result := (*heap)[length - 1]
    *heap = (*heap)[:length - 1]

    return result
}
