import { UserApi, BookmarkApi, FolderApi, Configuration } from "@/api/client";
import { inject, provide, reactive } from "vue";

export const DefaultDevmarksClient: unique symbol = Symbol(
  "DefaultDevmarksClient"
);

export class DevmarksClient {
  public userApi: UserApi;
  public bookmarkApi: BookmarkApi;
  private folderApi: FolderApi;

  constructor(configuration: Configuration = new Configuration()) {
    this.userApi = new UserApi(configuration);
    this.bookmarkApi = new BookmarkApi(configuration);
    this.folderApi = new FolderApi(configuration);
  }
}

export const createApi = (config: Configuration) => {
  return reactive(new DevmarksClient(config));
}

export const useApi = () => inject(DefaultDevmarksClient) as DevmarksClient;
export const provideApi = (config: Configuration) => provide(DefaultDevmarksClient, createApi(config));
