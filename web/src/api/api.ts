import { UserApi, BookmarkApi, FolderApi, Configuration, Bookmark } from "@/api/client";
import { BookmarkCreate, BookmarkUpdate } from "@/models/bookmark";
import { AxiosError, AxiosResponse } from "axios";
import { inject, provide, reactive } from "vue";

export const DefaultDevmarksClient: unique symbol = Symbol(
  "DefaultDevmarksClient"
);

interface Response<T> {
  success: boolean;
  response?: AxiosResponse;
  statusCode: number;
  data?: T;
  error?: AxiosError;
  message?: string;
}

const isAxiosError = (error: unknown): error is AxiosError => {
  return (error as AxiosError).response !== undefined;
};

const handleRequest: <T>(request: Promise<AxiosResponse<T>>) => Promise<Response<T>> = async <T>(request: Promise<AxiosResponse<T>>) => {
  try {
    const response = await request;
    return {
        success: true,
        response: response,
        data: response.data,
        statusCode: response.status,
      };
  } catch(error) {
    if(isAxiosError(error) && error.response) {
       return {
         success: false,
         error: error,
         statusCode: error.response.status,
         message: error.message,
       };
    } else {
      throw new Error("Malformed Response");
    }
  }
}

export class DevmarksClient {
  public userApi: UserApi;
  public bookmarkApi: BookmarkApi;
  private folderApi: FolderApi;

  constructor(configuration: Configuration = new Configuration()) {
    this.userApi = new UserApi(configuration);
    this.bookmarkApi = new BookmarkApi(configuration);
    this.folderApi = new FolderApi(configuration);
  }

  /**
   * getBookmarks
   */
  public async getBookmarks(): Promise<Response<Bookmark[]>> {
    return await handleRequest(
      this.bookmarkApi.getBookmarks()
    );
  }

  public async newBookmark(bookmark: BookmarkCreate) {
    return await handleRequest(
      this.bookmarkApi.createBookmark(bookmark)
    )
  }

  public async updateBookmark(bookmark: BookmarkUpdate) {
    return await handleRequest(
      this.bookmarkApi.updateBookmark(bookmark.id, bookmark)
    )
  }
}

export const createApi = (config: Configuration) => {
  return reactive(new DevmarksClient(config));
};

export const useApi = () => inject(DefaultDevmarksClient) as DevmarksClient;
export const provideApi = (config: Configuration) =>
  provide(DefaultDevmarksClient, createApi(config));
