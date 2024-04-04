"use client"

import styles from "./page.module.css";
import Image from "next/image";
import {UserResource} from "@/infrastracture/resource/user-resource";
import {postFollow} from "@/infrastracture/xhr/post-follow";
import {userIcon} from "@/presentation/user-icon";

export default async function UserArea({
    myId,
    user,
}: {
    myId: string,
    user: UserResource
}) {
    const clickFollowButton = async () => {
        await postFollow(myId, user.userId)
    }

    return (
        <div className={styles.userArea} aria-label={user.userName}>
            <Image
                src={userIcon(user.userId)}
                className={styles.userIcon}
                alt={user.userName}
                width="50"
                height="50"
            ></Image>
            <p>
                {user.userName}
            </p>
            { myId !== user.userId &&
              <button
                onClick={clickFollowButton}
              >
                フォローする
              </button>
            }
        </div>
    )
}
