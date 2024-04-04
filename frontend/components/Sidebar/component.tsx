import Link from "next/link";
import styles from "./page.module.css";
import mainStyles from "@/app/page.module.css";

export default function Sidebar({
    params,
}: {
    params: {
        myId: string
    }
}) {
    return (
        <nav aria-label="メインメニュー" role="navigation">
            <div className={styles.sidebar}>
                <Link className={mainStyles.card} href={`/home/${params.myId}`}>ホーム</Link>
                <Link className={mainStyles.card} href="/#">話題を検索</Link>
                <Link className={mainStyles.card} href="/#">通知</Link>
                <Link className={mainStyles.card} href="/#">メッセージ</Link>
                <Link className={mainStyles.card} href="/#">リスト</Link>
                <Link className={mainStyles.card} href="/#">ブックマーク</Link>
                <Link className={mainStyles.card} href="/#">コミュニティ</Link>
                <Link className={mainStyles.card} href="/#">プレミアム</Link>
                <Link className={mainStyles.card} href="/#">プロフィール</Link>
                <Link className={mainStyles.card} href="/#">もっと見る</Link>
            </div>
        </nav>
    )
}