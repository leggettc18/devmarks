import { UserApi, BookmarkApi, FolderApi, Configuration } from "@/api/client";

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
