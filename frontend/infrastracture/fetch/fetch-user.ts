import {RequestInfo} from "@/infrastracture/fetch/request-info";
import {UserResource, UsersResource} from "@/infrastracture/resource/user-resource";

export function FetchUser(
    userId: string,
): UserRequest {
    const request = fetch(`${process.env.API_SERVER}/v1/users/${userId}`)
    return new UserRequest(request)
}

export function FetchUsers(): UserRequest {
    const request = fetch(`${process.env.API_SERVER}/v1/users`)
    return new UserRequest(request)
}

export class UserRequest {
    request: Promise<RequestInfo>
    constructor(request: Promise<RequestInfo>) {
        this.request = request
    }
    async getResource (): Promise<UserResource> {
        return (await (await this.request).json())
    }

    async getListResource (): Promise<UsersResource> {
        return (await (await this.request).json())
    }
}
