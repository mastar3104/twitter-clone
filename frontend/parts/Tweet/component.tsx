import Image from "next/image";
import styles from "./page.module.css";
import {convertUnixTimeToYMD} from "@/presentation/day-fromat";
import {TweetResource} from "@/infrastracture/resource/tweet-resource";
import {userIcon} from "@/presentation/user-icon";

export default function Tweet({
    myId,
    tweet,
}: {
    myId: string,
    tweet: TweetResource,
}) {
    return (
        <div className={styles.tweetRows}>
            <a href={`/home/${myId}/user/${tweet.userId}`}>
                <Image
                    src={userIcon(tweet.userId)}
                    className={styles.userIcon}
                    alt={tweet.userName}
                    width="50"
                    height="50"
                ></Image>
            </a>
            <div>
                <div>
                    <a href={`/home/${myId}/user/${tweet.userId}`}>
                        <span>{tweet.userName}</span>
                    </a>
                    <span className={styles.tweetTime}>{convertUnixTimeToYMD(tweet.tweetTime)}</span>
                </div>
                <p className={styles.tweetContent}>
                    {tweet.content.split("\n").map((row, index) => {
                        return (
                            <p key={index}>{row}</p>
                        )
                    })}
                </p>
            </div>
        </div>
    )
}