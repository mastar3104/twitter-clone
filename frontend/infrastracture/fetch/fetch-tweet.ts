import {TweetListResource} from "@/infrastracture/resource/tweet-resource";
import {RequestInfo} from "@/infrastracture/fetch/request-info";

export function FetchTimelineList(
    userId: string,
): TweetRequest {
    const request = fetch(`${process.env.API_SERVER}/v1/timeline/${userId}`)
    return new TweetRequest(request)
}

export function FetchTweetList(
    userId: string,
): TweetRequest {
    const request = fetch(`${process.env.API_SERVER}/v1/tweets/${userId}`)
    return new TweetRequest(request)
}

export class TweetRequest {
    request: Promise<RequestInfo>

    constructor(request: Promise<RequestInfo>) {
        this.request = request
    }

    async getResource(): Promise<TweetListResource> {
        return (await this.request).json()
    }
}