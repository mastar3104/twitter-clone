"use client"

import styles from "./page.module.css";
import Image from "next/image";
import {UserResource} from "@/infrastracture/resource/user-resource";
import {useState} from "react";
import {postTweet} from "@/infrastracture/xhr/post-tweet";
import {userIcon} from "@/presentation/user-icon";

export default function TweetArea({
    user,
}: {
    user: UserResource
}) {
    const [content, setContent] = useState('');

    const clickTweetButton = async () => {
        await postTweet(content, user.userId)
        const tweetTextArea = document.getElementById("tweet-text-area") as HTMLInputElement;
        tweetTextArea.value = ""
        setContent("")
    }

    return (
        <div className={styles.tweetArea} aria-label="ツイートする">
            <a href={`/home/${user.userId}/user/${user.userId}`}>
                <Image
                    src={userIcon(user.userId)}
                    className={styles.userIcon}
                    alt={user.userName}
                    width="50"
                    height="50"
                ></Image>
            </a>
            <textarea
                id="tweet-text-area"
                placeholder="今何してる？"
                onChange={event => setContent(event.target.value)}
            ></textarea>
            <div className={styles.buttonArea}>
                <button
                    onClick={clickTweetButton}
                >
                    ツイートする
                </button>
            </div>
        </div>
    )
}