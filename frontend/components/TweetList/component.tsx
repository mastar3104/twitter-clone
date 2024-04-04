"use client"

import styles from "./page.module.css";
import {FetchTweetList} from "@/infrastracture/fetch/fetch-tweet";
import Tweet from "@/parts/Tweet/component";

export default async function TweetList({
    myId,
    userId,
}: {
    myId: string,
    userId: string,
}) {
    const tweetListRequest = FetchTweetList(userId)
    const tweetList = (await tweetListRequest.getResource())

    return (
        <section className={styles.tweetList} aria-label="タイムライン: ホームタイムライン">
            {tweetList.tweets.map((tweet) => {
                return (
                    <Tweet key={tweet.tweetId} myId={myId} tweet={tweet}></Tweet>
                )
            })}
        </section>
    )
}