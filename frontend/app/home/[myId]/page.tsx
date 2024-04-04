import styles from "@/app/page.module.css";
import Sidebar from "@/components/Sidebar/component";
import TweetArea from "@/components/TweetArea/component";
import RecommendUser from "@/components/RecommendUser/component";
import {FetchUser} from "@/infrastracture/fetch/fetch-user";
import TimelineList from "@/components/Timeline/component";

export default async function Home({
    params,
}: {
    params: { myId: string }
}) {
    const userRequest = FetchUser(params.myId)
    const user = await userRequest.getResource()

    return (
        <main className={styles.main}>
            <div className={styles.grid}>
                <Sidebar params={params}></Sidebar>
                <div>
                    <TweetArea
                        user={user}
                    ></TweetArea>
                    <TimelineList
                        myId={params.myId}
                    ></TimelineList>
                </div>
                <RecommendUser myId={params.myId}></RecommendUser>
            </div>
        </main>
    )
}