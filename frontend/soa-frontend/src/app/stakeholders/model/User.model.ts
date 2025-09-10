export type UserRole = 'Tourist' | 'Guide'; 

export interface User {
  id?: string;       
  username: string;
  email: string;
  password?: string; 
  role: UserRole;
}