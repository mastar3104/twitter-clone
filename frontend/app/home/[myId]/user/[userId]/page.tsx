import styles from "@/app/page.module.css";
import Sidebar from "@/components/Sidebar/component";
import TweetList from "@/components/TweetList/component";
import RecommendUser from "@/components/RecommendUser/component";
import UserArea from "@/components/UserArea/component";
import {FetchTweetList} from "@/infrastracture/fetch/fetch-tweet";
import {FetchUser} from "@/infrastracture/fetch/fetch-user";

export default async function User({
    params,
}: {
    params: {
        myId: string,
        userId: string
    }
}) {
    const userRequest = FetchUser(params.userId)
    const user = await userRequest.getResource()

    return (
        <main className={styles.main}>
            <div className={styles.grid}>
                <Sidebar params={params}></Sidebar>
                <div>
                    <UserArea
                        myId={params.myId}
                        user={user}
                    ></UserArea>
                    <TweetList
                        myId={params.myId}
                        userId={params.userId}
                    ></TweetList>
                </div>
                <RecommendUser myId={params.myId}></RecommendUser>
            </div>
        </main>
    )
}