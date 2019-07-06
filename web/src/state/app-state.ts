import { Team } from '@/store/modules/teams/types';

export interface Competition {
  id: number;
  name: string;
  type: string;
  external_entity_id: string|null;
  created_at: Date;
  updated_at: Date;
}

export interface DataState {
  teams: Team[];
  competitions: Competition[];
}

export const initialStateData: DataState = {
  teams: [],
  competitions: []
}

export interface UIState {
}
  
export const initialStateUi: UIState = {}

export interface AppState {
  data: DataState;
  ui: UIState;
}

export const initialState: AppState = {
  data: initialStateData,
  ui: initialStateUi
};