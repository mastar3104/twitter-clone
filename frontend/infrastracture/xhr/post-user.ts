import {UserResource} from "@/infrastracture/resource/user-resource";

export const postUser = async (userName: string): Promise<UserResource> => {
    const request = await fetch(`${process.env.API_SERVER}/v1/users`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            userName: userName,
        })
    })
    return await request.json()
}