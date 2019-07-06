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
  teams: [{
    "id": 1,
    "name": "kenig",
    "image_url": "test2",
    "external_entity_id": '123',
    "created_at": new Date("2019-05-10T08:28:01Z"),
    "updated_at": new Date("2019-05-10T08:28:01Z")
  },
  {
      "id": 3,
      "name": "kenig24",
      "image_url": "test23",
      "external_entity_id": null,
      "created_at": new Date("2019-05-12T18:53:44Z"),
      "updated_at": new Date("2019-05-12T18:53:44Z")
  }],
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