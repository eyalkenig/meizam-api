import { mutations } from './mutations'
import { actions } from './actions'
import { TeamState } from './types';

export const initialStateData: TeamState = {
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
  }]
}

export default {
  state: { ...initialStateData },
  mutations,
  actions
}