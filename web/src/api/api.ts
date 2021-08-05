import { UserApi, BookmarkApi, FolderApi, Configuration } from "@/api/client";
import { inject, provide, reactive } from "vue";

export const DefaultDevmarksClient: unique symbol = Symbol(
  "DefaultDevmarksClient"
);

export class DevmarksClient {
  private userApi: UserApi;
  private bookmarkApi: BookmarkApi;
  private folderApi: FolderApi;

  constructor(configuration: Configuration = new Configuration()) {
    this.userApi = new UserApi(configuration);
    this.bookmarkApi = new BookmarkApi(configuration);
    this.folderApi = new FolderApi(configuration);
  }
}

export const useApi = () => inject(DefaultDevmarksClient) as DevmarksClient;
