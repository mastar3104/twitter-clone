export interface TweetListResource {
    tweets: TweetResource[]
}

export interface TweetResource {
    userId: string,
    userName: string,
    tweetId: string,
    content: string,
    tweetTime: number,
}