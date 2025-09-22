export type AccountStatus = 'Activated' |'Deactivated' | 'Blocked';

export interface UserView {
  id: string;
  username: string;
  email: string;
  role: string;

  account_status: AccountStatus;
}
