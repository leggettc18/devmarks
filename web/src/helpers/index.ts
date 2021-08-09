export const asyncHandler: <T>(promise: Promise<T>) => Promise<[T, null] | [null, unknown]> = async <T>(promise: Promise<T>) => {
    try {
        const data = await promise;
        return [data, null];
    } catch(error){
        console.error(error);
        return [null, error];
    }
}