export interface UserResource {
    userId: string,
    userName: string,
}

export interface UsersResource {
    users: UserResource[]
}