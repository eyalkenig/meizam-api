export const FETCH_TEAMS = 'teams/FETCH_TEAMS'
export const SET_TEAMS = 'teams/SET_TEAMS'
export const GET_TEAMS = 'teams/GET_TEAMS'

export interface Team {
    id: number;
    name: string;
    image_url: string|null;
    external_entity_id: string|null;
    created_at: Date;
    updated_at: Date;
}

export interface TeamState {
    teams: Team[];
}