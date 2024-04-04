"use client"

import styles from "@/app/page.module.css";
import React, {useState} from "react";
import {postUser} from "@/infrastracture/xhr/post-user";
import {useRouter} from "next/navigation";

export default function Login() {
    const router = useRouter();

    const [userName, setUserName] = useState('');
    const [userId, setUserId] = useState('');

    const clickTweetButton = async () => {
        const user = await postUser(userName)
        setUserId(user.userId)
        const userIdArea = document.getElementById("user-id-area") as HTMLInputElement;
        userIdArea.value = user.userId
    }


    return (
        <>
            <div className={styles.grid}>
                <h3>サインアップ</h3>
                <p>ユーザ名:
                    <input
                        id="user-name-area"
                        type="text"
                        onChange={event => setUserName(event.target.value)}
                    ></input>
                </p>
                <button
                    onClick={clickTweetButton}
                >送信
                </button>
            </div>
            <div className={styles.grid}>
                <h3>サインイン</h3>
                <p>ユーザID:
                    <input
                        id="user-id-area"
                        type="text"
                        onChange={event => setUserId(event.target.value)}
                    ></input>
                </p>
                <button
                    onClick={() => {
                        router.push(`/home/${userId}`)
                    }}
                >
                    ログイン
                </button>
            </div>
        </>
    )
}