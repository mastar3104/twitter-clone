import styles from './page.module.css'
import React from "react";
import Login from "@/components/Login/component";

export default function Home() {

  return (
    <main className={styles.main}>
      <h2>Twitter Clone</h2>
      <Login></Login>
    </main>
  )
}
