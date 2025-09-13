export interface Comment {
  id?: string;
  username: string;
  text: string;
  postId: string | null;
  date_created: string; // ISO string, npr. "2025-07-25T12:00:00Z"
  date_modified: string; // ISO string, npr. "2025-07-25T12:00:00Z"
}