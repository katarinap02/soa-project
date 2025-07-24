export interface Comment {

  Username: string;
  Text: string;
  PostId: string;
  DateCreated: string; // ISO string, npr. "2025-07-25T12:00:00Z"
  DateModified: string; // ISO string, npr. "2025-07-25T12:00:00Z"
}