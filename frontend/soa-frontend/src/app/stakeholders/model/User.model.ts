export type UserRole = 'Tourist' | 'Guide';
export type AccountStatus = 'Activated' |'Deactivated' | 'Blocked';

export interface User {
  id?: string;
  username: string;
  email: string;
  password?: string;
  role: UserRole;
  account_status: AccountStatus;
}
