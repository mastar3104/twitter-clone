
export const convertUnixTimeToYMD = (unixTime: number) => {
    const date = new Date(unixTime * 1000);
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1);
    const day = String(date.getDate());

    return `${year}-${month}-${day}`;
}