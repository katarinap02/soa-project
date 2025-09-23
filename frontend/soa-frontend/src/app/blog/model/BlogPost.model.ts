export interface BlogPost {
  id?: string;
  username: string;
  title: string;
  description: string;
  date: string;
  Likes?: any[];

  likesCount?: number;
  likedByCurrentUser?: boolean;
}
