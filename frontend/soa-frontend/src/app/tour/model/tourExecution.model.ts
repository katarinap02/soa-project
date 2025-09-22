export type TourExecutionStatus = 'active' | 'completed' | 'abandoned';

export interface TourExecution {
  id?: string;             
  tourId: string;
  touristId: string;
  lastActivity: Date;
  status: TourExecutionStatus;
}
