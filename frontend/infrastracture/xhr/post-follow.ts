export const postFollow = async (userId: string, followUserId: string) => {
    await fetch(`${process.env.API_SERVER}/v1/users/${userId}/follow`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            followUserId: followUserId,
        })
    })
}