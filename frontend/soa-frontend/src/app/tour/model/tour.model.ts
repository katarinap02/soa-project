export interface Tour {
  id?: string;
  name: string;
  description?: string;
  price?: number;
  difficulty?: string;
  tags?: string[];
  status?: string;
  authorId?: string;
}
