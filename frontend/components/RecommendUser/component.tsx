"use client"

import mainStyles from "@/app/page.module.css";
import Image from "next/image";
import styles from "./page.module.css";
import {postFollow} from "@/infrastracture/xhr/post-follow";
import {FetchUsers} from "@/infrastracture/fetch/fetch-user";
import {userIcon} from "@/presentation/user-icon";

export default async function RecommendUser({
    myId,
}: {
    myId: string,
}) {
    const usersRequest = FetchUsers()
    const users = (await usersRequest.getListResource()).users

    const clickFollowButton = (userId: string) => {
        return async () => {
            await postFollow(myId, userId)
        }
    }

    return (
        <aside className={mainStyles.card}>
            <h2 className={styles.title}>おすすめユーザ</h2>
            {users.map((user) => {
                if (user.userId === myId) {
                    return
                }
                return (
                    <div key={user.userId} className={styles.userArea}>
                        <a href={`/home/${myId}/user/${user.userId}`}>
                            <Image
                                src={userIcon(user.userId)}
                                className={styles.userIcon}
                                alt={user.userName}
                                width="50"
                                height="50"
                            ></Image>
                        </a>
                        <a href={`/home/${myId}/user/${user.userId}`}>
                            <p>
                                {user.userName}
                            </p>
                        </a>
                        <button
                            onClick={clickFollowButton(user.userId)}
                        >
                            フォローする
                        </button>
                    </div>
                )
            })}
        </aside>
    )
}