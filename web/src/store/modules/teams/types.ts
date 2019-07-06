export const FETCH_TEAMS = 'teams/FETCH_TEAMS'
export const SET_TEAMS = 'teams/SET_TEAMS'
export const GET_TEAMS = 'teams/GET_TEAMS'

export interface Team {
    id: number;
    name: string;
    imageUrl: string|null;
    externalEntityId: string|null;
    createdAt: Date;
    updatedAt: Date;
}

export interface TeamState {
    teams: Team[];
}
