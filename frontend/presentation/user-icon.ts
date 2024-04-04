import { createHash } from 'crypto';

export const userIcon = (userId: string) : string => {
    const hash = createHash('sha256');
    hash.update(userId);
    const hashValue = hash.digest('hex');

    const decimalValue = parseInt(hashValue.slice(0, 2), 16);

    switch (decimalValue % 4) {
        case 0:
            return "/user1.png"
        case 1:
            return "/user2.png"
        case 2:
            return "/user3.png"
        default:
            return "/user4.png"
    }
}