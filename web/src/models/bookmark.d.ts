export interface Bookmark {
  id: string;
  name: string;
  url: string;
  color?: string;
}

export interface BookmarkCreate {
  name: string;
  url: string;
  color?: string;
}

export interface BookmarkUpdate {
  id: string;
  name: string;
  url: string;
  color?: string;
}
