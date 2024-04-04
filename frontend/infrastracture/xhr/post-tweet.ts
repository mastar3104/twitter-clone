export const postTweet = async (content: string, userId: string) => {
    if (content === "") {
        return
    }
    await fetch(`${process.env.API_SERVER}/v1/tweets/${userId}`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            content: content,
        })
    })
}