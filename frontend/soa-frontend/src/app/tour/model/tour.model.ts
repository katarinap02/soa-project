export interface Tour {
  id?: string;
  name: string;
  description?: string;
  price?: number;
  difficulty?: string;
  tags?: string[];
  status?: string;
  authorId?: string;
  publishDate?: string;
  archiveDate?: string;

  durations?: {
    walking?: number; // in minutes
    bicycle?: number;
    car?: number;
  };
}
