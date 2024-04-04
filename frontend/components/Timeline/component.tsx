"use client"

import styles from "./page.module.css";
import {FetchTimelineList} from "@/infrastracture/fetch/fetch-tweet";
import Tweet from "@/parts/Tweet/component";
import {useEffect, useState} from "react";
import {TweetListResource} from "@/infrastracture/resource/tweet-resource";

export default function TimelineList({
    myId,
}: {
    myId: string,
}) {
    const [timeline, setTimeline] = useState<TweetListResource>()

    useEffect(() => {
        // タイムラインは定期更新を行う
        const fetchData = async () => {
            const tweetListRequest = FetchTimelineList(myId)
            const resource = await tweetListRequest.getResource()
            setTimeline(resource)
        }
        fetchData().then()
        const interval = setInterval(() => {
            fetchData().then();
        }, 5000);
        return () => clearInterval(interval);
    }, [])

    return (
        <section className={styles.tweetList} aria-label="タイムライン: ホームタイムライン">
            {timeline && timeline.tweets.map((tweet) => {
                return (
                    <Tweet key={tweet.tweetId} myId={myId} tweet={tweet}></Tweet>
                )
            })}
        </section>
    )
}